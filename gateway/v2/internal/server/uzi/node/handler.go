package node

import (
	"context"

	api "gateway/internal/generated/http/api"
	services "gateway/internal/services"
)

type NodeHandler interface {
	UziIDNodesGet(ctx context.Context, params api.UziIDNodesGetParams) (api.UziIDNodesGetRes, error)
	UziNodesIDPatch(ctx context.Context, req *api.UziNodesIDPatchReq, params api.UziNodesIDPatchParams) (api.UziNodesIDPatchRes, error)
	UziNodesIDDelete(ctx context.Context, params api.UziNodesIDDeleteParams) (api.UziNodesIDDeleteRes, error)
}

type handler struct {
	services *services.Services
}

func NewHandler(services *services.Services) NodeHandler {
	return &handler{
		services: services,
	}
}
