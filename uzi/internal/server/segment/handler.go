package segment

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"

	pb "uzi/internal/generated/grpc/service"
	"uzi/internal/services"
)

type SegmentHandler interface {
	CreateSegment(ctx context.Context, in *pb.CreateSegmentIn) (*pb.CreateSegmentOut, error)

	GetSegmentsByNodeId(ctx context.Context, in *pb.GetSegmentsByNodeIdIn) (*pb.GetSegmentsByNodeIdOut, error)

	UpdateSegment(ctx context.Context, in *pb.UpdateSegmentIn) (*pb.UpdateSegmentOut, error)

	DeleteSegment(ctx context.Context, in *pb.DeleteSegmentIn) (*empty.Empty, error)
}

type handler struct {
	services *services.Services
}

func New(
	services *services.Services,
) SegmentHandler {
	return &handler{
		services: services,
	}
}
