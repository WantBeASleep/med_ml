package flow

import (
	"context"
	"fmt"

	pb "auth/internal/generated/grpc/service"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
)

var CreateUnRegisterUser flowfuncDepsInjector = func(deps *Deps) flowfunc {
	return func(ctx context.Context, data FlowData) (FlowData, error) {
		data = FlowData{}

		email := gofakeit.Email()

		resp, err := deps.Adapter.CreateUnRegisteredUser(ctx, &pb.CreateUnRegisteredUserIn{Email: email})
		if err != nil {
			return FlowData{}, fmt.Errorf("create unregistered user: %w", err)
		}

		data.UnRegisterUser = FlowUserUnRegister{
			Id:    uuid.MustParse(resp.Id),
			Email: email,
		}
		return data, nil
	}
}
