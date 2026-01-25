package cytology_image

import (
	"context"

	api "composition-api/internal/generated/http/api"
	services "composition-api/internal/services"
)

type CytologyImageHandler interface {
	CytologyCreateCreate(ctx context.Context, req *api.CytologyCreateCreateReq) (api.CytologyCreateCreateRes, error)
	CytologyRead(ctx context.Context, params api.CytologyReadParams) (api.CytologyReadRes, error)
	CytologySegmentsList(ctx context.Context, params api.CytologySegmentsListParams) (api.CytologySegmentsListRes, error)
	CytologyUpdateUpdate(ctx context.Context, req *api.CytologyUpdateUpdateReq, params api.CytologyUpdateUpdateParams) (api.CytologyUpdateUpdateRes, error)
	CytologyUpdatePartialUpdate(ctx context.Context, req *api.CytologyUpdatePartialUpdateReq, params api.CytologyUpdatePartialUpdateParams) (api.CytologyUpdatePartialUpdateRes, error)
	CytologyCopyCreate(ctx context.Context, req *api.CytologyCopyCreateReq) (api.CytologyCopyCreateRes, error)
	CytologyHistoryRead(ctx context.Context, params api.CytologyHistoryReadParams) (api.CytologyHistoryReadRes, error)
}

type handler struct {
	services *services.Services
}

func NewHandler(services *services.Services) CytologyImageHandler {
	return &handler{
		services: services,
	}
}
