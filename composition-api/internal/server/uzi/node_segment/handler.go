package node_segment

import (
	"context"

	api "composition-api/internal/generated/http/api"
	services "composition-api/internal/services"
)

type NodeSegmentHandler interface {
	UziIDNodesSegmentsPost(ctx context.Context, req *api.UziIDNodesSegmentsPostReq, params api.UziIDNodesSegmentsPostParams) (api.UziIDNodesSegmentsPostRes, error)
	UziImageIDNodesSegmentsGet(ctx context.Context, params api.UziImageIDNodesSegmentsGetParams) (api.UziImageIDNodesSegmentsGetRes, error)
}

type handler struct {
	services *services.Services
}

func NewHandler(services *services.Services) NodeSegmentHandler {
	return &handler{
		services: services,
	}
}
