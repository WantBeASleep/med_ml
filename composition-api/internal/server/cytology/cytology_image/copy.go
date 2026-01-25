package cytology_image

import (
	"context"
	"errors"
	"net/http"

	"composition-api/internal/domain"

	"github.com/AlekSi/pointer"
	"github.com/google/uuid"

	api "composition-api/internal/generated/http/api"
)

func (h *handler) CytologyCopyCreate(ctx context.Context, req *api.CytologyCopyCreateReq) (api.CytologyCopyCreateRes, error) {
	// req.ID теперь UUID после обновления swagger
	var id uuid.UUID

	if req.ID != (uuid.UUID{}) {
		id = req.ID
	} else if req.Pk.Set {
		id = req.Pk.Value
	} else {
		return &api.CytologyCopyCreateBadRequest{
			StatusCode: http.StatusBadRequest,
			Response: api.Error{
				Message: "ID обязателен",
			},
		}, nil
	}

	// Копируем исследование
	newImg, err := h.services.CytologyService.CopyCytologyImage(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrNotFound):
			return &api.CytologyCopyCreateNotFound{
				StatusCode: http.StatusNotFound,
				Response: api.Error{
					Message: "Цитологическое исследование не найдено",
				},
			}, nil
		case errors.Is(err, domain.ErrBadRequest):
			return &api.CytologyCopyCreateBadRequest{
				StatusCode: http.StatusBadRequest,
				Response: api.Error{
					Message: "Неверный формат запроса",
				},
			}, nil
		default:
			return nil, err
		}
	}

	// Возвращаем UUID нового скопированного исследования
	result := api.CytologyCopyCreateCreated{
		Pk: api.OptUUID{
			Value: newImg.Id,
			Set:   true,
		},
		ID: api.OptUUID{
			Value: newImg.Id,
			Set:   true,
		},
	}

	return pointer.To(result), nil
}
