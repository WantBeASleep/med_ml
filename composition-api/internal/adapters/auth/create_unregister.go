package auth

import (
	"context"

	pb "composition-api/internal/generated/grpc/clients/auth"

	"github.com/google/uuid"
)

func (a *adapter) CreateUnRegisteredUser(ctx context.Context, email string) (uuid.UUID, error) {
	res, err := a.client.CreateUnRegisteredUser(ctx, &pb.CreateUnRegisteredUserIn{
		Email: email,
	})
	if err != nil {
		return uuid.Nil, err
	}

	return uuid.MustParse(res.Id), nil
}
