package tiler

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Client interface {
	GetDZI(ctx context.Context, filePath string) (string, error)
	GetTile(ctx context.Context, filePath string, level, col, row int, format string) ([]byte, error)
}

type client struct {
	baseURL    string
	httpClient *http.Client
}

func NewClient(baseURL string) Client {
	// Добавляем протокол, если его нет
	url := strings.TrimSpace(baseURL)
	if url != "" && !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "http://" + url
	}

	return &client{
		baseURL:    strings.TrimSuffix(url, "/"),
		httpClient: &http.Client{},
	}
}

func (c *client) GetDZI(ctx context.Context, filePath string) (string, error) {
	startTime := time.Now()

	// Формируем URL для DZI
	u, err := url.Parse(c.baseURL)
	if err != nil {
		slog.Error("tiler client: failed to parse base URL", "base_url", c.baseURL, "err", err)
		return "", err
	}
	// Убираем ведущий слэш из filePath, если он есть, чтобы path.Join работал правильно
	filePath = strings.TrimPrefix(filePath, "/")
	u.Path = "/dzi/" + filePath

	requestURL := u.String()
	slog.Info("tiler client: GetDZI request",
		"url", requestURL,
		"file_path", filePath,
	)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL, nil)
	if err != nil {
		slog.Error("tiler client: failed to create request", "url", requestURL, "err", err)
		return "", err
	}

	resp, err := c.httpClient.Do(req)
	duration := time.Since(startTime)
	if err != nil {
		slog.Error("tiler client: GetDZI request failed",
			"url", requestURL,
			"file_path", filePath,
			"err", err,
			"duration_ms", duration.Milliseconds(),
		)
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		slog.Warn("tiler client: GetDZI unexpected status",
			"url", requestURL,
			"file_path", filePath,
			"status_code", resp.StatusCode,
			"duration_ms", duration.Milliseconds(),
		)
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error("tiler client: GetDZI failed to read body",
			"url", requestURL,
			"file_path", filePath,
			"err", err,
			"duration_ms", duration.Milliseconds(),
		)
		return "", err
	}

	slog.Info("tiler client: GetDZI success",
		"url", requestURL,
		"file_path", filePath,
		"body_size", len(body),
		"duration_ms", duration.Milliseconds(),
	)

	return string(body), nil
}

func (c *client) GetTile(ctx context.Context, filePath string, level, col, row int, format string) ([]byte, error) {
	startTime := time.Now()

	// Формируем URL для тайла
	// Формат: /dzi/{file_path}/files/{level}/{col}_{row}.{format}
	// Убираем ведущий слэш из filePath, если он есть
	filePath = strings.TrimPrefix(filePath, "/")
	tilePath := filePath + "/files/" + fmt.Sprintf("%d/%d_%d.%s", level, col, row, format)

	u, err := url.Parse(c.baseURL)
	if err != nil {
		slog.Error("tiler client: failed to parse base URL", "base_url", c.baseURL, "err", err)
		return nil, err
	}
	u.Path = "/dzi/" + tilePath

	requestURL := u.String()
	slog.Info("tiler client: GetTile request",
		"url", requestURL,
		"file_path", filePath,
		"level", level,
		"col", col,
		"row", row,
		"format", format,
	)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL, nil)
	if err != nil {
		slog.Error("tiler client: failed to create request", "url", requestURL, "err", err)
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	duration := time.Since(startTime)
	if err != nil {
		slog.Error("tiler client: GetTile request failed",
			"url", requestURL,
			"file_path", filePath,
			"level", level,
			"col", col,
			"row", row,
			"format", format,
			"err", err,
			"duration_ms", duration.Milliseconds(),
		)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		slog.Warn("tiler client: GetTile unexpected status",
			"url", requestURL,
			"file_path", filePath,
			"level", level,
			"col", col,
			"row", row,
			"format", format,
			"status_code", resp.StatusCode,
			"duration_ms", duration.Milliseconds(),
		)
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error("tiler client: GetTile failed to read body",
			"url", requestURL,
			"file_path", filePath,
			"level", level,
			"col", col,
			"row", row,
			"format", format,
			"err", err,
			"duration_ms", duration.Milliseconds(),
		)
		return nil, err
	}

	slog.Info("tiler client: GetTile success",
		"url", requestURL,
		"file_path", filePath,
		"level", level,
		"col", col,
		"row", row,
		"format", format,
		"tile_size", len(body),
		"duration_ms", duration.Milliseconds(),
	)

	return body, nil
}
