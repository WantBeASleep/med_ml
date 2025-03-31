// pattern: chain of responsibility
// https://refactoring.guru/ru/design-patterns/chain-of-responsibility

package flow

import (
	"context"

	"auth/internal/domain"
	pb "auth/internal/generated/grpc/service"

	"github.com/google/uuid"
)

type FlowUser struct {
	Id       uuid.UUID
	Email    string
	Password string
	Role     domain.Role
}

type FlowUserUnRegister struct {
	Id    uuid.UUID
	Email string
}

type RegisterUserTokens struct {
	Access  domain.Token
	Refresh domain.Token
}

type FlowData struct {
	RegisterUser   FlowUser
	UnRegisterUser FlowUserUnRegister
	Tokens         RegisterUserTokens
}

type Deps struct {
	Adapter pb.AuthSrvClient
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
