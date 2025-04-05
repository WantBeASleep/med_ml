package token

import (
	"crypto/rsa"
	"time"

	"auth/internal/domain"

	"github.com/google/uuid"
)

type Service interface {
	// return access and refresh tokens
	GenerateUserTokens(id uuid.UUID, role domain.Role) (domain.Token, domain.Token, error)
	ParseUserToken(token domain.Token) (uuid.UUID, domain.Role, error)
	ValidateToken(token domain.Token) bool
}

type service struct {
	accessLifeTime  time.Duration
	refreshLifeTime time.Duration
	privateKey      *rsa.PrivateKey
	publicKey       *rsa.PublicKey
}

func New(
	accessLifeTime time.Duration,
	refreshLifeTime time.Duration,
	privateKey *rsa.PrivateKey,
	publicKey *rsa.PublicKey,
) Service {
	return &service{
		accessLifeTime:  accessLifeTime,
		refreshLifeTime: refreshLifeTime,
		privateKey:      privateKey,
		publicKey:       publicKey,
	}
}
