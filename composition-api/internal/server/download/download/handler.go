package download

import (
	"context"

	api "composition-api/internal/generated/http/api"
	services "composition-api/internal/services"
)

type DownloadHandler interface {
	DownloadUziIDImageIDGet(ctx context.Context, params api.DownloadUziIDImageIDGetParams) (api.DownloadUziIDImageIDGetRes, error)
	DownloadCytologyCytologyIDOriginalImageIDGet(ctx context.Context, params api.DownloadCytologyCytologyIDOriginalImageIDGetParams) (api.DownloadCytologyCytologyIDOriginalImageIDGetRes, error)
}

type handler struct {
	services *services.Services
}

func NewHandler(services *services.Services) DownloadHandler {
	return &handler{
		services: services,
	}
}
