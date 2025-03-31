package token

import (
	"fmt"

	"github.com/google/uuid"

	"auth/internal/domain"
)

func (s *service) ParseUserToken(token domain.Token) (uuid.UUID, domain.Role, error) {
	claims, err := s.parseClaims(token.String())
	if err != nil {
		return uuid.Nil, "", fmt.Errorf("parse claims: %w", err)
	}

	unpackedClaims := map[string]string{}
	for k, v := range claims {
		unpacked, ok := v.(string)
		if !ok {
			return uuid.Nil, "", fmt.Errorf("invalid claim type: %T", v)
		}

		unpackedClaims[k] = unpacked
	}

	userId, ok := unpackedClaims["user_id"]
	if !ok {
		return uuid.Nil, "", fmt.Errorf("invalid user id")
	}

	roleParsed, ok := unpackedClaims["role"]
	if !ok {
		return uuid.Nil, "", fmt.Errorf("invalid role")
	}

	id, err := uuid.Parse(userId)
	if err != nil {
		return uuid.Nil, "", fmt.Errorf("parse user id: %w", err)
	}

	role, err := domain.Role.Parse("", roleParsed)
	if err != nil {
		return uuid.Nil, "", fmt.Errorf("parse role: %w", err)
	}

	return id, role, nil
}
