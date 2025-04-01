package auth

import (
	"context"
	"errors"
	"fmt"

	"auth/internal/domain"
	uentity "auth/internal/repository/user/entity"
)

func (s *service) Refresh(ctx context.Context, refreshToken domain.Token) (domain.Token, domain.Token, error) {
	userID, _, err := s.tokenSrv.ParseUserToken(refreshToken)
	if err != nil {
		return "", "", fmt.Errorf("parse token: %w", err)
	}

	userRepo := s.dao.NewUserRepo(ctx)
	userDB, err := userRepo.GetUserByID(userID)
	if err != nil {
		return "", "", fmt.Errorf("get user by pk: %w", err)
	}
	user := userDB.ToDomain()

	if refreshToken != *user.RefreshToken {
		return "", "", errors.New("tokens not equal")
	}

	access, refresh, err := s.tokenSrv.GenerateUserTokens(user.Id, user.Role)
	if err != nil {
		return "", "", fmt.Errorf("generate tokens pair: %w", err)
	}

	user.RefreshToken = &refresh
	if err := userRepo.UpdateUser(uentity.User{}.FromDomain(user)); err != nil {
		return "", "", fmt.Errorf("update user: %w", err)
	}

	return access, refresh, nil
}
