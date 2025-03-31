package password

import (
	"auth/internal/domain"
)

type Service interface {
	CreatePassword(pass string) (domain.Password, error)

	ComparePassword(cleanPass string, hashedPass domain.Password) bool
}

type service struct{}

func New() Service {
	return &service{}
}
