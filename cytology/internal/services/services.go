package services

import (
	"cytology/internal/repository"
	"cytology/internal/services/cytology_image"
	"cytology/internal/services/original_image"
	"cytology/internal/services/segmentation"
	"cytology/internal/services/segmentation_group"
)

type Services struct {
	CytologyImage      cytology_image.Service
	OriginalImage      original_image.Service
	SegmentationGroup  segmentation_group.Service
	Segmentation       segmentation.Service
}

func New(
	dao repository.DAO,
) *Services {
	cytologyImage := cytology_image.New(dao)
	originalImage := original_image.New(dao)
	segmentationGroup := segmentation_group.New(dao)
	segmentation := segmentation.New(dao)

	return &Services{
		CytologyImage:     cytologyImage,
		OriginalImage:     originalImage,
		SegmentationGroup: segmentationGroup,
		Segmentation:      segmentation,
	}
}
