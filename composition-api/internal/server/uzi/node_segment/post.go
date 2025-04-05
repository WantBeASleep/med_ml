package node_segment

import (
	"context"
	"encoding/json"

	api "composition-api/internal/generated/http/api"
	"composition-api/internal/server/mappers"
	"composition-api/internal/services/node_segment"
)

func (h *handler) UziIDNodesSegmentsPost(ctx context.Context, req *api.UziIDNodesSegmentsPostReq, params api.UziIDNodesSegmentsPostParams) (api.UziIDNodesSegmentsPostRes, error) {
	arg := node_segment.CreateNodeWithSegmentArg{}
	arg.UziID = params.ID
	arg.Node = node_segment.CreateNodeWithSegmentArg_Node{
		Tirads_23:   req.Node.Tirads23,
		Tirads_4:    req.Node.Tirads4,
		Tirads_5:    req.Node.Tirads5,
		Description: mappers.FromOptString(req.Node.Description),
	}

	for _, segment := range req.Segments {
		contor, err := json.Marshal(segment.Contor)
		if err != nil {
			return nil, err
		}

		arg.Segments = append(arg.Segments, node_segment.CreateNodeWithSegmentArg_Segment{
			ImageID:   segment.ImageID,
			Contor:    contor,
			Tirads_23: segment.Tirads23,
			Tirads_4:  segment.Tirads4,
			Tirads_5:  segment.Tirads5,
		})
	}

	nodeID, segmentIDs, err := h.services.NodeSegmentService.CreateNodeWithSegment(ctx, arg)
	if err != nil {
		return nil, err
	}

	return &api.UziIDNodesSegmentsPostOK{
		NodeID:     nodeID,
		SegmentIds: segmentIDs,
	}, nil
}
