package segment

import (
	"context"

	api "composition-api/internal/generated/http/api"
	services "composition-api/internal/services"
)

type SegmentHandler interface {
	UziSegmentPost(ctx context.Context, req *api.UziSegmentPostReq) (api.UziSegmentPostRes, error)
	UziNodesIDSegmentsGet(ctx context.Context, params api.UziNodesIDSegmentsGetParams) (api.UziNodesIDSegmentsGetRes, error)
	UziSegmentIDPatch(ctx context.Context, req *api.UziSegmentIDPatchReq, params api.UziSegmentIDPatchParams) (api.UziSegmentIDPatchRes, error)
	UziSegmentIDDelete(ctx context.Context, params api.UziSegmentIDDeleteParams) (api.UziSegmentIDDeleteRes, error)
}

type handler struct {
	services *services.Services
}

func NewHandler(services *services.Services) SegmentHandler {
	return &handler{
		services: services,
	}
}
