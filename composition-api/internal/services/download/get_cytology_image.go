package download

import (
	"context"
	"io"

	"github.com/google/uuid"
)

func (s *service) GetCytologyImage(ctx context.Context, cytologyID uuid.UUID, originalImageID uuid.UUID) (io.ReadCloser, error) {
	// Получаем original_image из БД, чтобы использовать правильный путь из ImagePath
	originalImage, err := s.cytologyService.GetOriginalImageById(ctx, originalImageID)
	if err != nil {
		return nil, err
	}

	// Используем путь из БД, который соответствует пути в S3
	if originalImage.ImagePath == "" {
		return nil, io.ErrUnexpectedEOF
	}

	return s.repo.NewFileRepo().GetFile(ctx, originalImage.ImagePath)
}
