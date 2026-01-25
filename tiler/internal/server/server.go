package server

import (
	"log/slog"
	"net/http"
	"strings"
	"time"

	"tiler/internal/services"

	"github.com/rs/cors"
)

// loggingMiddleware логирует все входящие HTTP запросы
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		slog.Info("tiler: incoming request",
			"method", r.Method,
			"path", r.URL.Path,
			"query", r.URL.RawQuery,
			"remote_addr", r.RemoteAddr,
			"user_agent", r.UserAgent(),
			"host", r.Host,
		)

		// Создаем ResponseWriter для отслеживания статуса
		lrw := &loggingResponseWriter{ResponseWriter: w, statusCode: http.StatusOK}

		next.ServeHTTP(lrw, r)

		duration := time.Since(startTime)
		slog.Info("tiler: request completed",
			"method", r.Method,
			"path", r.URL.Path,
			"status_code", lrw.statusCode,
			"duration_ms", duration.Milliseconds(),
		)
	})
}

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func NewServer(imageService services.ImageService) http.Handler {
	handler := NewHandler(imageService)

	mux := http.NewServeMux()

	// DZI endpoint
	mux.HandleFunc("/dzi/", func(w http.ResponseWriter, r *http.Request) {
		// Проверяем, это запрос на DZI XML или на тайл
		// Если путь содержит "/files/", то это запрос на тайл
		if strings.Contains(r.URL.Path, "/files/") {
			// Это запрос на тайл
			handler.GetTile(w, r)
		} else {
			// Это запрос на DZI XML
			handler.GetDZI(w, r)
		}
	})

	// CORS middleware
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "OPTIONS"},
		AllowedHeaders: []string{"*"},
	})

	// Применяем логирование ко всем запросам
	return loggingMiddleware(c.Handler(mux))
}
