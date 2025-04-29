package main

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"os"
	"time"

	"billing/internal/services/subscription_checker"

	"billing/internal/config"
	pb "billing/internal/generated/grpc/service"
	"billing/internal/repository"
	grpchandler "billing/internal/server"
	"billing/internal/services"

	grpclib "github.com/WantBeASleep/med_ml_lib/grpc"
	observergrpclib "github.com/WantBeASleep/med_ml_lib/observer/grpc"
	loglib "github.com/WantBeASleep/med_ml_lib/observer/log"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
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
	fmt.Println("DB_DSN:", os.Getenv("DB_DSN"))

	db, err := sqlx.Open("postgres", cfg.DB.Dsn)
	if err != nil {
		slog.Error("init db", "err", err)
		return failExitCode
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		slog.Error("ping db", "err", err)
		return failExitCode
	}

	dao := repository.NewRepository(db)

	services := services.New(
		dao,
		&cfg,
	)

	handler := grpchandler.New(services)

	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpclib.PanicRecover,
			observergrpclib.CrossServerCall,
			observergrpclib.LogServerCall,
		),
	)
	pb.RegisterBillingServiceServer(server, handler)

	lis, err := net.Listen("tcp", cfg.App.Url)
	if err != nil {
		slog.Error("take port", "err", err)
		return failExitCode
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	subscriptionChecker := subscription_checker.New(services.Subscription)
	go subscriptionChecker.Start(ctx, time.Minute)

	close := make(chan struct{})
	slog.Info("start serve", slog.String("app url", cfg.App.Url))
	go func() {
		if err := server.Serve(lis); err != nil {
			slog.Error("serve grpc", "err", err)
			panic("serve grpc")
		}
		close <- struct{}{}
	}()

	<-close

	return successExitCode
}
