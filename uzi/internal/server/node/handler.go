package node

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"

	pb "uzi/internal/generated/grpc/service"
	"uzi/internal/services"
)

type NodeHandler interface {
	GetNodesByUziId(ctx context.Context, in *pb.GetNodesByUziIdIn) (*pb.GetNodesByUziIdOut, error)
	UpdateNode(ctx context.Context, in *pb.UpdateNodeIn) (*pb.UpdateNodeOut, error)
	DeleteNode(ctx context.Context, in *pb.DeleteNodeIn) (*empty.Empty, error)
}

type handler struct {
	services *services.Services
}

func New(
	services *services.Services,
) NodeHandler {
	return &handler{
		services: services,
	}
}
