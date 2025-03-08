// время покажет, но пока выглядит будто бесполезный пакет, раз воткнул сюда дженерики, можно уж было
// вообще везде либу использовать просто пихая интерфейс

package dbus

import (
	"context"

	uzicompletepb "uzi/internal/generated/dbus/produce/uzicomplete"
	uzisplittedpb "uzi/internal/generated/dbus/produce/uzisplitted"

	dbuslib "github.com/WantBeASleep/med_ml_lib/dbus"
)

type Producer interface {
	SendUziSplitted(ctx context.Context, msg *uzisplittedpb.UziSplitted) error
	SendUziComplete(ctx context.Context, msg *uzicompletepb.UziComplete) error
}

type producer struct {
	producerUziSplitted dbuslib.Producer[*uzisplittedpb.UziSplitted]
	producerUziComplete dbuslib.Producer[*uzicompletepb.UziComplete]
}

func New(
	producerUziSplitted dbuslib.Producer[*uzisplittedpb.UziSplitted],
	producerUziComplete dbuslib.Producer[*uzicompletepb.UziComplete],
) Producer {
	return &producer{
		producerUziSplitted: producerUziSplitted,
		producerUziComplete: producerUziComplete,
	}
}

func (a *producer) SendUziSplitted(ctx context.Context, msg *uzisplittedpb.UziSplitted) error {
	return a.producerUziSplitted.Send(ctx, msg)
}

func (a *producer) SendUziComplete(ctx context.Context, msg *uzicompletepb.UziComplete) error {
	return a.producerUziComplete.Send(ctx, msg)
}
