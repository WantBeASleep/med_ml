package cytology_image

import (
	"context"
	"errors"
	"fmt"

	"cytology/internal/domain"
	cytologyImageEntity "cytology/internal/repository/cytology_image/entity"
	"cytology/internal/repository/entity"

	"github.com/google/uuid"
)

func (s *service) GetCytologyImageByID(ctx context.Context, id uuid.UUID) (domain.CytologyImage, error) {
	img, err := s.dao.NewCytologyImageQuery(ctx).GetCytologyImageByID(id)
	if err != nil {
		if errors.Is(err, entity.ErrNotFound) {
			return domain.CytologyImage{}, domain.ErrNotFound
		}
		return domain.CytologyImage{}, fmt.Errorf("get cytology image: %w", err)
	}

	return img.ToDomain(), nil
}

func (s *service) GetCytologyImagesByExternalID(ctx context.Context, externalID uuid.UUID) ([]domain.CytologyImage, error) {
	images, err := s.dao.NewCytologyImageQuery(ctx).GetCytologyImagesByExternalID(externalID)
	if err != nil {
		if errors.Is(err, entity.ErrNotFound) {
			return nil, domain.ErrNotFound
		}
		return nil, fmt.Errorf("get cytology images: %w", err)
	}

	return cytologyImageEntity.CytologyImage{}.SliceToDomain(images), nil
}

func (s *service) GetCytologyImagesByDoctorIdAndPatientId(ctx context.Context, doctorID, patientID uuid.UUID) ([]domain.CytologyImage, error) {
	images, err := s.dao.NewCytologyImageQuery(ctx).GetCytologyImagesByDoctorIdAndPatientId(doctorID, patientID)
	if err != nil {
		if errors.Is(err, entity.ErrNotFound) {
			return nil, domain.ErrNotFound
		}
		return nil, fmt.Errorf("get cytology images: %w", err)
	}

	return cytologyImageEntity.CytologyImage{}.SliceToDomain(images), nil
}
