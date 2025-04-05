package download

import (
	"context"

	api "composition-api/internal/generated/http/api"
)

func (h *handler) DownloadUziIDImageIDGet(ctx context.Context, params api.DownloadUziIDImageIDGetParams) (api.DownloadUziIDImageIDGetRes, error) {
	image, err := h.services.DownloadService.GetImage(ctx, params.UziID, params.ImageID)
	if err != nil {
		return nil, err
	}

	return &api.DownloadUziIDImageIDGetOK{
		Data: image,
	}, nil
}
