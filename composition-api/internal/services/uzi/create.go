package uzi

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/google/uuid"

	adapter "composition-api/internal/adapters/uzi"
	uziuploadpb "composition-api/internal/generated/dbus/produce/uziupload"
)

func (s *service) Create(ctx context.Context, in CreateUziArg) (uuid.UUID, error) {
	uziID, err := s.adapters.Uzi.CreateUzi(ctx, adapter.CreateUziIn{
		Projection: in.Projection,
		ExternalID: in.ExternalID.String(),
		DeviceID:   in.DeviceID,
	})
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("create uzi in microservice: %w", err)
	}

	err = s.dao.NewFileRepo().LoadFile(ctx, filepath.Join(uziID.String(), uziID.String()), in.File)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("load uzi file to s3: %w", err)
	}

	// TODO: сделать сагу
	err = s.dbus.SendUziUpload(ctx, &uziuploadpb.UziUpload{UziId: uziID.String()})
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("send uzi upload to dbus: %w", err)
	}

	return uziID, nil
}
