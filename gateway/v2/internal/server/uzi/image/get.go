package image

import (
	"context"

	"github.com/AlekSi/pointer"

	api "gateway/internal/generated/http/api"
	mappers "gateway/internal/server/uzi/mappers"
)

func (h *handler) UziIDImagesGet(ctx context.Context, params api.UziIDImagesGetParams) (api.UziIDImagesGetRes, error) {
	images, err := h.services.ImageService.GetImagesByUziID(ctx, params.ID)
	if err != nil {
		return nil, err
	}

	return pointer.To(api.UziIDImagesGetOKApplicationJSON(mappers.SliceImage(images))), nil
}
