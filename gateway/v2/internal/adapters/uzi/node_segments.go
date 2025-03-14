package uzi

import (
	"context"

	"github.com/google/uuid"

	"gateway/internal/adapters/uzi/mappers"
	domain "gateway/internal/domain/uzi"
	pb "gateway/internal/generated/grpc/clients/uzi"
)

func (a *adapter) CreateNodeWithSegments(ctx context.Context, in CreateNodeWithSegmentsIn) (uuid.UUID, []uuid.UUID, error) {
	req := &pb.CreateNodeWithSegmentsIn{
		Node: &pb.CreateNodeWithSegmentsIn_Node{
			UziId:     in.Node.UziID.String(),
			Ai:        in.Node.Ai,
			Tirads_23: in.Node.Tirads_23,
			Tirads_4:  in.Node.Tirads_4,
			Tirads_5:  in.Node.Tirads_5,
		},
	}

	for _, segment := range in.Segments {
		req.Segments = append(req.Segments, &pb.CreateNodeWithSegmentsIn_Segment{
			ImageId:   segment.ImageID.String(),
			Contor:    segment.Contor,
			Tirads_23: segment.Tirads_23,
			Tirads_4:  segment.Tirads_4,
			Tirads_5:  segment.Tirads_5,
		})
	}

	res, err := a.client.CreateNodeWithSegments(ctx, req)
	if err != nil {
		return uuid.Nil, nil, err
	}

	segmentIDs := make([]uuid.UUID, 0, len(res.SegmentIds))
	for _, id := range res.SegmentIds {
		segmentIDs = append(segmentIDs, uuid.MustParse(id))
	}

	return uuid.MustParse(res.NodeId), segmentIDs, nil
}

func (a *adapter) GetNodesWithSegmentsByImageId(ctx context.Context, id uuid.UUID) ([]domain.Node, []domain.Segment, error) {
	res, err := a.client.GetNodesWithSegmentsByImageId(ctx, &pb.GetNodesWithSegmentsByImageIdIn{Id: id.String()})
	if err != nil {
		return nil, nil, err
	}

	nodes := mappers.SliceNode(res.Nodes)
	segments := mappers.SliceSegment(res.Segments)

	return nodes, segments, nil
}
