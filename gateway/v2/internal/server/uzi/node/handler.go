package node

import (
	"context"

	api "gateway/internal/generated/http/api"
)

type Handler interface {
	GetNodesByUziID(ctx context.Context, params api.UziIDNodesGetParams) (api.UziIDNodesGetRes, error)
	UziNodesIDPatch(ctx context.Context, req *api.UziNodesIDPatchReq, params api.UziNodesIDPatchParams) (api.UziNodesIDPatchRes, error)
	UziNodesIDDelete(ctx context.Context, params api.UziNodesIDDeleteParams) (api.UziNodesIDDeleteRes, error)
}

type handler struct {
	
}

func NewHandler() Handler {
	return &handler{}
}
