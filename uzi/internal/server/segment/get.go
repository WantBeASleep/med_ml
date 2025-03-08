package segment

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "uzi/internal/generated/grpc/service"
	"uzi/internal/server/mappers"
)

func (h *handler) GetSegmentsByNodeId(ctx context.Context, in *pb.GetSegmentsByNodeIdIn) (*pb.GetSegmentsByNodeIdOut, error) {
	if _, err := uuid.Parse(in.NodeId); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "node_id is not a valid uuid: %s", err.Error())
	}

	segments, err := h.services.Segment.GetSegmentsByNodeID(ctx, uuid.MustParse(in.NodeId))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}

	out := new(pb.GetSegmentsByNodeIdOut)
	out.Segments = mappers.SliceSegmentFromDomain(segments)

	return out, nil
}
