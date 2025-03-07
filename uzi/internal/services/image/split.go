package image

import (
	"context"
	"errors"
	"fmt"
	"path/filepath"

	uzisplittedpb "uzi/internal/generated/dbus/produce/uzisplitted"

	"uzi/internal/domain"
	"uzi/internal/repository/image/entity"
	"uzi/internal/services/splitter"

	"github.com/google/uuid"
)

var ErrUziDoesntExist = errors.New("uzi doesnt exist")

func (s *service) SplitUzi(ctx context.Context, id uuid.UUID) error {
	fileRepo := s.dao.NewFileRepo()

	exists, err := s.dao.NewUziQuery(ctx).CheckExist(id)
	if err != nil {
		return fmt.Errorf("check exists uzi: %w", err)
	}
	if !exists {
		return ErrUziDoesntExist
	}

	file, closer, err := fileRepo.GetFileViaTemp(ctx, filepath.Join(id.String(), id.String()))
	if err != nil {
		return fmt.Errorf("get file via temp: %w", err)
	}
	defer closer()

	splitterSrv := splitter.New()
	splitted, err := splitterSrv.SplitToPng(file)
	if err != nil {
		return fmt.Errorf("split img to png: %w", err)
	}

	images := make([]domain.Image, len(splitted))
	for i := range images {
		images[i].Id = uuid.New()
		images[i].UziID = id
		images[i].Page = i + 1
	}

	for i, image := range images {
		if err := fileRepo.LoadFile(
			ctx,
			filepath.Join(id.String(), image.Id.String(), image.Id.String()),
			splitted[i],
		); err != nil {
			return fmt.Errorf("load file to S3: %w", err)
		}
	}

	if err := s.dao.NewImageQuery(ctx).InsertImages(entity.Image{}.SliceFromDomain(images)...); err != nil {
		return fmt.Errorf("insert images: %w", err)
	}

	imageIds := make([]uuid.UUID, 0, len(images))
	for _, image := range images {
		imageIds = append(imageIds, image.Id)
	}

	if err := s.adapter.SendUziSplitted(ctx, &uzisplittedpb.UziSplitted{
		UziId:   id.String(),
		PagesId: uuid.UUIDs(imageIds).Strings(),
	}); err != nil {
		return fmt.Errorf("send to uzisplitted topic: %w", err)
	}

	return nil
}
