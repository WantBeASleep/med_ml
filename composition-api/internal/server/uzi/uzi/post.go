package uzi

import (
	"context"

	"github.com/AlekSi/pointer"

	domain "composition-api/internal/domain/uzi"
	api "composition-api/internal/generated/http/api"
	mappers "composition-api/internal/server/mappers"
	"composition-api/internal/server/security"
	uziSrv "composition-api/internal/services/uzi"
)

var uziProjectionMap = map[api.UziPostReqProjection]domain.UziProjection{
	api.UziPostReqProjectionCross: domain.UziProjectionCross,
	api.UziPostReqProjectionLong:  domain.UziProjectionLong,
}

func (h *handler) UziPost(ctx context.Context, req *api.UziPostReq) (api.UziPostRes, error) {
	token, err := security.ParseToken(ctx)
	if err != nil {
		return nil, err
	}

	uziID, err := h.services.UziService.Create(ctx, uziSrv.CreateUziArg{
		File:        req.File,
		Projection:  uziProjectionMap[req.Projection],
		ExternalID:  req.ExternalID,
		Author:      token.Id,
		DeviceID:    req.DeviceID,
		Description: mappers.FromOptString(req.Description),
	})
	if err != nil {
		return nil, err
	}

	return pointer.To(api.SimpleUuid{ID: uziID}), nil
}
