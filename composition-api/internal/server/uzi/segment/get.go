package segment

import (
	"context"

	"github.com/AlekSi/pointer"

	api "composition-api/internal/generated/http/api"
	"composition-api/internal/server/uzi/mappers"
)

func (h *handler) UziNodesIDSegmentsGet(ctx context.Context, params api.UziNodesIDSegmentsGetParams) (api.UziNodesIDSegmentsGetRes, error) {
	segments, err := h.services.SegmentService.GetByNodeID(ctx, params.ID)
	if err != nil {
		return nil, err
	}

	response, err := mappers.Segment{}.SliceDomain(segments)
	if err != nil {
		return nil, err
	}

	return pointer.To(api.UziNodesIDSegmentsGetOKApplicationJSON(response)), nil
}
