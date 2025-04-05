package register

import (
	"context"

	pb "auth/internal/generated/grpc/service"
	"auth/internal/server/mappers"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) RegisterUser(ctx context.Context, in *pb.RegisterUserIn) (*pb.RegisterUserOut, error) {
	id, err := h.services.UserService.RegisterUser(
		ctx,
		in.Email,
		in.Password,
		mappers.RoleReversedMap[in.Role],
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}

	return &pb.RegisterUserOut{Id: id.String()}, nil
}
