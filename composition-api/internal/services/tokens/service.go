package tokens

import (
	"context"

	"composition-api/internal/adapters"
	domain "composition-api/internal/domain/auth"
)

type Service interface {
	Login(ctx context.Context, email, password string) (domain.Token, domain.Token, error)
	Refresh(ctx context.Context, refreshToken domain.Token) (domain.Token, domain.Token, error)
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
