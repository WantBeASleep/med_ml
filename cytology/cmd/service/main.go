// TODO: убрать мусор отсюда сделать нормальную инициализацию
package main

import (
	"context"
	"log/slog"
	"net"
	"os"

	grpclib "github.com/WantBeASleep/med_ml_lib/grpc"
	observergrpclib "github.com/WantBeASleep/med_ml_lib/observer/grpc"
	loglib "github.com/WantBeASleep/med_ml_lib/observer/log"

	"cytology/internal/config"

	"github.com/ilyakaznacheev/cleanenv"

	"cytology/internal/repository"

	services "cytology/internal/services"

	pb "cytology/internal/generated/grpc/service"

	grpchandler "cytology/internal/server"

	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

const (
	successExitCode = 0
	failExitCode    = 1
)

func main() {
	os.Exit(run())
}

func run() (exitCode int) {
	loglib.InitLogger(loglib.WithEnv())

	cfg := config.Config{}
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		slog.Error("init config", "err", err)
		return failExitCode

	}

	db, err := sqlx.Open("postgres", cfg.DB.Dsn)
	if err != nil {
		slog.Error("init db", "err", err)
		return failExitCode
	}
	defer db.Close()

	client, err := minio.New(cfg.S3.Endpoint, &minio.Options{
		Secure: false,
		Creds:  credentials.NewStaticV4(cfg.S3.Access_Token, cfg.S3.Secret_Token, ""),
	})
	if err != nil {
		slog.Error("init s3", "err", err)
		return failExitCode
	}

	bucketName := "cytology"
	exists, err := client.BucketExists(context.Background(), bucketName)
	if err != nil {
		slog.Error("check bucket exists", "err", err)
		return failExitCode
	}
	if !exists {
		err = client.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{})
		if err != nil {
			slog.Error("create bucket", "err", err)
			return failExitCode
		}
		slog.Info("created bucket", "bucket", bucketName)
	}

	if err := db.Ping(); err != nil {
		slog.Error("ping db", "err", err)
		return failExitCode
	}

	dao := repository.NewRepository(db, client, bucketName)

	services := services.New(
		dao,
	)

	handler := grpchandler.New(services)

	// Увеличиваем максимальный размер сообщения для передачи больших изображений
	// 4GB должно быть достаточно для больших медицинских изображений
	const maxMsgSize = 4 * 1024 * 1024 * 1024 // 4GB
	server := grpc.NewServer(
		grpc.MaxRecvMsgSize(maxMsgSize),
		grpc.MaxSendMsgSize(maxMsgSize),
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle:     30 * time.Minute,
			MaxConnectionAge:      1 * time.Hour,
			MaxConnectionAgeGrace: 10 * time.Minute,
			Time:                  10 * time.Minute,
			Timeout:               20 * time.Second,
		}),
		grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
			MinTime:             10 * time.Second,
			PermitWithoutStream: true,
		}),
		grpc.ChainUnaryInterceptor(
			grpclib.PanicRecover,
			observergrpclib.CrossServerCall,
			observergrpclib.LogServerCall,
		),
	)
	pb.RegisterCytologySrvServer(server, handler)

	lis, err := net.Listen("tcp", cfg.App.Url)
	if err != nil {
		slog.Error("take port", "err", err)
		return failExitCode
	}

	close := make(chan struct{})
	slog.Info("start serve", slog.String("app url", cfg.App.Url))
	go func() {
		if err := server.Serve(lis); err != nil {
			slog.Error("take port", "err", err)
			panic("serve grpc")
		}
		close <- struct{}{}
	}()

	<-close

	return successExitCode
}
