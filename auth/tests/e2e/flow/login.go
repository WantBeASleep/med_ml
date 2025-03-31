package flow

import (
	"context"
	"fmt"

	"auth/internal/domain"
	pb "auth/internal/generated/grpc/service"
)

var Login flowfuncDepsInjector = func(deps *Deps) flowfunc {
	return func(ctx context.Context, data FlowData) (FlowData, error) {
		data = FlowData{}

		resp, err := deps.Adapter.Login(ctx, &pb.LoginIn{
			Email:    data.RegisterUser.Email,
			Password: data.RegisterUser.Password,
		})
		if err != nil {
			return FlowData{}, fmt.Errorf("login: %w", err)
		}

		data.Tokens = RegisterUserTokens{
			Access:  domain.Token(resp.AccessToken),
			Refresh: domain.Token(resp.RefreshToken),
		}

		return data, nil
	}
}
