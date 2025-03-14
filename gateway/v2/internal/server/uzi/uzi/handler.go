package uzi

import (
	"context"

	api "gateway/internal/generated/http/api"
	services "gateway/internal/services"
)

type UziHandler interface {
	UziPost(ctx context.Context, req *api.UziPostReq) (api.UziPostRes, error)
	UziIDGet(ctx context.Context, params api.UziIDGetParams) (api.UziIDGetRes, error)
	UzisExternalIDGet(ctx context.Context, params api.UzisExternalIDGetParams) (api.UzisExternalIDGetRes, error)
	UziEchographicsUziIDGet(ctx context.Context, params api.UziEchographicsUziIDGetParams) (api.UziEchographicsUziIDGetRes, error)
	UziIDPatch(ctx context.Context, req *api.UziIDPatchReq, params api.UziIDPatchParams) (api.UziIDPatchRes, error)
	UziEchographicsUziIDPatch(ctx context.Context, req *api.Echographics, params api.UziEchographicsUziIDPatchParams) (api.UziEchographicsUziIDPatchRes, error)
}

type handler struct {
	services *services.Services
}

func NewHandler(services *services.Services) UziHandler {
	return &handler{
		services: services,
	}
}
