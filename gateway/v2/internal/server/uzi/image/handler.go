package image

import (
	"context"

	api "gateway/internal/generated/http/api"
)

type Handler interface {
	UziIDImagesGet(ctx context.Context, params api.UziIDImagesGetParams) (api.UziIDImagesGetRes, error)
}

type handler struct {
}

func NewHandler() Handler {
	return &handler{}
}
