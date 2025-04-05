package token

import (
	"fmt"

	"github.com/google/uuid"

	"auth/internal/domain"
)

func (s *service) ParseUserToken(token domain.Token) (uuid.UUID, domain.Role, error) {
	// TODO: логика валидации вросла в parseClaims, исправить
	claims, err := s.parseClaims(token.String())
	if err != nil {
		return uuid.Nil, "", fmt.Errorf("parse claims: %w", err)
	}

	userId, ok := claims["id"]
	if !ok {
		return uuid.Nil, "", fmt.Errorf("user id not found")
	}

	roleParsed, ok := claims["role"]
	if !ok {
		return uuid.Nil, "", fmt.Errorf("role not found")
	}

	id, err := uuid.Parse(userId.(string))
	if err != nil {
		return uuid.Nil, "", fmt.Errorf("parse user id: %w", err)
	}

	role, err := domain.Role.Parse("", roleParsed.(string))
	if err != nil {
		return uuid.Nil, "", fmt.Errorf("parse role: %w", err)
	}

	return id, role, nil
}
