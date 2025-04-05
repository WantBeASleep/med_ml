package node

import (
	"context"

	api "composition-api/internal/generated/http/api"
)

func (h *handler) UziNodesIDDelete(ctx context.Context, params api.UziNodesIDDeleteParams) (api.UziNodesIDDeleteRes, error) {
	err := h.services.NodeService.DeleteNode(ctx, params.ID)
	if err != nil {
		return nil, err
	}

	return &api.UziNodesIDDeleteOK{}, nil
}
