package register

import (
	"context"
	"errors"

	"auth/internal/domain"
	pb "auth/internal/generated/grpc/service"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) CreateUnRegisteredUser(ctx context.Context, in *pb.CreateUnRegisteredUserIn) (*pb.CreateUnRegisteredUserOut, error) {
	id, err := h.services.UserService.CreateUnRegisteredUser(ctx, in.Email)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrConflict):
			return nil, status.Errorf(codes.AlreadyExists, "Пользователь с таким email уже существует")
		default:
			return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
		}
	}

	return &pb.CreateUnRegisteredUserOut{Id: id.String()}, nil
}
