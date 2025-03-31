package auth

import (
	"context"

	pb "auth/internal/generated/grpc/service"
	"auth/internal/services"
)

type AuthHandler interface {
	Login(ctx context.Context, in *pb.LoginIn) (*pb.LoginOut, error)
	Refresh(ctx context.Context, in *pb.RefreshIn) (*pb.RefreshOut, error)
}

type handler struct {
	services *services.Services
}

func New(
	services *services.Services,
) AuthHandler {
	return &handler{
		services: services,
	}
}
