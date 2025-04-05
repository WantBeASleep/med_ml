package tokens

import (
	"context"

	domain "composition-api/internal/domain/auth"
)

func (s *service) Refresh(ctx context.Context, refreshToken domain.Token) (domain.Token, domain.Token, error) {
	return s.adapters.Auth.Refresh(ctx, refreshToken)
}
