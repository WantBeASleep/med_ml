// время покажет, но пока выглядит будто бесполезный пакет, раз воткнул сюда дженерики, можно уж было
// вообще везде либу использовать просто пихая интерфейс

package producers

import (
	"context"

	"github.com/IBM/sarama"
	dbuslib "github.com/WantBeASleep/med_ml_lib/dbus"

	uziuploadpb "composition-api/internal/generated/dbus/produce/uziupload"
)

type Producer interface {
	SendUziUpload(ctx context.Context, msg *uziuploadpb.UziUpload) error
}

type producer struct {
	producerUziUpload dbuslib.Producer[*uziuploadpb.UziUpload]
}

func New(
	client sarama.SyncProducer,
) Producer {
	producerUziUpload := dbuslib.NewProducer[*uziuploadpb.UziUpload](
		client,
		"uziupload",
	)

	return &producer{
		producerUziUpload: producerUziUpload,
	}
}

func (a *producer) SendUziUpload(ctx context.Context, msg *uziuploadpb.UziUpload) error {
	return a.producerUziUpload.Send(ctx, msg)
}
