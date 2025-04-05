package security

import (
	"context"
	"crypto/rsa"
	"errors"
	"fmt"

	domain "composition-api/internal/domain/auth"
	api "composition-api/internal/generated/http/api"

	"composition-api/internal/config"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrUnauthorized = errors.New("unauthorized")
)

type handler struct {
	publicKey *rsa.PublicKey
}

func New(cfg *config.Config) *handler {
	publicKey, err := cfg.ParseRsaPublicKey()
	if err != nil {
		panic(err)
	}

	return &handler{publicKey: publicKey}
}

func (h *handler) HandleBearerAuth(ctx context.Context, operationName api.OperationName, t api.BearerAuth) (context.Context, error) {
	parsed, err := jwt.Parse(t.Token, func(t *jwt.Token) (interface{}, error) { return h.publicKey, nil })
	if err != nil {
		return nil, fmt.Errorf("parse token: %w", err)
	}
	if !parsed.Valid {
		return nil, ErrInvalidToken
	}

	claims, ok := parsed.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid jwt claims")
	}

	idAny, ok := claims["id"]
	if !ok {
		return nil, ErrInvalidToken
	}
	idString, ok := idAny.(string)
	if !ok {
		return nil, ErrInvalidToken
	}
	id, err := uuid.Parse(idString)
	if err != nil {
		return nil, ErrInvalidToken
	}

	roleAny, ok := claims["role"]
	if !ok {
		return nil, ErrInvalidToken
	}
	roleString, ok := roleAny.(string)
	if !ok {
		return nil, ErrInvalidToken
	}
	role, err := domain.Role.Parse("", roleString)
	if err != nil {
		return nil, ErrInvalidToken
	}

	return context.WithValue(ctx, tokenKey, Token{Id: id, Role: role}), nil
}
