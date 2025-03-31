package auth

import (
	"context"

	"auth/internal/domain"
	"auth/internal/repository"
	"auth/internal/services/password"
	"auth/internal/services/token"
)

type Service interface {
	Login(ctx context.Context, email, password string) (domain.Token, domain.Token, error)
	Refresh(ctx context.Context, refreshToken domain.Token) (domain.Token, domain.Token, error)
}

type service struct {
	dao         repository.DAO
	passwordSrv password.Service
	tokenSrv    token.Service
}

func New(
	dao repository.DAO,
	passwordSrv password.Service,
	tokenSrv token.Service,
) Service {
	return &service{
		dao:         dao,
		passwordSrv: passwordSrv,
		tokenSrv:    tokenSrv,
	}
}
