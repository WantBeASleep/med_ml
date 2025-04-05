package node

import (
	"context"

	"github.com/AlekSi/pointer"

	domain "composition-api/internal/domain/uzi"
	api "composition-api/internal/generated/http/api"
	apimappers "composition-api/internal/server/mappers"
	mappers "composition-api/internal/server/uzi/mappers"
	nodeSrv "composition-api/internal/services/node"
)

func (h *handler) UziNodesIDPatch(ctx context.Context, req *api.UziNodesIDPatchReq, params api.UziNodesIDPatchParams) (api.UziNodesIDPatchRes, error) {
	var validation *domain.NodeValidation
	switch {
	case req.Validation.Null:
		validation = pointer.To(domain.NodeValidationNull)
	case req.Validation.IsSet():
		validation = pointer.To(domain.NodeValidation(req.Validation.Value))
	}

	node, err := h.services.NodeService.UpdateNode(ctx, nodeSrv.UpdateNodeArg{
		Id:         params.ID,
		Validation: validation,
		Tirads_23:  apimappers.FromOptFloat64(req.Tirads23),
		Tirads_4:   apimappers.FromOptFloat64(req.Tirads4),
		Tirads_5:   apimappers.FromOptFloat64(req.Tirads5),
	})
	if err != nil {
		return nil, err
	}

	return pointer.To(mappers.Node{}.Domain(node)), nil
}
