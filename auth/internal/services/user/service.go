package user

import (
	"context"

	"auth/internal/domain"
	"auth/internal/repository"
	"auth/internal/services/password"

	"github.com/google/uuid"
)

type Service interface {
	RegisterUser(ctx context.Context, email, password string, role domain.Role) (uuid.UUID, error)
	// Пользователь не зарегистрирован в системе, но сведения о его почте сохраняются
	//
	// Кейс: на незареганного пациента(почту) создаются узи, при регистрации пользователя
	// на эту почту, все узи будут отображаться
	CreateUnRegisteredUser(ctx context.Context, email string) (uuid.UUID, error)
}

type service struct {
	dao         repository.DAO
	passwordSrv password.Service
}

func New(
	dao repository.DAO,
	passwordSrv password.Service,
) Service {
	return &service{
		dao:         dao,
		passwordSrv: passwordSrv,
	}
}
