package flow

import (
	"context"

	"med/internal/domain"
	pb "med/internal/generated/grpc/service"
)

type FlowData struct {
	Doctor  domain.Doctor
	Patient domain.Patient
	Card    domain.Card
}

type Deps struct {
	Adapter pb.MedSrvClient
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
