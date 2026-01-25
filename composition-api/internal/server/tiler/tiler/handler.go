package tiler

import (
	"bytes"
	"context"
	"io"
	"net/http"

	api "composition-api/internal/generated/http/api"
	services "composition-api/internal/services"
)

// TilerHandler интерфейс для обработчиков tiler
type TilerHandler interface {
	TilerDziFilePathGet(ctx context.Context, params api.TilerDziFilePathGetParams) (api.TilerDziFilePathGetRes, error)
	TilerDziFilePathFilesLevelColRowFormatGet(ctx context.Context, params api.TilerDziFilePathFilesLevelColRowFormatGetParams) (api.TilerDziFilePathFilesLevelColRowFormatGetRes, error)
}

type handler struct {
	services *services.Services
}

func NewHandler(services *services.Services) TilerHandler {
	return &handler{
		services: services,
	}
}

func (h *handler) TilerDziFilePathGet(ctx context.Context, params api.TilerDziFilePathGetParams) (api.TilerDziFilePathGetRes, error) {
	dzi, err := h.services.TilerService.GetDZI(ctx, params.FilePath)
	if err != nil {
		return &api.TilerDziFilePathGetBadRequest{
			StatusCode: http.StatusBadRequest,
			Response: api.Error{
				Message: "Failed to get DZI: " + err.Error(),
			},
		}, nil
	}

	return &api.TilerDziFilePathGetOK{
		Data: bytes.NewReader([]byte(dzi)),
	}, nil
}

func (h *handler) TilerDziFilePathFilesLevelColRowFormatGet(ctx context.Context, params api.TilerDziFilePathFilesLevelColRowFormatGetParams) (api.TilerDziFilePathFilesLevelColRowFormatGetRes, error) {
	// Конвертируем Format enum в string
	formatStr := string(params.Format)

	tile, err := h.services.TilerService.GetTile(ctx, params.FilePath, params.Level, params.Col, params.Row, formatStr)
	if err != nil {
		return &api.TilerDziFilePathFilesLevelColRowFormatGetBadRequest{
			StatusCode: http.StatusBadRequest,
			Response: api.Error{
				Message: "Failed to get tile: " + err.Error(),
			},
		}, nil
	}
	defer tile.Close()

	data, err := io.ReadAll(tile)
	if err != nil {
		return &api.TilerDziFilePathFilesLevelColRowFormatGetBadRequest{
			StatusCode: http.StatusBadRequest,
			Response: api.Error{
				Message: "Failed to read tile: " + err.Error(),
			},
		}, nil
	}

	// Определяем Content-Type в зависимости от формата
	switch params.Format {
	case api.TilerDziFilePathFilesLevelColRowFormatGetFormatJpeg, api.TilerDziFilePathFilesLevelColRowFormatGetFormatJPG:
		return &api.TilerDziFilePathFilesLevelColRowFormatGetOKImageJpeg{
			Data: bytes.NewReader(data),
		}, nil
	case api.TilerDziFilePathFilesLevelColRowFormatGetFormatPNG:
		return &api.TilerDziFilePathFilesLevelColRowFormatGetOKImagePNG{
			Data: bytes.NewReader(data),
		}, nil
	default:
		return &api.TilerDziFilePathFilesLevelColRowFormatGetOKImageJpeg{
			Data: bytes.NewReader(data),
		}, nil
	}
}
