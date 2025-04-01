package user

import (
	"context"
	"fmt"

	"auth/internal/domain"

	"github.com/google/uuid"

	uentity "auth/internal/repository/user/entity"
)

func (s *service) RegisterUser(
	ctx context.Context,
	email string,
	password string,
	role domain.Role,
) (uuid.UUID, error) {
	return s.createUser(ctx, email, role, WithPassword(password))
}

func (s *service) CreateUnRegisteredUser(
	ctx context.Context,
	email string,
) (uuid.UUID, error) {
	return s.createUser(ctx, email, domain.RolePatient)
}

func (s *service) createUser(
	ctx context.Context,
	email string,
	role domain.Role,
	opts ...registerOption,
) (uuid.UUID, error) {

	// TODO: логика перерга юзера зареганого по лайте

	options := &registerOptions{}
	for _, opt := range opts {
		opt(options)
	}

	user := domain.User{
		Email: email,
		Role:  role,
	}

	id := uuid.New()
	user.Id = id

	if options.password != nil {
		pass, err := s.passwordSrv.CreatePassword(*options.password)
		if err != nil {
			return uuid.Nil, fmt.Errorf("create password: %w", err)
		}
		user.Password = &pass
	}

	userRepo := s.dao.NewUserRepo(ctx)
	if err := userRepo.InsertUser(uentity.User{}.FromDomain(user)); err != nil {
		return uuid.Nil, fmt.Errorf("create user: %w", err)
	}

	return id, nil
}
