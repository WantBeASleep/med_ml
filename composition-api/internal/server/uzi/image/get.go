package image

import (
	"context"

	"github.com/AlekSi/pointer"

	api "composition-api/internal/generated/http/api"
	mappers "composition-api/internal/server/uzi/mappers"
)

func (h *handler) UziIDImagesGet(ctx context.Context, params api.UziIDImagesGetParams) (api.UziIDImagesGetRes, error) {
	images, err := h.services.ImageService.GetImagesByUziID(ctx, params.ID)
	if err != nil {
		return nil, err
	}

	return pointer.To(api.UziIDImagesGetOKApplicationJSON(mappers.SliceImage(images))), nil
}
