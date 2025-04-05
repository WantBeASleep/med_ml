package uzi

import (
	"context"

	api "composition-api/internal/generated/http/api"
)

func (h *handler) UziIDDelete(ctx context.Context, params api.UziIDDeleteParams) (api.UziIDDeleteRes, error) {
	err := h.services.UziService.DeleteByID(ctx, params.ID)
	if err != nil {
		return nil, err
	}

	return &api.UziIDDeleteOK{}, nil
}
