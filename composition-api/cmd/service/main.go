package main

import (
	_ "embed"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"strings"

	"composition-api/internal/dbus/producers"

	"github.com/IBM/sarama"

	"time"

	loglib "github.com/WantBeASleep/med_ml_lib/observer/log"
	"github.com/flowchartsman/swaggerui"
	"github.com/go-chi/chi/v5"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"

	"composition-api/internal/adapters"
	"composition-api/internal/config"
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

	billingConn, err := grpc.NewClient(
		cfg.Adapters.BillingUrl,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		slog.Error("init billingConn", slog.Any("err", err))
		return failExitCode
	}

	// Увеличиваем максимальный размер сообщения для cytology (для передачи больших изображений)
	// 4GB должно быть достаточно для больших медицинских изображений
	const maxMsgSize = 4 * 1024 * 1024 * 1024 // 4GB

	// Пытаемся подключиться с несколькими попытками, так как сервис может быть еще не готов
	var cytologyConn *grpc.ClientConn
	maxRetries := 5
	for i := 0; i < maxRetries; i++ {
		conn, err := grpc.NewClient(
			cfg.Adapters.CytologyUrl,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithKeepaliveParams(keepalive.ClientParameters{
				Time:                10 * time.Minute,
				Timeout:             20 * time.Second,
				PermitWithoutStream: true,
			}),
			grpc.WithDefaultCallOptions(
				grpc.MaxCallRecvMsgSize(maxMsgSize),
				grpc.MaxCallSendMsgSize(maxMsgSize),
			),
		)
		if err == nil {
			cytologyConn = conn
			break
		}
		if i < maxRetries-1 {
			slog.Warn("failed to connect to cytology service, retrying", slog.Int("attempt", i+1), slog.Any("err", err))
			time.Sleep(2 * time.Second)
		} else {
			slog.Error("init cytologyConn", slog.Any("err", err))
			return failExitCode
		}
	}

	adapters := adapters.NewAdapters(uziConn, authConn, medConn, billingConn, cytologyConn, cfg.Adapters.TilerUrl)

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

	// Создаем сервер с увеличенным лимитом для multipart форм (4GB)
	server, err := api.NewServer(
		handlers,
		security,
		api.WithMaxMultipartMemory(4<<30), // 4GB для больших файлов
	)
	if err != nil {
		slog.Error("init server", slog.Any("err", err))
		return failExitCode
	}

	r := chi.NewRouter()

	r.Mount("/api/v1/", http.StripPrefix("/api/v1", server))
	r.Mount("/docs/", http.StripPrefix("/docs", swaggerui.Handler(spec)))

	// Проксирование запросов к tiler напрямую на tiler_service
	r.Mount("/tiler/", http.StripPrefix("/tiler", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		// Формируем URL для проксирования на tiler_service
		tilerURL := cfg.Adapters.TilerUrl
		if !strings.HasPrefix(tilerURL, "http://") && !strings.HasPrefix(tilerURL, "https://") {
			tilerURL = "http://" + tilerURL
		}
		// Убираем завершающий слэш, если есть
		tilerURL = strings.TrimSuffix(tilerURL, "/")

		// Правильно формируем URL
		proxyURL, err := url.Parse(tilerURL)
		if err != nil {
			slog.Error("tiler proxy: invalid tiler URL", "tiler_url", tilerURL, "err", err)
			http.Error(w, fmt.Sprintf("Invalid tiler URL: %v", err), http.StatusInternalServerError)
			return
		}
		// Объединяем базовый URL с путем запроса
		// r.URL.Path уже содержит ведущий слэш после StripPrefix
		proxyURL.Path = r.URL.Path

		proxyURLStr := proxyURL.String()
		slog.Info("tiler proxy: proxying request",
			"method", r.Method,
			"original_path", r.URL.Path,
			"proxy_url", proxyURLStr,
			"remote_addr", r.RemoteAddr,
		)

		// Создаем новый запрос к tiler_service
		proxyReq, err := http.NewRequest(r.Method, proxyURLStr, r.Body)
		if err != nil {
			slog.Error("tiler proxy: failed to create request", "proxy_url", proxyURLStr, "err", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Копируем заголовки
		for key, values := range r.Header {
			for _, value := range values {
				proxyReq.Header.Add(key, value)
			}
		}

		// Выполняем запрос с увеличенными таймаутами для больших файлов
		client := &http.Client{
			Timeout: 30 * time.Minute, // 30 минут для загрузки больших файлов
		}
		resp, err := client.Do(proxyReq)
		duration := time.Since(startTime)
		if err != nil {
			slog.Error("tiler proxy: request failed",
				"proxy_url", proxyURLStr,
				"original_path", r.URL.Path,
				"err", err,
				"duration_ms", duration.Milliseconds(),
			)
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}
		defer resp.Body.Close()

		slog.Info("tiler proxy: request completed",
			"proxy_url", proxyURLStr,
			"original_path", r.URL.Path,
			"status_code", resp.StatusCode,
			"duration_ms", duration.Milliseconds(),
		)

		// Копируем заголовки ответа
		for key, values := range resp.Header {
			for _, value := range values {
				w.Header().Add(key, value)
			}
		}

		// Копируем статус и тело ответа
		w.WriteHeader(resp.StatusCode)
		io.Copy(w, resp.Body)
	})))

	// Настройка HTTP сервера с увеличенными таймаутами для больших файлов
	srv := &http.Server{
		Addr:           cfg.App.Url,
		Handler:        r,
		ReadTimeout:    30 * time.Minute,  // 30 минут на чтение запроса
		WriteTimeout:   30 * time.Minute,  // 30 минут на запись ответа
		IdleTimeout:    120 * time.Second, // 2 минуты для idle соединений
		MaxHeaderBytes: 1 << 20,           // 1MB для заголовков
	}

	slog.Info("start serve", slog.String("url", cfg.App.Url))
	if err := srv.ListenAndServe(); err != nil {
		slog.Error("listen and serve", slog.Any("err", err))
		return failExitCode
	}

	return successExitCode
}
