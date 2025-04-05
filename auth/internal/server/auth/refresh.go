package auth

import (
	"context"

	"auth/internal/domain"
	pb "auth/internal/generated/grpc/service"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) Refresh(ctx context.Context, in *pb.RefreshIn) (*pb.RefreshOut, error) {
	access, refresh, err := h.services.AuthService.Refresh(ctx, domain.Token(in.RefreshToken))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}

	return &pb.RefreshOut{
		AccessToken:  access.String(),
		RefreshToken: refresh.String(),
	}, nil
}
