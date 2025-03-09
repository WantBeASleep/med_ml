package device

import (
	"context"

	api "gateway/internal/generated/http/api"
	services "gateway/internal/services"
)

type Handler interface {
	UziDevicePost(ctx context.Context, req *api.UziDevicePostReq) (api.UziDevicePostRes, error)
	UziDevicesGet(ctx context.Context) (api.UziDevicesGetRes, error)
}

type handler struct {
	services *services.Services
}

func NewHandler(services *services.Services) Handler {
	return &handler{
		services: services,
	}
}
