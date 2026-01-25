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

func (h *handler) CytologySegmentUpdateRead(ctx context.Context, params api.CytologySegmentUpdateReadParams) (api.CytologySegmentUpdateReadRes, error) {
	seg, err := h.services.CytologyService.GetSegmentationById(ctx, params.ID)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return &api.CytologySegmentUpdateReadNotFound{
				StatusCode: http.StatusNotFound,
				Response: api.Error{
					Message: "Сегментация не найдена",
				},
			}, nil
		}
		return nil, err
	}

	result := mappers.Segmentation{}.ToCytologySegmentUpdateReadOK(seg)
	return pointer.To(result), nil
}

func (h *handler) CytologySegmentUpdateUpdate(ctx context.Context, req *api.CytologySegmentUpdateUpdateReq, params api.CytologySegmentUpdateUpdateParams) (api.CytologySegmentUpdateUpdateRes, error) {
	arg := mappers.Segmentation{}.UpdateArgFromCytologySegmentUpdateUpdateReq(params.ID, req)

	seg, err := h.services.CytologyService.UpdateSegmentation(ctx, arg)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrNotFound):
			return &api.CytologySegmentUpdateUpdateNotFound{
				StatusCode: http.StatusNotFound,
				Response: api.Error{
					Message: "Сегментация не найдена",
				},
			}, nil
		case errors.Is(err, domain.ErrBadRequest):
			return &api.CytologySegmentUpdateUpdateBadRequest{
				StatusCode: http.StatusBadRequest,
				Response: api.Error{
					Message: "Неверный формат запроса",
				},
			}, nil
		default:
			return nil, err
		}
	}

	result := mappers.Segmentation{}.ToCytologySegmentUpdateUpdateOK(seg, req)
	return pointer.To(result), nil
}

func (h *handler) CytologySegmentUpdatePartialUpdate(ctx context.Context, req *api.CytologySegmentUpdatePartialUpdateReq, params api.CytologySegmentUpdatePartialUpdateParams) (api.CytologySegmentUpdatePartialUpdateRes, error) {
	arg := mappers.Segmentation{}.UpdateArgFromCytologySegmentUpdatePartialUpdateReq(params.ID, req)

	seg, err := h.services.CytologyService.UpdateSegmentation(ctx, arg)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrNotFound):
			return &api.CytologySegmentUpdatePartialUpdateNotFound{
				StatusCode: http.StatusNotFound,
				Response: api.Error{
					Message: "Сегментация не найдена",
				},
			}, nil
		case errors.Is(err, domain.ErrBadRequest):
			return &api.CytologySegmentUpdatePartialUpdateBadRequest{
				StatusCode: http.StatusBadRequest,
				Response: api.Error{
					Message: "Неверный формат запроса",
				},
			}, nil
		default:
			return nil, err
		}
	}

	result := mappers.Segmentation{}.ToCytologySegmentUpdatePartialUpdateOK(seg, req)
	return pointer.To(result), nil
}

func (h *handler) CytologySegmentUpdateDelete(ctx context.Context, params api.CytologySegmentUpdateDeleteParams) (api.CytologySegmentUpdateDeleteRes, error) {
	err := h.services.CytologyService.DeleteSegmentation(ctx, params.ID)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return &api.CytologySegmentUpdateDeleteNotFound{
				StatusCode: http.StatusNotFound,
				Response: api.Error{
					Message: "Сегментация не найдена",
				},
			}, nil
		}
		return nil, err
	}

	return &api.CytologySegmentUpdateDeleteNoContent{}, nil
}
