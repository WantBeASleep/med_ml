package flow

import (
	"context"
	"fmt"
	"log/slog"

	"auth/internal/domain"
	pb "auth/internal/generated/grpc/service"
	"auth/internal/server/mappers"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
)

var RegisterUser flowfuncDepsInjector = func(deps *Deps) flowfunc {
	return func(ctx context.Context, data FlowData) (FlowData, error) {
		email := gofakeit.Email()
		password := gofakeit.Password(true, true, true, true, false, 10)
		role := domain.RoleDoctor

		resp, err := deps.Adapter.RegisterUser(ctx, &pb.RegisterUserIn{
			Email:    email,
			Password: password,
			Role:     mappers.RoleMap[role],
		})
		if err != nil {
			slog.ErrorContext(ctx, "register user", slog.Any("err", err))
			return FlowData{}, fmt.Errorf("register user: %w", err)
		}

		data.RegisterUser = FlowUser{
			Id:       uuid.MustParse(resp.Id),
			Email:    email,
			Password: password,
			Role:     role,
		}
		return data, nil
	}
}
