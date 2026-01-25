package cytology

import (
	"composition-api/internal/server/cytology/cytology_image"
	"composition-api/internal/server/cytology/original_image"
	"composition-api/internal/server/cytology/segmentation"
	services "composition-api/internal/services"
)

type CytologyRoute interface {
	cytology_image.CytologyImageHandler
	original_image.OriginalImageHandler
	segmentation.SegmentationHandler
}

type cytologyRoute struct {
	cytology_image.CytologyImageHandler
	original_image.OriginalImageHandler
	segmentation.SegmentationHandler
}

func NewCytologyRoute(services *services.Services) CytologyRoute {
	cytologyImageHandler := cytology_image.NewHandler(services)
	originalImageHandler := original_image.NewHandler(services)
	segmentationHandler := segmentation.NewHandler(services)

	return &cytologyRoute{
		CytologyImageHandler: cytologyImageHandler,
		OriginalImageHandler: originalImageHandler,
		SegmentationHandler:  segmentationHandler,
	}
}
