package image

import (
	"context"

	pb "uzi/internal/generated/grpc/service"
	"uzi/internal/services"
)

type ImageHandler interface {
	GetImagesByUziId(ctx context.Context, in *pb.GetImagesByUziIdIn) (*pb.GetImagesByUziIdOut, error)
}

type handler struct {
	services *services.Services
}

func New(
	services *services.Services,
) ImageHandler {
	return &handler{
		services: services,
	}
}
