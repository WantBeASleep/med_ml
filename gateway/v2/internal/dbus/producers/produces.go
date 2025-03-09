// время покажет, но пока выглядит будто бесполезный пакет, раз воткнул сюда дженерики, можно уж было
// вообще везде либу использовать просто пихая интерфейс

package producers

import (
	"context"

	uziuploadpb "gateway/internal/generated/dbus/produce/uziupload"

	dbuslib "github.com/WantBeASleep/med_ml_lib/dbus"
)

type Producer interface {
	SendUziUpload(ctx context.Context, msg *uziuploadpb.UziUpload) error
}

type producer struct {
	producerUziUpload dbuslib.Producer[*uziuploadpb.UziUpload]
}

func New(
	producerUziUpload dbuslib.Producer[*uziuploadpb.UziUpload],
) Producer {
	return &producer{
		producerUziUpload: producerUziUpload,
	}
}

func (a *producer) SendUziUpload(ctx context.Context, msg *uziuploadpb.UziUpload) error {
	return a.producerUziUpload.Send(ctx, msg)
}
