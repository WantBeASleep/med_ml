package cytology_image

import (
	"context"
	"errors"
	"fmt"
	"time"

	"cytology/internal/domain"
	cytologyImageEntity "cytology/internal/repository/cytology_image/entity"
	"cytology/internal/repository/entity"
	original_image "cytology/internal/services/original_image"

	"github.com/google/uuid"
)

func (s *service) CreateCytologyImage(ctx context.Context, arg CreateCytologyImageArg) (uuid.UUID, error) {
	ctx, err := s.dao.BeginTx(ctx)
	if err != nil {
		return uuid.Nil, fmt.Errorf("begin transaction: %w", err)
	}
	defer func() { _ = s.dao.RollbackTx(ctx) }()

	img := domain.CytologyImage{
		Id:                uuid.New(),
		ExternalID:        arg.ExternalID,
		DoctorID:          arg.DoctorID,
		PatientID:         arg.PatientID,
		DiagnosticNumber:  arg.DiagnosticNumber,
		DiagnosticMarking: arg.DiagnosticMarking,
		MaterialType:      arg.MaterialType,
		DiagnosDate:       time.Now(),
		IsLast:            true,
		Calcitonin:        arg.Calcitonin,
		CalcitoninInFlush: arg.CalcitoninInFlush,
		Thyroglobulin:     arg.Thyroglobulin,
		Details:           arg.Details,
		PrevID:            arg.PrevID,
		ParentPrevID:      arg.ParentPrevID,
		CreateAt:          time.Now(),
	}

	if err := s.dao.NewCytologyImageQuery(ctx).InsertCytologyImage(cytologyImageEntity.CytologyImage{}.FromDomain(img)); err != nil {
		var valErr *entity.DBValidationError
		if errors.As(err, &valErr) {
			return uuid.Nil, domain.ErrUnprocessableEntity
		}
		return uuid.Nil, fmt.Errorf("insert cytology image: %w", err)
	}

	// Если передан файл, создаем original_image
	if len(arg.File) > 0 && arg.ContentType != "" {
		originalImageService := original_image.New(s.dao)
		_, err = originalImageService.CreateOriginalImage(ctx, original_image.CreateOriginalImageArg{
			CytologyID:  img.Id,
			File:        arg.File,
			ContentType: arg.ContentType,
			DelayTime:   nil,
		})
		if err != nil {
			return uuid.Nil, fmt.Errorf("create original image: %w", err)
		}
	}

	if err := s.dao.CommitTx(ctx); err != nil {
		return uuid.Nil, fmt.Errorf("commit transaction: %w", err)
	}

	return img.Id, nil
}
