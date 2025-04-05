package uzi

import (
	"context"

	api "composition-api/internal/generated/http/api"
	services "composition-api/internal/services"
)

type UziHandler interface {
	UziPost(ctx context.Context, req *api.UziPostReq) (api.UziPostRes, error)
	UziIDGet(ctx context.Context, params api.UziIDGetParams) (api.UziIDGetRes, error)
	UzisExternalIDGet(ctx context.Context, params api.UzisExternalIDGetParams) (api.UzisExternalIDGetRes, error)
	UzisAuthorIDGet(ctx context.Context, params api.UzisAuthorIDGetParams) (api.UzisAuthorIDGetRes, error)
	UziIDEchographicsGet(ctx context.Context, params api.UziIDEchographicsGetParams) (api.UziIDEchographicsGetRes, error)
	UziIDPatch(ctx context.Context, req *api.UziIDPatchReq, params api.UziIDPatchParams) (api.UziIDPatchRes, error)
	UziIDEchographicsPatch(ctx context.Context, req *api.Echographics, params api.UziIDEchographicsPatchParams) (api.UziIDEchographicsPatchRes, error)
	UziIDDelete(ctx context.Context, params api.UziIDDeleteParams) (api.UziIDDeleteRes, error)
}

type handler struct {
	services *services.Services
}

func NewHandler(services *services.Services) UziHandler {
	return &handler{
		services: services,
	}
}
