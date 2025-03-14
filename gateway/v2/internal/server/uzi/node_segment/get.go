package node_segment

import (
	"context"

	api "gateway/internal/generated/http/api"
	mappers "gateway/internal/server/uzi/mappers"
)

func (h *handler) UziImageIDNodesSegmentsGet(ctx context.Context, params api.UziImageIDNodesSegmentsGetParams) (api.UziImageIDNodesSegmentsGetRes, error) {
	nodes, segments, err := h.services.NodeSegmentService.GetNodeWithSegmentsByImageID(ctx, params.ID)
	if err != nil {
		return nil, err
	}

	segmentsResp, err := mappers.SliceSegment(segments)
	if err != nil {
		return nil, err
	}

	return &api.UziImageIDNodesSegmentsGetOK{
		Nodes:    mappers.SliceNode(nodes),
		Segments: segmentsResp,
	}, nil
}
