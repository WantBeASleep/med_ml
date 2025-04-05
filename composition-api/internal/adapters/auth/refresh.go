package auth

import (
	"context"

	domain "composition-api/internal/domain/auth"
	pb "composition-api/internal/generated/grpc/clients/auth"
)

func (a *adapter) Refresh(ctx context.Context, refreshToken domain.Token) (domain.Token, domain.Token, error) {
	res, err := a.client.Refresh(ctx, &pb.RefreshIn{
		RefreshToken: refreshToken.String(),
	})
	if err != nil {
		return domain.Token(""), domain.Token(""), err
	}

	return domain.Token(res.AccessToken), domain.Token(res.RefreshToken), nil
}
