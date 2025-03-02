// TODO: убрать мусор отсюда сделать нормальную инициализацию
package main

import (
	"context"
	"log/slog"
	"net"
	"os"

	dbuslib "github.com/WantBeASleep/med_ml_lib/dbus"
	grpclib "github.com/WantBeASleep/med_ml_lib/grpc"
	observerdbuslib "github.com/WantBeASleep/med_ml_lib/observer/dbus"
	observergrpclib "github.com/WantBeASleep/med_ml_lib/observer/grpc"
	loglib "github.com/WantBeASleep/med_ml_lib/observer/log"

	"uzi/internal/config"

	"github.com/ilyakaznacheev/cleanenv"

	"uzi/internal/repository"

	devicesrv "uzi/internal/services/device"
	imagesrv "uzi/internal/services/image"
	isnsrv "uzi/internal/services/image-segment-node"
	nodesrv "uzi/internal/services/node"
	segmentsrv "uzi/internal/services/segment"
	uzisrv "uzi/internal/services/uzi"

	pb "uzi/internal/generated/grpc/service"

	grpchandler "uzi/internal/controllers/grpc"
	devicehandler "uzi/internal/controllers/grpc/device"
	imagehandler "uzi/internal/controllers/grpc/image"
	isnhandler "uzi/internal/controllers/grpc/image-segment-node"
	nodehandler "uzi/internal/controllers/grpc/node"
	segmenthandler "uzi/internal/controllers/grpc/segment"
	uzihandler "uzi/internal/controllers/grpc/uzi"

	uziprocessedsubscriber "uzi/internal/controllers/dbus/uziprocessed"
	uziuploadsubscriber "uzi/internal/controllers/dbus/uziupload"

	dbusadapters "uzi/internal/adapters/dbus"
	uziprocessed "uzi/internal/generated/dbus/consume/uziprocessed"
	uziupload "uzi/internal/generated/dbus/consume/uziupload"
	uzicompletepb "uzi/internal/generated/dbus/produce/uzicomplete"
	uzisplittedpb "uzi/internal/generated/dbus/produce/uzisplitted"

	"github.com/IBM/sarama"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"google.golang.org/grpc"
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

	if err := db.Ping(); err != nil {
		slog.Error("ping db", "err", err)
		return failExitCode
	}

	producer, err := sarama.NewSyncProducer(cfg.Broker.Addrs, nil)
	if err != nil {
		slog.Error("init sarama producer", "err", err)
		return failExitCode
	}

	producerUziSplitted := dbuslib.NewProducer[*uzisplittedpb.UziSplitted](
		producer,
		"uzisplitted",
		dbuslib.WithProducerMiddlewares[*uzisplittedpb.UziSplitted](
			observerdbuslib.CrossEventProduce,
			observerdbuslib.LogEventProduce,
		),
	)

	producerUziComplete := dbuslib.NewProducer[*uzicompletepb.UziComplete](
		producer,
		"uzicomplete",
		dbuslib.WithProducerMiddlewares[*uzicompletepb.UziComplete](
			observerdbuslib.CrossEventProduce,
			observerdbuslib.LogEventProduce,
		),
	)

	dbusAdapter := dbusadapters.New(producerUziSplitted, producerUziComplete)

	dao := repository.NewRepository(db, client, "uzi")

	deviceSrv := devicesrv.New(dao)
	uziSrv := uzisrv.New(dao)
	imageSrv := imagesrv.New(dao, dbusAdapter)
	isnSrv := isnsrv.New(dao)
	nodeSrv := nodesrv.New(dao)
	serviceSrv := segmentsrv.New(dao)

	// grpc
	deviceHandler := devicehandler.New(deviceSrv)
	uziHandler := uzihandler.New(uziSrv)
	imageHandler := imagehandler.New(imageSrv)
	isnHandler := isnhandler.New(isnSrv)
	nodeHandler := nodehandler.New(nodeSrv)
	segmentHandler := segmenthandler.New(serviceSrv)

	handler := grpchandler.New(
		deviceHandler,
		uziHandler,
		imageHandler,
		isnHandler,
		segmentHandler,
		nodeHandler,
	)

	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpclib.PanicRecover,
			observergrpclib.CrossServerCall,
			observergrpclib.LogServerCall,
		),
	)
	pb.RegisterUziSrvServer(server, handler)

	// dbus
	uziuploadSubscriber := uziuploadsubscriber.New(imageSrv)
	uziprocessedSubscriber := uziprocessedsubscriber.New(nodeSrv)

	uziUploadHandler := dbuslib.NewGroupSubscriber(
		"uziupload",
		cfg.Broker.Addrs,
		"uziupload",
		uziuploadSubscriber,
		dbuslib.WithSubscriberMiddlewares[*uziupload.UziUpload](
			observerdbuslib.CrossEventConsume,
			observerdbuslib.LogEventConsume,
		),
	)

	uziprocessedHandler := dbuslib.NewGroupSubscriber(
		"uziprocessed",
		cfg.Broker.Addrs,
		"uziprocessed",
		uziprocessedSubscriber,
		dbuslib.WithSubscriberMiddlewares[*uziprocessed.UziProcessed](
			observerdbuslib.CrossEventConsume,
			observerdbuslib.LogEventConsume,
		),
	)

	lis, err := net.Listen("tcp", cfg.App.Url)
	if err != nil {
		slog.Error("take port", "err", err)
		return failExitCode
	}

	close := make(chan struct{})
	// ЛЮТОЕ MVP
	slog.Info("start serve", slog.String("app url", cfg.App.Url))
	go func() {
		if err := server.Serve(lis); err != nil {
			slog.Error("take port", "err", err)
			panic("serve grpc")
		}
		close <- struct{}{}
	}()
	go func() {
		// пока без DI
		if err := uziUploadHandler.Start(context.Background()); err != nil {
			slog.Error("start uziupload handler", "err", err)
		}
	}()
	go func() {
		if err := uziprocessedHandler.Start(context.Background()); err != nil {
			slog.Error("start uziprocessedHandler handler", "err", err)
		}
	}()

	<-close

	return successExitCode
}
