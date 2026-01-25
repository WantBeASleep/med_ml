package original_image

import (
	"context"

	pb "cytology/internal/generated/grpc/service"
	"cytology/internal/services"
)

type OriginalImageHandler interface {
	CreateOriginalImage(ctx context.Context, req *pb.CreateOriginalImageIn) (*pb.CreateOriginalImageOut, error)
	GetOriginalImageById(ctx context.Context, req *pb.GetOriginalImageByIdIn) (*pb.GetOriginalImageByIdOut, error)
	GetOriginalImagesByCytologyId(ctx context.Context, req *pb.GetOriginalImagesByCytologyIdIn) (*pb.GetOriginalImagesByCytologyIdOut, error)
	UpdateOriginalImage(ctx context.Context, req *pb.UpdateOriginalImageIn) (*pb.UpdateOriginalImageOut, error)
}

type handler struct {
	services *services.Services
}

func New(services *services.Services) OriginalImageHandler {
	return &handler{
		services: services,
	}
}
