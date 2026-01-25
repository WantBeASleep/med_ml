package download

import (
	"context"

	api "composition-api/internal/generated/http/api"
)

func (h *handler) DownloadCytologyCytologyIDOriginalImageIDGet(ctx context.Context, params api.DownloadCytologyCytologyIDOriginalImageIDGetParams) (api.DownloadCytologyCytologyIDOriginalImageIDGetRes, error) {
	image, err := h.services.DownloadService.GetCytologyImage(ctx, params.CytologyID, params.OriginalImageID)
	if err != nil {
		return nil, err
	}

	return &api.DownloadCytologyCytologyIDOriginalImageIDGetOKApplicationOctetStream{
		Data: image,
	}, nil
}
