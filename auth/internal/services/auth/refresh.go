package auth

import (
	"context"
	"fmt"

	"auth/internal/domain"
	rtentity "auth/internal/repository/refresh_token/entity"
)

// TODO: неактуальный rt надо удалять, реализовать, когда будет сделан крон удаления устаревших rt.
func (s *service) Refresh(ctx context.Context, refreshToken domain.Token) (domain.Token, domain.Token, error) {
	userID, _, err := s.tokenSrv.ParseUserToken(refreshToken)
	if err != nil {
		return "", "", domain.ErrBadRequest
	}

	refreshTokenRepo := s.dao.NewRefreshTokenRepo(ctx)
	exists, err := refreshTokenRepo.ExistsRefreshToken(userID, refreshToken.String())
	if err != nil {
		return "", "", fmt.Errorf("check exists refresh token: %w", err)
	}
	if !exists {
		return "", "", domain.ErrUnauthorized
	}

	// итого сюда нужно транзакцию, но не критично если ее не будет
	if err := refreshTokenRepo.DeleteRefreshTokens([]rtentity.RefreshToken{
		{
			Id:           userID,
			RefreshToken: refreshToken.String(),
		},
	}); err != nil {
		return "", "", fmt.Errorf("delete refresh token: %w", err)
	}

	userRepo := s.dao.NewUserRepo(ctx)
	userDB, err := userRepo.GetUserByID(userID)
	if err != nil {
		return "", "", fmt.Errorf("get user by pk: %w", err)
	}
	user := userDB.ToDomain()

	access, refresh, err := s.tokenSrv.GenerateUserTokens(user.Id, user.Role)
	if err != nil {
		return "", "", fmt.Errorf("generate tokens pair: %w", err)
	}

	if err := refreshTokenRepo.InsertRefreshToken(rtentity.RefreshToken{
		Id:           user.Id,
		RefreshToken: refresh.String(),
	}); err != nil {
		return "", "", fmt.Errorf("insert new refresh token: %w", err)
	}

	return access, refresh, nil
}
