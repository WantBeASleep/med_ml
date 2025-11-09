package auth

import (
	"context"

	adapter_errors "composition-api/internal/adapters/errors"
	domain "composition-api/internal/domain/auth"
	pb "composition-api/internal/generated/grpc/clients/auth"

	"github.com/google/uuid"
)

var roleMap = map[domain.Role]pb.Role{
	domain.RoleDoctor:  pb.Role_ROLE_DOCTOR,
	domain.RolePatient: pb.Role_ROLE_PATIENT,
}

func (a *adapter) RegisterUser(ctx context.Context, email, password string, role domain.Role) (uuid.UUID, error) {
	res, err := a.client.RegisterUser(ctx, &pb.RegisterUserIn{
		Email:    email,
		Password: password,
		Role:     roleMap[role],
	})
	if err != nil {
		return uuid.Nil, adapter_errors.HandleGRPCError(err)
	}

	return uuid.MustParse(res.Id), nil
}
