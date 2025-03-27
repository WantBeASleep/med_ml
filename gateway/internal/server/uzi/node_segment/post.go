package node_segment

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"

	api "composition-api/internal/generated/http/api"
	nodeSegmentSrv "composition-api/internal/services/node_segment"
)

func (h *handler) UziNodesSegmentsPost(ctx context.Context, req *api.UziNodesSegmentsPostReq) (api.UziNodesSegmentsPostRes, error) {
	arg := nodeSegmentSrv.CreateNodeWithSegmentArg{
		Node: struct {
			UziID     uuid.UUID
			Ai        bool
			Tirads_23 float64
			Tirads_4  float64
			Tirads_5  float64
		}{
			UziID:     req.Node.UziID,
			Ai:        req.Node.Ai,
			Tirads_23: req.Node.Tirads23,
			Tirads_4:  req.Node.Tirads4,
			Tirads_5:  req.Node.Tirads5,
		},
	}

	for _, segment := range req.Segments {
		contor, err := json.Marshal(segment.Contor)
		if err != nil {
			return nil, err
		}

		arg.Segments = append(arg.Segments, struct {
			ImageID   uuid.UUID
			Contor    []byte
			Tirads_23 float64
			Tirads_4  float64
			Tirads_5  float64
		}{
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

	return &api.UziNodesSegmentsPostOK{
		NodeID:     nodeID,
		SegmentIds: segmentIDs,
	}, nil
}
