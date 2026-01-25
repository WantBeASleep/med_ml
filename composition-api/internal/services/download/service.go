package download

import (
	"context"
	"io"

	"composition-api/internal/repository"
	"composition-api/internal/services/cytology"

	"github.com/google/uuid"
)

type Service interface {
	GetImage(ctx context.Context, uziID uuid.UUID, imageID uuid.UUID) (io.ReadCloser, error)
	GetCytologyImage(ctx context.Context, cytologyID uuid.UUID, originalImageID uuid.UUID) (io.ReadCloser, error)
}

type service struct {
	repo            repository.DAO
	cytologyService cytology.Service
}

func New(
	repo repository.DAO,
	cytologyService cytology.Service,
) Service {
	return &service{
		repo:            repo,
		cytologyService: cytologyService,
	}
}
