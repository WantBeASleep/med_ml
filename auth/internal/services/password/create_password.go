package password

import (
	"auth/internal/domain"
)

func (s *service) CreatePassword(pass string) (domain.Password, error) {
	salt := s.generateSalt()

	hash, err := s.hashPassword(pass, salt)
	if err != nil {
		return domain.Password{}, err
	}

	return domain.Password{Hash: hash, Salt: salt}, nil
}
