package user

import (
	"context"
	"errors"
	"fmt"

	"auth/internal/domain"

	"github.com/google/uuid"

	"auth/internal/repository/entity"
	uentity "auth/internal/repository/user/entity"
)

func (s *service) RegisterUser(
	ctx context.Context,
	email string,
	password string,
	role domain.Role,
) (uuid.UUID, error) {
	if email == "" || password == "" {
		return uuid.Nil, domain.ErrBadRequest
	}

	pass, err := s.passwordSrv.CreatePassword(password)
	if err != nil {
		return uuid.Nil, fmt.Errorf("create password: %w", err)
	}

	userRepo := s.dao.NewUserRepo(ctx)
	userDB, err := userRepo.GetUserByEmail(email)
	switch {
	case err == nil:
		user := userDB.ToDomain()
		if user.Password != nil {
			return uuid.Nil, domain.ErrConflict
		}
		if user.Role != role {
			return uuid.Nil, domain.ErrUnprocessableEntity
		}
		user.Password = &pass
		if err := userRepo.UpdateUserPassword(user.Id, pass.String()); err != nil {
			return uuid.Nil, fmt.Errorf("update user password: %w", err)
		}
		return user.Id, nil

	case errors.Is(err, entity.ErrNotFound):
		user := domain.User{
			Id:       uuid.New(),
			Email:    email,
			Password: &pass,
			Role:     role,
		}
		if err := userRepo.InsertUser(uentity.User{}.FromDomain(user)); err != nil {
			var dbErr *entity.DBConflictError
			if errors.As(err, &dbErr) {
				return uuid.Nil, domain.ErrConflict
			}
			var valErr *entity.DBValidationError
			if errors.As(err, &valErr) {
				return uuid.Nil, domain.ErrUnprocessableEntity
			}
			return uuid.Nil, fmt.Errorf("create user: %w", err)
		}

		return user.Id, nil
	default:
		return uuid.Nil, fmt.Errorf("get user by email: %w", err)
	}
}

func (s *service) CreateUnRegisteredUser(
	ctx context.Context,
	email string,
) (uuid.UUID, error) {
	user := domain.User{
		Id:    uuid.New(),
		Email: email,
		Role:  domain.RolePatient,
	}

	if err := s.dao.NewUserRepo(ctx).InsertUser(uentity.User{}.FromDomain(user)); err != nil {
		var dbErr *entity.DBConflictError
		if errors.As(err, &dbErr) {
			return uuid.Nil, domain.ErrConflict
		}
		return uuid.Nil, fmt.Errorf("create user: %w", err)
	}

	return user.Id, nil
}
