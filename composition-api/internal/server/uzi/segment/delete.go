package segment

import (
	"context"

	api "composition-api/internal/generated/http/api"
)

func (h *handler) UziSegmentIDDelete(ctx context.Context, params api.UziSegmentIDDeleteParams) (api.UziSegmentIDDeleteRes, error) {
	err := h.services.SegmentService.Delete(ctx, params.ID)
	if err != nil {
		return nil, err
	}
	return &api.UziSegmentIDDeleteOK{}, nil
}
