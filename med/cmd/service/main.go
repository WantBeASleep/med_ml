package main

import (
	"log/slog"
	"net"
	"os"

	grpclib "github.com/WantBeASleep/med_ml_lib/grpc"
	observergrpclib "github.com/WantBeASleep/med_ml_lib/observer/grpc"
	loglib "github.com/WantBeASleep/med_ml_lib/observer/log"

	"med/internal/config"

	"github.com/ilyakaznacheev/cleanenv"
	_ "github.com/joho/godotenv/autoload"

	"med/internal/repository"

	cardsrv "med/internal/services/card"
	doctorsrv "med/internal/services/doctor"
	patientsrv "med/internal/services/patient"

	pb "med/internal/generated/grpc/service"
	grpchandler "med/internal/grpc"

	cardhandler "med/internal/grpc/card"
	doctorhandler "med/internal/grpc/doctor"
	patienthandler "med/internal/grpc/patient"

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

	patientSrv := patientsrv.New(dao)
	doctorSrv := doctorsrv.New(dao)
	cardSrv := cardsrv.New(dao)

	patientHandler := patienthandler.New(patientSrv)
	doctorHandler := doctorhandler.New(doctorSrv)
	cardHandler := cardhandler.New(cardSrv)

	handler := grpchandler.New(
		patientHandler,
		doctorHandler,
		cardHandler,
	)

	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpclib.PanicRecover,
			observergrpclib.CrossServerCall,
			observergrpclib.LogServerCall,
		),
	)
	pb.RegisterMedSrvServer(server, handler)

	lis, err := net.Listen("tcp", cfg.App.Url)
	if err != nil {
		slog.Error("take port", "err", err)
		return failExitCode
	}

	slog.Info("start serve", slog.String("app url", cfg.App.Url))
	if err := server.Serve(lis); err != nil {
		slog.Error("take port", "err", err)
		return failExitCode
	}

	return successExitCode
}
