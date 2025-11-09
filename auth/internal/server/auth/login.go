package auth

import (
	"context"
	"errors"

	"auth/internal/domain"
	pb "auth/internal/generated/grpc/service"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) Login(ctx context.Context, in *pb.LoginIn) (*pb.LoginOut, error) {
	access, refresh, err := h.services.AuthService.Login(ctx, in.Email, in.Password)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrUnauthorized):
			return nil, status.Errorf(codes.Unauthenticated, "Неверный email или пароль")
		default:
			return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
		}
	}

	return &pb.LoginOut{
		AccessToken:  access.String(),
		RefreshToken: refresh.String(),
	}, nil
}
