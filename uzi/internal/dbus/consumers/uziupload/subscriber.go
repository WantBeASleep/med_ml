package uziupload

import (
	"context"
	"fmt"

	"github.com/WantBeASleep/med_ml_lib/dbus"
	"github.com/google/uuid"

	pb "uzi/internal/generated/dbus/consume/uziupload"
	"uzi/internal/services"
)

type subscriber struct {
	services *services.Services
}

func New(
	services *services.Services,
) dbus.Consumer[*pb.UziUpload] {
	return &subscriber{
		services: services,
	}
}

func (h *subscriber) Consume(ctx context.Context, event *pb.UziUpload) error {
	if _, err := uuid.Parse(event.UziId); err != nil {
		return fmt.Errorf("uzi id is not uuid: %s", event.UziId)
	}

	if err := h.services.Image.SplitUzi(ctx, uuid.MustParse(event.UziId)); err != nil {
		return fmt.Errorf("process uziupload: %w", err)
	}
	return nil
}
