package auth

import (
	"context"

	domain "composition-api/internal/domain/auth"
	pb "composition-api/internal/generated/grpc/clients/auth"
)

func (a *adapter) Login(ctx context.Context, email, password string) (domain.Token, domain.Token, error) {
	res, err := a.client.Login(ctx, &pb.LoginIn{
		Email:    email,
		Password: password,
	})
	if err != nil {
		return domain.Token(""), domain.Token(""), err
	}

	return domain.Token(res.AccessToken), domain.Token(res.RefreshToken), nil
}
