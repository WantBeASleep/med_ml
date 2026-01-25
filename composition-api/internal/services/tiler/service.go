package tiler

import (
	"bytes"
	"context"
	"io"

	"composition-api/internal/adapters/tiler"
)

type Service interface {
	GetDZI(ctx context.Context, filePath string) (string, error)
	GetTile(ctx context.Context, filePath string, level, col, row int, format string) (io.ReadCloser, error)
}

type service struct {
	client tiler.Client
}

func New(client tiler.Client) Service {
	return &service{
		client: client,
	}
}

func (s *service) GetDZI(ctx context.Context, filePath string) (string, error) {
	return s.client.GetDZI(ctx, filePath)
}

func (s *service) GetTile(ctx context.Context, filePath string, level, col, row int, format string) (io.ReadCloser, error) {
	data, err := s.client.GetTile(ctx, filePath, level, col, row, format)
	if err != nil {
		return nil, err
	}

	return io.NopCloser(bytes.NewReader(data)), nil
}
