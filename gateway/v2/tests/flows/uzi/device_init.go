package uzi

import (
	"context"
	"fmt"

	"github.com/brianvoe/gofakeit/v7"

	"gateway/internal/generated/http/api"
)

var DeviceInit flowfuncDepsInjector = func(deps *Deps) flowfunc {
	return func(ctx context.Context, data FlowData) (FlowData, error) {
		data = FlowData{}

		deviceName := gofakeit.MovieName()

		resp, err := deps.Adapter.UziDevicePost(ctx, &api.UziDevicePostReq{Name: deviceName})
		if err != nil {
			return FlowData{}, fmt.Errorf("create device: %w", err)
		}

		switch v := resp.(type) {
		case *api.UziDevicePostOK:
			data.Got.DeviceID = v.ID
			data.Expected.DeviceName = deviceName

		case *api.ErrorStatusCode:
			return FlowData{}, fmt.Errorf("device post error: %w", v)

		default:
			return FlowData{}, fmt.Errorf("unexpected device post response: %T", v)
		}

		return data, nil
	}
}
