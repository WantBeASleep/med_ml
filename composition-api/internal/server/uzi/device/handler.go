package device

import (
	"context"

	api "composition-api/internal/generated/http/api"
	services "composition-api/internal/services"
)

type DeviceHandler interface {
	UziDevicePost(ctx context.Context, req *api.UziDevicePostReq) (api.UziDevicePostRes, error)
	UziDevicesGet(ctx context.Context) (api.UziDevicesGetRes, error)
}

type handler struct {
	services *services.Services
}

func NewHandler(services *services.Services) DeviceHandler {
	return &handler{
		services: services,
	}
}
