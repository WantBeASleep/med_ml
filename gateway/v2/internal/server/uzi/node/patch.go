package node

import (
	"context"

	"github.com/AlekSi/pointer"

	api "gateway/internal/generated/http/api"
	apimappers "gateway/internal/server/mappers"
	mappers "gateway/internal/server/uzi/mappers"
	nodeSrv "gateway/internal/services/node"
)

func (h *handler) UziNodesIDPatch(ctx context.Context, req *api.UziNodesIDPatchReq, params api.UziNodesIDPatchParams) (api.UziNodesIDPatchRes, error) {
	node, err := h.services.NodeService.UpdateNode(ctx, nodeSrv.UpdateNodeArg{
		Id:        params.ID,
		Tirads_23: apimappers.FromOptFloat64(req.Tirads23),
		Tirads_4:  apimappers.FromOptFloat64(req.Tirads4),
		Tirads_5:  apimappers.FromOptFloat64(req.Tirads5),
	})
	if err != nil {
		return nil, err
	}

	return pointer.To(mappers.Node(node)), nil
}
