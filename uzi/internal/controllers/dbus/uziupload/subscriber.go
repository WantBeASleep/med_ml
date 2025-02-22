package uziupload

import (
	"context"
	"errors"
	"fmt"

	"github.com/WantBeASleep/med_ml_lib/dbus"

	pb "uzi/internal/generated/dbus/consume/uziupload"
	"uzi/internal/services/image"

	"github.com/google/uuid"
)

type subscriber struct {
	imageSrv image.Service
}

func New(
	imageSrv image.Service,
) dbus.Consumer[*pb.UziUpload] {
	return &subscriber{
		imageSrv: imageSrv,
	}
}

func (h *subscriber) Consume(ctx context.Context, event *pb.UziUpload) error {
	if err := h.imageSrv.SplitUzi(ctx, uuid.MustParse(event.UziId)); err != nil {
		return errors.New("wrong msg type. uziupload required")
	}

	if err := h.imageSrv.SplitUzi(ctx, uuid.MustParse(event.UziId)); err != nil {
		return fmt.Errorf("process uziupload: %w", err)
	}
	return nil
}
