package segment

import (
	"context"
	"encoding/json"

	api "composition-api/internal/generated/http/api"
	segmentSrv "composition-api/internal/services/segment"
)

func (h *handler) UziSegmentPost(ctx context.Context, req *api.UziSegmentPostReq) (api.UziSegmentPostRes, error) {
	contor, err := json.Marshal(req.Contor)
	if err != nil {
		return nil, err
	}

	segmentID, err := h.services.SegmentService.Create(ctx, segmentSrv.CreateSegmentArg{
		ImageID:   req.ImageID,
		NodeID:    req.NodeID,
		Contor:    contor,
		Tirads_23: req.Tirads23,
		Tirads_4:  req.Tirads4,
		Tirads_5:  req.Tirads5,
	})
	if err != nil {
		return nil, err
	}
	return &api.SimpleUuid{ID: segmentID}, nil
}
