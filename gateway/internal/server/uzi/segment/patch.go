package segment

import (
	"context"

	api "composition-api/internal/generated/http/api"
	apimappers "composition-api/internal/server/mappers"
	mappers "composition-api/internal/server/uzi/mappers"
	segmentSrv "composition-api/internal/services/segment"
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
