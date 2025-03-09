package node

import (
	"context"

	"uzi/internal/domain"
	"uzi/internal/repository"

	"github.com/google/uuid"
)

type Service interface {
	GetNodesByUziID(ctx context.Context, id uuid.UUID) ([]domain.Node, error)

	UpdateNode(ctx context.Context, arg UpdateNodeArg) (domain.Node, error)
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
