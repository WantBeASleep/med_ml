package token

import (
	"fmt"
	"time"

	"auth/internal/domain"

	"github.com/google/uuid"
)

func (s *service) GenerateUserTokens(id uuid.UUID, role domain.Role) (domain.Token, domain.Token, error) {
	claims := map[string]any{
		"id":   id.String(),
		"role": role.String(),
	}

	accessToken, err := s.generateToken(claims, WithExpirationTime(time.Now().Add(s.accessLifeTime)))
	if err != nil {
		return "", "", fmt.Errorf("generate access token: %w", err)
	}

	refreshToken, err := s.generateToken(claims, WithExpirationTime(time.Now().Add(s.refreshLifeTime)))
	if err != nil {
		return "", "", fmt.Errorf("generate refresh token: %w", err)
	}

	return domain.Token(accessToken), domain.Token(refreshToken), nil
}
