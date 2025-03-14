package node_segment

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "uzi/internal/generated/grpc/service"
	"uzi/internal/server/mappers"
)

func (h *handler) GetNodesWithSegmentsByImageId(ctx context.Context, in *pb.GetNodesWithSegmentsByImageIdIn) (*pb.GetNodesWithSegmentsByImageIdOut, error) {
	if _, err := uuid.Parse(in.Id); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "id is not a valid uuid: %s", err.Error())
	}

	nodes, segments, err := h.services.NodeSegment.GetNodesWithSegmentsByImageID(ctx, uuid.MustParse(in.Id))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}

	out := new(pb.GetNodesWithSegmentsByImageIdOut)
	out.Nodes = mappers.SliceNodeFromDomain(nodes)
	out.Segments = mappers.SliceSegmentFromDomain(segments)

	return out, nil
}
