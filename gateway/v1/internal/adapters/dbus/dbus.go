// время покажет, но пока выглядит будто бесполезный пакет, раз воткнул сюда дженерики, можно уж было
// вообще везде либу использовать просто пихая интерфейс

package dbus

import (
	"context"

	uziuploadpb "gateway/internal/generated/dbus/produce/uziupload"

	dbuslib "github.com/WantBeASleep/med_ml_lib/dbus"
)

type DbusAdapter interface {
	SendUziUpload(ctx context.Context, msg *uziuploadpb.UziUpload) error
}

type adapter struct {
	producerUziUpload dbuslib.Producer[*uziuploadpb.UziUpload]
}

func New(
	producerUziUpload dbuslib.Producer[*uziuploadpb.UziUpload],
) DbusAdapter {
	return &adapter{
		producerUziUpload: producerUziUpload,
	}
}

func (a *adapter) SendUziUpload(ctx context.Context, msg *uziuploadpb.UziUpload) error {
	return a.producerUziUpload.Send(ctx, msg)
}
