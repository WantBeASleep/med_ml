package main

import (
	_ "embed"
	"log/slog"
	"net/http"
	"os"

	"github.com/IBM/sarama"
	loglib "github.com/WantBeASleep/med_ml_lib/observer/log"
	"github.com/flowchartsman/swaggerui"
	"github.com/go-chi/chi/v5"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"composition-api/internal/adapters"
	"composition-api/internal/config"
	"composition-api/internal/dbus/producers"
	api "composition-api/internal/generated/http/api"
	"composition-api/internal/repository"
	"composition-api/internal/server"
	"composition-api/internal/server/security"
	"composition-api/internal/services"
)

//go:embed server.yml
var spec []byte

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
		slog.Error("init config", slog.Any("err", err))
		return failExitCode
	}

	// adapters
	uziConn, err := grpc.NewClient(
		cfg.Adapters.UziUrl,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		slog.Error("init uziConn", slog.Any("err", err))
		return failExitCode
	}
	authConn, err := grpc.NewClient(
		cfg.Adapters.AuthUrl,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		slog.Error("init authConn", slog.Any("err", err))
		return failExitCode
	}
	medConn, err := grpc.NewClient(
		cfg.Adapters.MedUrl,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		slog.Error("init medConn", slog.Any("err", err))
		return failExitCode
	}

	adapters := adapters.NewAdapters(uziConn, authConn, medConn)

	// infra
	s3Client, err := minio.New(cfg.S3.Endpoint, &minio.Options{
		Secure: false,
		Creds:  credentials.NewStaticV4(cfg.S3.Access_Token, cfg.S3.Secret_Token, ""),
	})
	if err != nil {
		slog.Error("init s3", slog.Any("err", err))
		return failExitCode
	}

	dao := repository.NewRepository(s3Client, "uzi")

	dbusClient, err := sarama.NewSyncProducer(cfg.Dbus.Addrs, nil)
	if err != nil {
		slog.Error("init sarama producer", slog.Any("err", err))
		return failExitCode
	}

	producer := producers.New(dbusClient)

	// services
	services := services.New(adapters, producer, dao)

	// server
	handlers := server.New(services)

	// security
	security := security.New(&cfg)

	server, err := api.NewServer(handlers, security)
	if err != nil {
		slog.Error("init server", slog.Any("err", err))
		return failExitCode
	}

	r := chi.NewRouter()
	r.Mount("/api/v1/", http.StripPrefix("/api/v1", server))
	r.Mount("/swagger/", http.StripPrefix("/swagger", swaggerui.Handler(spec)))

	slog.Info("start serve", slog.String("url", cfg.App.Url))
	if err := http.ListenAndServe(cfg.App.Url, r); err != nil {
		slog.Error("listen and serve", slog.Any("err", err))
		return failExitCode
	}

	return successExitCode
}
