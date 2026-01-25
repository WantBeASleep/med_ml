package grpc

import (
	"cytology/internal/generated/grpc/service"
	"cytology/internal/server/cytology_image"
	"cytology/internal/server/original_image"
	"cytology/internal/server/segmentation"
	"cytology/internal/server/segmentation_group"
	"cytology/internal/services"
)

type Handler struct {
	cytology_image.CytologyImageHandler
	original_image.OriginalImageHandler
	segmentation_group.SegmentationGroupHandler
	segmentation.SegmentationHandler

	service.UnsafeCytologySrvServer
}

func New(
	services *services.Services,
) *Handler {
	cytologyImageHandler := cytology_image.New(services)
	originalImageHandler := original_image.New(services)
	segmentationGroupHandler := segmentation_group.New(services)
	segmentationHandler := segmentation.New(services)

	return &Handler{
		CytologyImageHandler:     cytologyImageHandler,
		OriginalImageHandler:      originalImageHandler,
		SegmentationGroupHandler: segmentationGroupHandler,
		SegmentationHandler:      segmentationHandler,
	}
}
