package segmentation

import (
	"context"

	pb "cytology/internal/generated/grpc/service"
	"cytology/internal/services"

	"google.golang.org/protobuf/types/known/emptypb"
)

type SegmentationHandler interface {
	CreateSegmentation(ctx context.Context, req *pb.CreateSegmentationIn) (*pb.CreateSegmentationOut, error)
	GetSegmentationById(ctx context.Context, req *pb.GetSegmentationByIdIn) (*pb.GetSegmentationByIdOut, error)
	GetSegmentsByGroupId(ctx context.Context, req *pb.GetSegmentsByGroupIdIn) (*pb.GetSegmentsByGroupIdOut, error)
	UpdateSegmentation(ctx context.Context, req *pb.UpdateSegmentationIn) (*pb.UpdateSegmentationOut, error)
	DeleteSegmentation(ctx context.Context, req *pb.DeleteSegmentationIn) (*emptypb.Empty, error)
}

type handler struct {
	services *services.Services
}

func New(services *services.Services) SegmentationHandler {
	return &handler{
		services: services,
	}
}
