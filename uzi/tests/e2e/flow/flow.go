// pattern: chain of responsibility
// https://refactoring.guru/ru/design-patterns/chain-of-responsibility

package flow

import (
	"context"

	"github.com/IBM/sarama"
	minio "github.com/minio/minio-go/v7"

	"uzi/internal/domain"
	pb "uzi/internal/generated/grpc/service"
)

type FlowData struct {
	Device   domain.Device
	Uzi      domain.Uzi
	Images   []domain.Image
	Nodes    []domain.Node
	Segments []domain.Segment
}

type Deps struct {
	Adapter pb.UziSrvClient
	Dbus    sarama.SyncProducer

	S3     *minio.Client
	Bucket string
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
