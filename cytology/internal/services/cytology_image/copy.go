package cytology_image

import (
	"context"
	"errors"
	"fmt"
	"time"

	"cytology/internal/domain"
	cytologyImageEntity "cytology/internal/repository/cytology_image/entity"
	"cytology/internal/repository/entity"
	segmentationEntity "cytology/internal/repository/segmentation/entity"
	segmentationGroupEntity "cytology/internal/repository/segmentation_group/entity"

	"github.com/google/uuid"
)

func (s *service) CopyCytologyImage(ctx context.Context, id uuid.UUID) (domain.CytologyImage, error) {
	ctx, err := s.dao.BeginTx(ctx)
	if err != nil {
		return domain.CytologyImage{}, fmt.Errorf("begin transaction: %w", err)
	}
	defer func() { _ = s.dao.RollbackTx(ctx) }()

	// Получаем текущее исследование
	oldImg, err := s.GetCytologyImageByID(ctx, id)
	if err != nil {
		return domain.CytologyImage{}, fmt.Errorf("get cytology image: %w", err)
	}

	// Проверяем, что это последняя версия
	if !oldImg.IsLast {
		return domain.CytologyImage{}, domain.ErrBadRequest
	}

	// Устанавливаем is_last=false для старого
	oldImg.IsLast = false
	if err := s.dao.NewCytologyImageQuery(ctx).UpdateCytologyImage(cytologyImageEntity.CytologyImage{}.FromDomain(oldImg)); err != nil {
		return domain.CytologyImage{}, fmt.Errorf("update old cytology image: %w", err)
	}

	// Определяем parent_prev_id
	parentPrevID := oldImg.ParentPrevID
	if parentPrevID == nil {
		parentPrevID = &id
	}

	// Создаем новую копию
	newImg := domain.CytologyImage{
		Id:                uuid.New(),
		ExternalID:        oldImg.ExternalID,
		DoctorID:          oldImg.DoctorID,
		PatientID:         oldImg.PatientID,
		DiagnosticNumber:  oldImg.DiagnosticNumber,
		DiagnosticMarking: oldImg.DiagnosticMarking,
		MaterialType:      oldImg.MaterialType,
		DiagnosDate:       time.Now(),
		IsLast:            true,
		Calcitonin:        oldImg.Calcitonin,
		CalcitoninInFlush: oldImg.CalcitoninInFlush,
		Thyroglobulin:     oldImg.Thyroglobulin,
		Details:           oldImg.Details,
		PrevID:            &id,
		ParentPrevID:      parentPrevID,
		CreateAt:          time.Now(),
	}

	if err := s.dao.NewCytologyImageQuery(ctx).InsertCytologyImage(cytologyImageEntity.CytologyImage{}.FromDomain(newImg)); err != nil {
		var valErr *entity.DBValidationError
		if errors.As(err, &valErr) {
			return domain.CytologyImage{}, domain.ErrUnprocessableEntity
		}
		return domain.CytologyImage{}, fmt.Errorf("insert cytology image: %w", err)
	}

	// Копируем сегменты
	if err := s.copySegments(ctx, id, newImg.Id); err != nil {
		return domain.CytologyImage{}, fmt.Errorf("copy segments: %w", err)
	}

	if err := s.dao.CommitTx(ctx); err != nil {
		return domain.CytologyImage{}, fmt.Errorf("commit transaction: %w", err)
	}

	return newImg, nil
}

func (s *service) copySegments(ctx context.Context, oldCytologyID, newCytologyID uuid.UUID) error {
	// Получаем все группы сегментов для старого исследования
	oldGroups, err := s.dao.NewSegmentationGroupQuery(ctx).GetSegmentationGroupsByCytologyID(oldCytologyID)
	if err != nil && !errors.Is(err, domain.ErrNotFound) {
		return fmt.Errorf("get segmentation groups: %w", err)
	}

	// Создаем новые группы и копируем сегменты
	for _, oldGroup := range oldGroups {
		oldGroupDomain := oldGroup.ToDomain()
		newGroup := domain.SegmentationGroup{
			Id:         0, // ID будет сгенерирован БД
			CytologyID: newCytologyID,
			SegType:    oldGroupDomain.SegType,
			GroupType:  oldGroupDomain.GroupType,
			IsAI:       oldGroupDomain.IsAI,
			Details:    oldGroupDomain.Details,
			CreateAt:   time.Now(),
		}

		newGroupID, err := s.dao.NewSegmentationGroupQuery(ctx).InsertSegmentationGroup(segmentationGroupEntity.SegmentationGroup{}.FromDomain(newGroup))
		if err != nil {
			return fmt.Errorf("insert segmentation group: %w", err)
		}

		// Получаем сегменты для старой группы
		oldSegments, err := s.dao.NewSegmentationQuery(ctx).GetSegmentsByGroupID(oldGroupDomain.Id)
		if err != nil && !errors.Is(err, domain.ErrNotFound) {
			return fmt.Errorf("get segments: %w", err)
		}

		// Копируем сегменты
		for _, oldSegment := range oldSegments {
			oldSegmentDomain := oldSegment.ToDomain()
			newSegment := domain.Segmentation{
				Id:                  0, // ID будет сгенерирован БД
				SegmentationGroupID: newGroupID,
				Points:              oldSegmentDomain.Points,
				CreateAt:            time.Now(),
			}

			_, err = s.dao.NewSegmentationQuery(ctx).InsertSegmentation(segmentationEntity.Segmentation{}.FromDomain(newSegment))
			if err != nil {
				return fmt.Errorf("insert segmentation: %w", err)
			}
		}
	}

	return nil
}
