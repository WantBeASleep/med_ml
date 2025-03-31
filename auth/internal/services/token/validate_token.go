package token

import (
	"auth/internal/domain"
)

func (s *service) ValidateToken(token domain.Token) bool {
	_, err := s.parseClaims(token.String())
	return err == nil
}
