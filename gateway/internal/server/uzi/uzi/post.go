package uzi

import (
	"context"

	"github.com/AlekSi/pointer"

	api "composition-api/internal/generated/http/api"
	uziSrv "composition-api/internal/services/uzi"
)

func (h *handler) UziPost(ctx context.Context, req *api.UziPostReq) (api.UziPostRes, error) {
	uziID, err := h.services.UziService.Create(ctx, uziSrv.CreateUziArg{
		File:       req.File,
		Projection: req.Projection,
		ExternalID: req.ExternalID,
		DeviceID:   req.DeviceID,
	})
	if err != nil {
		return nil, err
	}

	return pointer.To(api.SimpleUuid{ID: uziID}), nil
}
