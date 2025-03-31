package auth

import (
	"context"

	pb "auth/internal/generated/grpc/service"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) Login(ctx context.Context, in *pb.LoginIn) (*pb.LoginOut, error) {
	access, refresh, err := h.services.AuthService.Login(ctx, in.Email, in.Password)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}

	return &pb.LoginOut{
		AccessToken:  access.String(),
		RefreshToken: refresh.String(),
	}, nil
}
