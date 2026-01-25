package card

import (
	"context"

	"composition-api/internal/adapters"
	domain "composition-api/internal/domain/med"

	"github.com/google/uuid"
)

type Service interface {
	CreateCard(ctx context.Context, card domain.Card) (domain.Card, error)

	GetCard(ctx context.Context, doctorID, patientID uuid.UUID) (domain.Card, error)

	UpdateCard(ctx context.Context, card domain.Card) (domain.Card, error)
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
