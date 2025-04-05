// TODO: плохо расположен пакет, но Я ТАК УСТААЛ УЖЕ, ПРОСТО ПОТОМ mv *
package auth

import (
	"context"

	domain "composition-api/internal/domain/auth"
	"composition-api/internal/server/security"
)

func IsDoctor(ctx context.Context) (bool, error) {
	token, err := security.ParseToken(ctx)
	if err != nil {
		return false, err
	}

	return token.Role == domain.RoleDoctor, nil
}

func IsPatient(ctx context.Context) (bool, error) {
	token, err := security.ParseToken(ctx)
	if err != nil {
		return false, err
	}

	return token.Role == domain.RolePatient, nil
}
