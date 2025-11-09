package uzi

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"composition-api/internal/domain"
	uzi_domain "composition-api/internal/domain/uzi"

	"github.com/AlekSi/pointer"

	api "composition-api/internal/generated/http/api"
	mappers "composition-api/internal/server/mappers"
	"composition-api/internal/server/security"
	uziSrv "composition-api/internal/services/uzi"
)

var uziProjectionMap = map[api.UziPostReqProjection]uzi_domain.UziProjection{
	api.UziPostReqProjectionCross: uzi_domain.UziProjectionCross,
	api.UziPostReqProjectionLong:  uzi_domain.UziProjectionLong,
}

func (h *handler) UziPost(ctx context.Context, req *api.UziPostReq) (api.UziPostRes, error) {
	token, err := security.ParseToken(ctx)
	if err != nil {
		return nil, err
	}

	contentType := req.File.Header.Get("Content-Type")
	if contentType != "image/tiff" {
		return &api.UziPostBadRequest{
			StatusCode: http.StatusBadRequest,
			Response: api.Error{
				Message: fmt.Sprintf("Неверный формат файла, ожидается: image/tiff, получено: %s", contentType),
			},
		}, nil
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
		switch {
		case errors.Is(err, domain.ErrBadRequest):
			return &api.UziPostBadRequest{
				StatusCode: http.StatusBadRequest,
				Response: api.Error{
					Message: "Неверный формат запроса или файла",
				},
			}, nil
		case errors.Is(err, domain.ErrUnprocessableEntity):
			return &api.UziPostUnprocessableEntity{
				StatusCode: http.StatusUnprocessableEntity,
				Response: api.Error{
					Message: "Ошибка валидации данных",
				},
			}, nil
		default:
			return nil, err
		}
	}

	return pointer.To(api.SimpleUuid{ID: uziID}), nil
}
