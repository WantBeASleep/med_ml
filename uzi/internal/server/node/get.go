package node

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "uzi/internal/generated/grpc/service"
	"uzi/internal/server/mappers"
)

func (h *handler) GetNodesByUziId(ctx context.Context, in *pb.GetNodesByUziIdIn) (*pb.GetNodesByUziIdOut, error) {
	if _, err := uuid.Parse(in.UziId); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "uzi_id is not a valid uuid: %s", err.Error())
	}

	nodes, err := h.services.Node.GetNodesByUziID(ctx, uuid.MustParse(in.UziId))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}

	out := new(pb.GetNodesByUziIdOut)
	out.Nodes = mappers.SliceNodeFromDomain(nodes)

	return out, nil
}
