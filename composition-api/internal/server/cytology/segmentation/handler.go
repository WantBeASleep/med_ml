package segmentation

import (
	"context"

	api "composition-api/internal/generated/http/api"
	services "composition-api/internal/services"
)

type SegmentationHandler interface {
	CytologySegmentGroupCreateCreate(ctx context.Context, req *api.CytologySegmentGroupCreateCreateReq, params api.CytologySegmentGroupCreateCreateParams) (api.CytologySegmentGroupCreateCreateRes, error)
	CytologySegmentUpdateRead(ctx context.Context, params api.CytologySegmentUpdateReadParams) (api.CytologySegmentUpdateReadRes, error)
	CytologySegmentUpdateUpdate(ctx context.Context, req *api.CytologySegmentUpdateUpdateReq, params api.CytologySegmentUpdateUpdateParams) (api.CytologySegmentUpdateUpdateRes, error)
	CytologySegmentUpdatePartialUpdate(ctx context.Context, req *api.CytologySegmentUpdatePartialUpdateReq, params api.CytologySegmentUpdatePartialUpdateParams) (api.CytologySegmentUpdatePartialUpdateRes, error)
	CytologySegmentUpdateDelete(ctx context.Context, params api.CytologySegmentUpdateDeleteParams) (api.CytologySegmentUpdateDeleteRes, error)
}

type handler struct {
	services *services.Services
}

func NewHandler(services *services.Services) SegmentationHandler {
	return &handler{
		services: services,
	}
}
