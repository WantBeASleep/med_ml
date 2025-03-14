// pattern: chain of responsibility
// https://refactoring.guru/ru/design-patterns/chain-of-responsibility

// Пакет FLOW - описывает последовательность основных действий для удобства тестирования

package uzi

import (
	"context"

	"github.com/google/uuid"

	domain "gateway/internal/domain/uzi"
	"gateway/internal/generated/http/api"
)

// FlowData - данные полученные при выполнении flow обработки узи
type FlowData struct {
	// Полученные данные от сервера
	//
	// например все ID генерируются на сервере и мы не можем их задавать в тесте
	Got struct {
		DeviceID int
		UziID    uuid.UUID
		Images   []domain.Image
		Nodes    []domain.Node
		Segments []domain.Segment
	}
	// Отправляемые данные на сервер
	Expected struct {
		DeviceName string

		UziProjection string
		UziExternalID uuid.UUID
	}
}

type Deps struct {
	Adapter api.Client
}

// точка запуска flow
type Flow interface {
	Do(ctx context.Context) (FlowData, error)
}

// часть общего flow
type flowfunc func(ctx context.Context, data FlowData) (FlowData, error)

type flowelem struct {
	flowfunc flowfunc
	next     *flowelem
}

func (f *flowelem) do(ctx context.Context, data FlowData) (FlowData, error) {
	flowRes, err := f.flowfunc(ctx, data)
	if err != nil {
		return FlowData{}, err
	}

	if f.next != nil {
		return f.next.do(ctx, flowRes)
	}
	return flowRes, nil
}

type flowfuncDepsInjector func(deps *Deps) flowfunc

type _flow struct {
	head *flowelem
}

func (f *_flow) Do(ctx context.Context) (FlowData, error) {
	return f.head.do(ctx, FlowData{})
}

func New(deps *Deps, flows ...flowfuncDepsInjector) Flow {
	if len(flows) == 0 {
		panic("flows is empty")
	}

	flowHead := &flowelem{}

	prevFlow := flowHead
	for _, flow := range flows {
		flowelem := &flowelem{flowfunc: flow(deps)}
		prevFlow.next = flowelem
		prevFlow = flowelem
	}

	return &_flow{head: flowHead.next}
}
