package node

import (
	"context"

	"github.com/google/uuid"

	"gateway/internal/adapters"
	domain "gateway/internal/domain/uzi"
)

type Service interface {
	GetNodesByUziID(ctx context.Context, uziID uuid.UUID) ([]domain.Node, error)
	UpdateNode(ctx context.Context, arg UpdateNodeArg) (domain.Node, error)
	DeleteNode(ctx context.Context, id uuid.UUID) error
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
