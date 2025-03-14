package segment

import (
	"context"

	"github.com/AlekSi/pointer"

	api "gateway/internal/generated/http/api"
	"gateway/internal/server/uzi/mappers"
)

func (h *handler) UziNodesIDSegmentsGet(ctx context.Context, params api.UziNodesIDSegmentsGetParams) (api.UziNodesIDSegmentsGetRes, error) {
	segments, err := h.services.SegmentService.GetByNodeID(ctx, params.ID)
	if err != nil {
		return nil, err
	}

	response, err := mappers.SliceSegment(segments)
	if err != nil {
		return nil, err
	}

	return pointer.To(api.UziNodesIDSegmentsGetOKApplicationJSON(response)), nil
}
