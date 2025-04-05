package segment

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "uzi/internal/generated/grpc/service"
	"uzi/internal/services/segment"
)

func (h *handler) CreateSegment(ctx context.Context, in *pb.CreateSegmentIn) (*pb.CreateSegmentOut, error) {
	if _, err := uuid.Parse(in.ImageId); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "image_id is not a valid uuid: %s", err.Error())
	}

	if _, err := uuid.Parse(in.NodeId); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "node_id is not a valid uuid: %s", err.Error())
	}

	if !json.Valid(in.Contor) {
		return nil, status.Errorf(codes.InvalidArgument, "contor is not a valid json")
	}

	id, err := h.services.Segment.CreateManualSegment(ctx, segment.CreateSegmentArg{
		ImageID:  uuid.MustParse(in.ImageId),
		NodeID:   uuid.MustParse(in.NodeId),
		Contor:   in.Contor,
		Tirads23: in.Tirads_23,
		Tirads4:  in.Tirads_4,
		Tirads5:  in.Tirads_5,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}

	return &pb.CreateSegmentOut{
		Id: id.String(),
	}, nil
}
