package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"chat_service/internal/config"
	grpc_server "chat_service/internal/controller/grpc"
	http_server "chat_service/internal/controller/http"
	"chat_service/internal/core/cache_service"
	"chat_service/internal/core/chat_service"
	"chat_service/internal/core/message_service"
	"chat_service/internal/repository/med"
	"chat_service/internal/repository/postgres"
	"chat_service/internal/repository/redis"
	"chat_service/internal/usecase"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// TODO: make app package instead of main
func main() {
	log.Logger = zerolog.New(zerolog.ConsoleWriter{
		Out: os.Stderr,
	}).
		Level(zerolog.DebugLevel).
		With().
		Timestamp().
		Caller().
		Logger()

	// TODO: Try to load config file, but don't fail if it doesn't exist
	configPath := "./config/config.yaml"
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		configPath = ""
		log.Warn().
			Msg("config file not found, using environment variables and defaults")
	}

	cfg, err := config.Parse(configPath)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("load config")
	}

	log.Info().
		Str("postgresDSN", cfg.Postgres.DSN).
		Str("redisAddr", cfg.Redis.Addr).
		Str("medServiceAddr", cfg.MedService.Addr).
		Msg("Configuration loaded")

	pg, err := postgres.New(cfg.Postgres)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("init postgres")
	}

	defer pg.Close()

	redisRepo, err := redis.New(cfg.Redis)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("init redis")
	}

	defer redisRepo.Close()

	medClient, err := med.New(cfg.MedService.Addr)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("init med client")
	}

	defer medClient.Close()

	// Repositories
	chatRepo := postgres.NewChatRepository(pg)
	messageRepo := postgres.NewMessageRepository(pg)

	// Core services
	cacheService := cache_service.NewCacheService(redisRepo, medClient, cfg.MedService.CacheTTL)

	// Domain services
	chatSvc := chat_service.NewChatService(chatRepo)
	messageSvc := message_service.NewMessageService(messageRepo, cacheService)

	// Use cases
	chatUsecase := usecase.NewChatUsecase(chatSvc, messageSvc)

	// Servers
	httpServer := http_server.NewServer(cfg, chatSvc, messageSvc)

	grpcAddr := fmt.Sprintf("%s:%d", cfg.GRPC.Host, cfg.GRPC.Port)

	grpcServer, err := grpc_server.NewServer(grpcAddr, chatUsecase)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("create gRPC server")
	}

	// Start HTTP server
	go func() {
		log.Info().
			Str("host", cfg.HTTP.Host).
			Uint16("port", cfg.HTTP.Port).
			Msg("starting HTTP server")

		if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal().
				Err(err).
				Msg("HTTP server")
		}
	}()

	// Start gRPC server
	go func() {
		log.Info().
			Str("host", cfg.GRPC.Host).
			Uint16("port", cfg.GRPC.Port).
			Msg("starting gRPC server")

		if err := grpcServer.Start(); err != nil {
			log.Fatal().
				Err(err).
				Msg("gRPC server")
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	log.Info().
		Msg("service started")

	<-quit

	log.Info().
		Msg("shutdown signal received")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		log.Error().
			Err(err).
			Msg("HTTP server shutdown")
	}

	grpcServer.Stop()

	log.Info().
		Msg("service stopped")
}
