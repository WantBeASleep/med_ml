package tokens

import (
	"context"

	domain "composition-api/internal/domain/auth"
)

func (s *service) Login(ctx context.Context, email, password string) (domain.Token, domain.Token, error) {
	return s.adapters.Auth.Login(ctx, email, password)
}
