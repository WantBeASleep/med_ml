package server

import (
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"tiler/internal/services"
)

type Handler struct {
	imageService services.ImageService
}

func NewHandler(imageService services.ImageService) *Handler {
	return &Handler{
		imageService: imageService,
	}
}

// GetDZI возвращает DZI XML для DeepZoom
func (h *Handler) GetDZI(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	// Извлекаем путь к файлу из URL
	// Формат: /dzi/{file_path:path}
	path := strings.TrimPrefix(r.URL.Path, "/dzi/")
	if path == "" {
		slog.Warn("GetDZI: empty file path", "method", r.Method, "path", r.URL.Path, "remote_addr", r.RemoteAddr)
		http.Error(w, "File path is required", http.StatusBadRequest)
		return
	}

	// Декодируем URL-encoded путь
	decodedPath, err := url.PathUnescape(path)
	if err != nil {
		// Если декодирование не удалось, используем исходный путь
		decodedPath = path
		slog.Warn("GetDZI: failed to unescape path, using original", "path", path, "err", err)
	}

	slog.Info("GetDZI: request received",
		"method", r.Method,
		"path", r.URL.Path,
		"decoded_path", decodedPath,
		"remote_addr", r.RemoteAddr,
		"user_agent", r.UserAgent(),
	)

	ctx := r.Context()
	dzi, err := h.imageService.GetDZI(ctx, decodedPath)
	duration := time.Since(startTime)

	if err != nil {
		slog.Error("GetDZI: failed to get DZI",
			"decoded_path", decodedPath,
			"err", err,
			"duration_ms", duration.Milliseconds(),
		)
		http.Error(w, fmt.Sprintf("Failed to get DZI: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/xml")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(dzi.XML))

	slog.Info("GetDZI: success",
		"decoded_path", decodedPath,
		"xml_size", len(dzi.XML),
		"duration_ms", duration.Milliseconds(),
	)
}

// GetTile возвращает тайл изображения
func (h *Handler) GetTile(w http.ResponseWriter, r *http.Request) {
	// Обработка паники для предотвращения падения контейнера
	defer func() {
		if err := recover(); err != nil {
			slog.Error("GetTile: panic recovered",
				"err", err,
				"path", r.URL.Path,
				"remote_addr", r.RemoteAddr,
			)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
	}()

	startTime := time.Now()

	// Извлекаем параметры из URL
	// Формат: /dzi/{file_path:path}/files/{level:int}/{col:int}_{row:int}.{format:str}
	path := r.URL.Path

	// Парсим путь
	// Пример: /dzi/path/to/image/files/5/10_20.jpeg
	parts := strings.Split(path, "/files/")
	if len(parts) != 2 {
		slog.Warn("GetTile: invalid path format", "method", r.Method, "path", path, "remote_addr", r.RemoteAddr)
		http.Error(w, "Invalid tile path format", http.StatusBadRequest)
		return
	}

	filePath := strings.TrimPrefix(parts[0], "/dzi/")
	// Декодируем URL-encoded путь
	decodedFilePath, err := url.PathUnescape(filePath)
	if err != nil {
		// Если декодирование не удалось, используем исходный путь
		decodedFilePath = filePath
		slog.Warn("GetTile: failed to unescape path, using original", "path", filePath, "err", err)
	}
	tilePart := parts[1]

	// Парсим level/col_row.format
	tileParts := strings.Split(tilePart, "/")
	if len(tileParts) != 2 {
		slog.Warn("GetTile: invalid tile part format", "tile_part", tilePart, "path", path)
		http.Error(w, "Invalid tile path format", http.StatusBadRequest)
		return
	}

	level, err := strconv.Atoi(tileParts[0])
	if err != nil {
		slog.Warn("GetTile: invalid level", "level_str", tileParts[0], "path", path, "err", err)
		http.Error(w, "Invalid level", http.StatusBadRequest)
		return
	}

	// Парсим col_row.format
	colRowFormat := strings.Split(tileParts[1], ".")
	if len(colRowFormat) != 2 {
		slog.Warn("GetTile: invalid format", "tile_part", tileParts[1], "path", path)
		http.Error(w, "Invalid tile format", http.StatusBadRequest)
		return
	}

	format := colRowFormat[1]
	colRow := strings.Split(colRowFormat[0], "_")
	if len(colRow) != 2 {
		slog.Warn("GetTile: invalid coordinates", "col_row_str", colRowFormat[0], "path", path)
		http.Error(w, "Invalid tile coordinates", http.StatusBadRequest)
		return
	}

	col, err := strconv.Atoi(colRow[0])
	if err != nil {
		slog.Warn("GetTile: invalid column", "col_str", colRow[0], "path", path, "err", err)
		http.Error(w, "Invalid column", http.StatusBadRequest)
		return
	}

	row, err := strconv.Atoi(colRow[1])
	if err != nil {
		slog.Warn("GetTile: invalid row", "row_str", colRow[1], "path", path, "err", err)
		http.Error(w, "Invalid row", http.StatusBadRequest)
		return
	}

	slog.Info("GetTile: request received",
		"method", r.Method,
		"path", r.URL.Path,
		"decoded_file_path", decodedFilePath,
		"level", level,
		"col", col,
		"row", row,
		"format", format,
		"remote_addr", r.RemoteAddr,
		"user_agent", r.UserAgent(),
	)

	ctx := r.Context()
	tile, err := h.imageService.GetTile(ctx, decodedFilePath, level, col, row, format)
	duration := time.Since(startTime)

	if err != nil {
		slog.Error("GetTile: failed to get tile",
			"decoded_file_path", decodedFilePath,
			"level", level,
			"col", col,
			"row", row,
			"format", format,
			"err", err,
			"duration_ms", duration.Milliseconds(),
		)
		http.Error(w, fmt.Sprintf("Failed to get tile: %v", err), http.StatusInternalServerError)
		return
	}

	// Устанавливаем Content-Type
	var contentType string
	switch strings.ToLower(format) {
	case "jpeg", "jpg":
		contentType = "image/jpeg"
	case "png":
		contentType = "image/png"
	default:
		contentType = "image/jpeg"
	}

	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(http.StatusOK)
	w.Write(tile.Data)

	slog.Info("GetTile: success",
		"decoded_file_path", decodedFilePath,
		"level", level,
		"col", col,
		"row", row,
		"format", format,
		"tile_size", len(tile.Data),
		"duration_ms", duration.Milliseconds(),
	)
}
