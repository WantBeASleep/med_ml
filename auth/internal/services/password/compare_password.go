package password

import (
	"auth/internal/domain"
)

func (s *service) ComparePassword(cleanPass string, hashedPass domain.Password) bool {
	hash, err := s.hashPassword(cleanPass, hashedPass.Salt)
	if err != nil {
		return false
	}

	return hash == hashedPass.Hash
}
