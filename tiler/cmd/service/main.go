package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"

	loglib "github.com/WantBeASleep/med_ml_lib/observer/log"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"

	"tiler/internal/config"
	"tiler/internal/server"
	"tiler/internal/services"
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

	// Инициализируем S3 клиент
	s3Client, err := minio.New(cfg.S3.Endpoint, &minio.Options{
		Secure: false,
		Creds:  credentials.NewStaticV4(cfg.S3.Access_Token, cfg.S3.Secret_Token, ""),
	})
	if err != nil {
		slog.Error("init s3", "err", err)
		return failExitCode
	}

	// Проверяем существование bucket
	bucketName := cfg.S3.BucketName
	exists, err := s3Client.BucketExists(context.Background(), bucketName)
	if err != nil {
		slog.Error("check bucket exists", "err", err)
		return failExitCode
	}
	if !exists {
		err = s3Client.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{})
		if err != nil {
			slog.Error("create bucket", "err", err)
			return failExitCode
		}
		slog.Info("created bucket", "bucket", bucketName)
	}

	// Создаем сервисы
	s3Service := services.NewS3Client(s3Client, bucketName)
	imageService := services.NewImageService(s3Service, cfg.App.TileSize, cfg.App.Overlap)

	// Создаем HTTP сервер
	httpHandler := server.NewServer(imageService)

	// Запускаем HTTP сервер
	slog.Info("start serve", slog.String("app url", cfg.App.URL))
	if err := http.ListenAndServe(cfg.App.URL, httpHandler); err != nil {
		slog.Error("serve http", "err", err)
		return failExitCode
	}

	return successExitCode
}
