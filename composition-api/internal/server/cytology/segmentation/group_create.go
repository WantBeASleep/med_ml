package segmentation

import (
	"context"
	"errors"
	"net/http"

	"composition-api/internal/domain"

	"github.com/AlekSi/pointer"

	api "composition-api/internal/generated/http/api"
	mappers "composition-api/internal/server/cytology/mappers"
)

func (h *handler) CytologySegmentGroupCreateCreate(ctx context.Context, req *api.CytologySegmentGroupCreateCreateReq, params api.CytologySegmentGroupCreateCreateParams) (api.CytologySegmentGroupCreateCreateRes, error) {
	groupArg, segArg := mappers.SegmentationGroup{}.CreateArgFromCytologySegmentGroupCreateCreateReq(params.CytologyImgID, req)

	// Сначала создаем группу сегментации
	groupID, err := h.services.CytologyService.CreateSegmentationGroup(ctx, groupArg)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrBadRequest):
			return &api.CytologySegmentGroupCreateCreateBadRequest{
				StatusCode: http.StatusBadRequest,
				Response: api.Error{
					Message: "Неверный формат запроса",
				},
			}, nil
		case errors.Is(err, domain.ErrNotFound):
			return &api.CytologySegmentGroupCreateCreateNotFound{
				StatusCode: http.StatusNotFound,
				Response: api.Error{
					Message: "Цитологическое исследование не найдено",
				},
			}, nil
		case errors.Is(err, domain.ErrUnprocessableEntity):
			return &api.CytologySegmentGroupCreateCreateUnprocessableEntity{
				StatusCode: http.StatusUnprocessableEntity,
				Response: api.Error{
					Message: "Ошибка валидации данных",
				},
			}, nil
		default:
			return nil, err
		}
	}

	// Затем создаем сегмент с точками
	segArg.SegmentationGroupID = groupID
	_, err = h.services.CytologyService.CreateSegmentation(ctx, segArg)
	if err != nil {
		// Если не удалось создать сегмент, возвращаем ошибку
		// TODO: Возможно, стоит удалить созданную группу
		return nil, err
	}

	// Возвращаем согласно swagger.json
	result := api.CytologySegmentGroupCreateCreateCreated{
		ID: api.OptInt{
			Value: groupID,
			Set:   true,
		},
		Data: api.OptCytologySegmentGroupCreateCreateCreatedData{
			Value: api.CytologySegmentGroupCreateCreateCreatedData{
				Points: mappers.SegmentationGroup{}.ToCytologySegmentGroupCreateCreateCreatedDataPoints(req.Data.Points),
			},
			Set: true,
		},
		SegType: api.OptString{
			Value: string(req.SegType),
			Set:   true,
		},
	}

	return pointer.To(result), nil
}
