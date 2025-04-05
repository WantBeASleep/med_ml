package node

import (
	"context"

	"github.com/AlekSi/pointer"

	api "composition-api/internal/generated/http/api"
	mappers "composition-api/internal/server/uzi/mappers"
)

func (h *handler) UziIDNodesGet(ctx context.Context, params api.UziIDNodesGetParams) (api.UziIDNodesGetRes, error) {
	nodes, err := h.services.NodeService.GetNodesByUziID(ctx, params.ID)
	if err != nil {
		return nil, err
	}

	return pointer.To(api.UziIDNodesGetOKApplicationJSON(mappers.Node{}.SliceDomain(nodes))), nil
}
