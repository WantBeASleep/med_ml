package segment

import (
	"context"

	api "gateway/internal/generated/http/api"
)

type Handler interface {
	UziSegmentPost(ctx context.Context, req *api.UziSegmentPostReq) (api.UziSegmentPostRes, error)
	UziSegmentIDPatch(ctx context.Context, req *api.UziSegmentIDPatchReq, params api.UziSegmentIDPatchParams) (api.UziSegmentIDPatchRes, error)
	UziSegmentIDDelete(ctx context.Context, params api.UziSegmentIDDeleteParams) (api.UziSegmentIDDeleteRes, error)
}

type handler struct {
}

func NewHandler() Handler {
	return &handler{}
}
