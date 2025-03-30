package node_segment

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "uzi/internal/generated/grpc/service"
	"uzi/internal/services/node_segment"
)

func (h *handler) CreateNodeWithSegments(ctx context.Context, in *pb.CreateNodeWithSegmentsIn) (*pb.CreateNodeWithSegmentsOut, error) {
	if _, err := uuid.Parse(in.UziId); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "uzi_id is not a valid uuid: %s", err.Error())
	}

	for _, v := range in.Segments {
		if _, err := uuid.Parse(v.ImageId); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "image_id is not a valid uuid: %s", err.Error())
		}

		if !json.Valid(v.Contor) {
			return nil, status.Errorf(codes.InvalidArgument, "contor is not a valid json")
		}
	}

	segments := make([]node_segment.CreateNodesWithSegmentsArgSegment, 0, len(in.Segments))
	for _, v := range in.Segments {
		segments = append(segments, node_segment.CreateNodesWithSegmentsArgSegment{
			ImageID:  uuid.MustParse(v.ImageId),
			Contor:   v.Contor,
			Tirads23: v.Tirads_23,
			Tirads4:  v.Tirads_4,
			Tirads5:  v.Tirads_5,
		})
	}

	node := node_segment.CreateNodesWithSegmentsArgNode{
		Tirads23: in.Node.Tirads_23,
		Tirads4:  in.Node.Tirads_4,
		Tirads5:  in.Node.Tirads_5,
	}

	ids, err := h.services.NodeSegment.CreateManualNodesWithSegments(ctx,
		uuid.MustParse(in.UziId),
		[]node_segment.CreateNodesWithSegmentsArg{
			{
				Node:     node,
				Segments: segments,
			},
		},
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}

	out := new(pb.CreateNodeWithSegmentsOut)
	out.NodeId = ids[0].NodeID.String()
	for _, v := range ids[0].SegmentsID {
		out.SegmentIds = append(out.SegmentIds, v.String())
	}

	return out, nil
}
