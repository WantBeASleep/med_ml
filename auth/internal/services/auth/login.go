package auth

import (
	"context"
	"errors"
	"fmt"

	"auth/internal/domain"
	uentity "auth/internal/repository/user/entity"
)

var (
	ErrUserNotRegistered = errors.New("user not registered")
	ErrWrongPassword     = errors.New("wrong password")
)

func (s *service) Login(ctx context.Context, email, password string) (domain.Token, domain.Token, error) {
	userRepo := s.dao.NewUserRepo(ctx)
	userDB, err := userRepo.GetUserByEmail(email)
	if err != nil {
		return "", "", fmt.Errorf("get user by email: %w", err)
	}
	user := userDB.ToDomain()

	if user.Password == nil {
		return "", "", ErrUserNotRegistered
	}

	if !s.passwordSrv.ComparePassword(password, *user.Password) {
		return "", "", ErrWrongPassword
	}

	accessToken, refreshToken, err := s.tokenSrv.GenerateUserTokens(user.Id, user.Role)
	if err != nil {
		return "", "", fmt.Errorf("generate tokens: %w", err)
	}

	user.RefreshToken = &refreshToken
	if err := userRepo.UpdateUser(uentity.User{}.FromDomain(user)); err != nil {
		return "", "", fmt.Errorf("update user: %w", err)
	}

	return accessToken, refreshToken, nil
}
