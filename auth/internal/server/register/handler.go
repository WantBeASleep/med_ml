package register

import (
	"context"

	pb "auth/internal/generated/grpc/service"
	"auth/internal/services"
)

type RegisterHandler interface {
	RegisterUser(ctx context.Context, in *pb.RegisterUserIn) (*pb.RegisterUserOut, error)
	CreateUnRegisteredUser(ctx context.Context, in *pb.CreateUnRegisteredUserIn) (*pb.CreateUnRegisteredUserOut, error)
}

type handler struct {
	services *services.Services
}

func New(
	services *services.Services,
) RegisterHandler {
	return &handler{
		services: services,
	}
}
