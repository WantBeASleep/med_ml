package segment

import (
	"context"

	api "gateway/internal/generated/http/api"
	apimappers "gateway/internal/server/mappers"
	mappers "gateway/internal/server/uzi/mappers"
	segmentSrv "gateway/internal/services/segment"
)

func (h *handler) UziSegmentIDPatch(ctx context.Context, req *api.UziSegmentIDPatchReq, params api.UziSegmentIDPatchParams) (api.UziSegmentIDPatchRes, error) {
	segment, err := h.services.SegmentService.Update(ctx, segmentSrv.UpdateSegmentArg{
		Id:        params.ID,
		Tirads_23: apimappers.FromOptFloat64(req.Tirads23),
		Tirads_4:  apimappers.FromOptFloat64(req.Tirads4),
		Tirads_5:  apimappers.FromOptFloat64(req.Tirads5),
	})
	if err != nil {
		return nil, err
	}

	resp, err := mappers.Segment(segment)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
