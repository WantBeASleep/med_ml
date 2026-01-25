package segmentation_group

import (
	"context"

	pb "cytology/internal/generated/grpc/service"
	"cytology/internal/services"

	"google.golang.org/protobuf/types/known/emptypb"
)

type SegmentationGroupHandler interface {
	CreateSegmentationGroup(ctx context.Context, req *pb.CreateSegmentationGroupIn) (*pb.CreateSegmentationGroupOut, error)
	GetSegmentationGroupsByCytologyId(ctx context.Context, req *pb.GetSegmentationGroupsByCytologyIdIn) (*pb.GetSegmentationGroupsByCytologyIdOut, error)
	UpdateSegmentationGroup(ctx context.Context, req *pb.UpdateSegmentationGroupIn) (*pb.UpdateSegmentationGroupOut, error)
	DeleteSegmentationGroup(ctx context.Context, req *pb.DeleteSegmentationGroupIn) (*emptypb.Empty, error)
}

type handler struct {
	services *services.Services
}

func New(services *services.Services) SegmentationGroupHandler {
	return &handler{
		services: services,
	}
}
