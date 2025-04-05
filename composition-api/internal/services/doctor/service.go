package doctor

import (
	"context"

	"composition-api/internal/adapters"
	domain "composition-api/internal/domain/med"

	"github.com/google/uuid"
)

type Service interface {
	GetDoctor(ctx context.Context, id uuid.UUID) (domain.Doctor, error)
}

type service struct {
	adapters *adapters.Adapters
}

func New(
	adapters *adapters.Adapters,
) Service {
	return &service{
		adapters: adapters,
	}
}
