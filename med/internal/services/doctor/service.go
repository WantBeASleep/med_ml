package doctor

import (
	"context"

	"med/internal/domain"
	"med/internal/repository"

	"github.com/google/uuid"
)

type Service interface {
	RegisterDoctor(ctx context.Context, doctor domain.Doctor) error

	GetDoctor(ctx context.Context, id uuid.UUID) (domain.Doctor, error)
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
