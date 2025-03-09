package flow

import (
	"context"
	"fmt"

	"github.com/brianvoe/gofakeit/v7"

	"uzi/internal/domain"
	pb "uzi/internal/generated/grpc/service"
)

var DeviceInit flowfuncDepsInjector = func(deps *Deps) flowfunc {
	return func(ctx context.Context, data FlowData) (FlowData, error) {
		data = FlowData{}

		deviceName := gofakeit.MovieName()

		resp, err := deps.Adapter.CreateDevice(ctx, &pb.CreateDeviceIn{Name: deviceName})
		if err != nil {
			return FlowData{}, fmt.Errorf("create device: %w", err)
		}

		data.Device = domain.Device{Id: int(resp.Id), Name: deviceName}
		return data, nil
	}
}
