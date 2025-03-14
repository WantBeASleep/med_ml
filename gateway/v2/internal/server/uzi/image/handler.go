package image

import (
	"context"

	api "gateway/internal/generated/http/api"
	services "gateway/internal/services"
)

type ImageHandler interface {
	UziIDImagesGet(ctx context.Context, params api.UziIDImagesGetParams) (api.UziIDImagesGetRes, error)
}

type handler struct {
	services *services.Services
}

func NewHandler(services *services.Services) ImageHandler {
	return &handler{
		services: services,
	}
}
