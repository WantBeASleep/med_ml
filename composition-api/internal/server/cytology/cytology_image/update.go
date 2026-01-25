package cytology_image

import (
	"context"
	"errors"
	"net/http"

	"composition-api/internal/domain"

	"github.com/AlekSi/pointer"
	"github.com/google/uuid"

	api "composition-api/internal/generated/http/api"
	mappers "composition-api/internal/server/cytology/mappers"
)

func (h *handler) CytologyUpdateUpdate(ctx context.Context, req *api.CytologyUpdateUpdateReq, params api.CytologyUpdateUpdateParams) (api.CytologyUpdateUpdateRes, error) {
	id, err := uuid.Parse(params.ID)
	if err != nil {
		return &api.CytologyUpdateUpdateBadRequest{
			StatusCode: http.StatusBadRequest,
			Response: api.Error{
				Message: "Неверный формат ID",
			},
		}, nil
	}

	arg := mappers.CytologyImage{}.UpdateArgFromCytologyUpdateUpdateReq(id, req)

	img, err := h.services.CytologyService.UpdateCytologyImage(ctx, arg)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrNotFound):
			return &api.CytologyUpdateUpdateNotFound{
				StatusCode: http.StatusNotFound,
				Response: api.Error{
					Message: "Цитологическое исследование не найдено",
				},
			}, nil
		case errors.Is(err, domain.ErrBadRequest):
			return &api.CytologyUpdateUpdateBadRequest{
				StatusCode: http.StatusBadRequest,
				Response: api.Error{
					Message: "Неверный формат запроса",
				},
			}, nil
		default:
			return nil, err
		}
	}

	result := mappers.CytologyImage{}.ToCytologyUpdateUpdateOK(img, req)
	return pointer.To(result), nil
}

func (h *handler) CytologyUpdatePartialUpdate(ctx context.Context, req *api.CytologyUpdatePartialUpdateReq, params api.CytologyUpdatePartialUpdateParams) (api.CytologyUpdatePartialUpdateRes, error) {
	id, err := uuid.Parse(params.ID)
	if err != nil {
		return &api.CytologyUpdatePartialUpdateBadRequest{
			StatusCode: http.StatusBadRequest,
			Response: api.Error{
				Message: "Неверный формат ID",
			},
		}, nil
	}

	arg := mappers.CytologyImage{}.UpdateArgFromCytologyUpdatePartialUpdateReq(id, req)

	img, err := h.services.CytologyService.UpdateCytologyImage(ctx, arg)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrNotFound):
			return &api.CytologyUpdatePartialUpdateNotFound{
				StatusCode: http.StatusNotFound,
				Response: api.Error{
					Message: "Цитологическое исследование не найдено",
				},
			}, nil
		case errors.Is(err, domain.ErrBadRequest):
			return &api.CytologyUpdatePartialUpdateBadRequest{
				StatusCode: http.StatusBadRequest,
				Response: api.Error{
					Message: "Неверный формат запроса",
				},
			}, nil
		default:
			return nil, err
		}
	}

	result := mappers.CytologyImage{}.ToCytologyUpdatePartialUpdateOK(img, req)
	return pointer.To(result), nil
}
