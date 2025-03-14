package node_segment

import (
	"context"

	pb "uzi/internal/generated/grpc/service"
	"uzi/internal/services"
)

type NodeSegmentHandler interface {
	CreateNodeWithSegments(ctx context.Context, in *pb.CreateNodeWithSegmentsIn) (*pb.CreateNodeWithSegmentsOut, error)

	GetNodesWithSegmentsByImageId(ctx context.Context, in *pb.GetNodesWithSegmentsByImageIdIn) (*pb.GetNodesWithSegmentsByImageIdOut, error)
}

type handler struct {
	services *services.Services
}

func New(
	services *services.Services,
) NodeSegmentHandler {
	return &handler{
		services: services,
	}
}
