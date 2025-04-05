package download

import (
	"context"
	"io"

	"composition-api/internal/repository"

	"github.com/google/uuid"
)

type Service interface {
	GetImage(ctx context.Context, uziID uuid.UUID, imageID uuid.UUID) (io.ReadCloser, error)
}

type service struct {
	repo repository.DAO
}

func New(
	repo repository.DAO,
) Service {
	return &service{
		repo: repo,
	}
}
