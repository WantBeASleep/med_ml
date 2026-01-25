package card

import (
	"context"

	"med/internal/domain"
	"med/internal/repository"

	"github.com/google/uuid"
)

type Service interface {
	CreateCard(ctx context.Context, card domain.Card) (domain.Card, error)

	GetCard(ctx context.Context, doctorID, patientID uuid.UUID) (domain.Card, error)

	UpdateCard(ctx context.Context, doctorID, patientID uuid.UUID, update UpdateCardArg) (domain.Card, error)
}

type service struct {
	dao repository.DAO
}

func New(
	dao repository.DAO,
) Service {
	return &service{
		dao: dao,
	}
}
