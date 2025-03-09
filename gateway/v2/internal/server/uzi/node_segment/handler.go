package node_segment

import (
	"context"

	api "gateway/internal/generated/http/api"
	services "gateway/internal/services"
)

type Handler interface {
	UziNodesSegmentsPost(ctx context.Context, req *api.UziNodesSegmentsPostReq) (api.UziNodesSegmentsPostRes, error)
	UziImageIDNodesSegmentsGet(ctx context.Context, params api.UziImageIDNodesSegmentsGetParams) (api.UziImageIDNodesSegmentsGetRes, error)
}

type handler struct {
	
}

func NewHandler() Handler {
	
}
