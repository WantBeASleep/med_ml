package device

import (
	"context"

	"github.com/AlekSi/pointer"

	api "gateway/internal/generated/http/api"
	mappers "gateway/internal/server/uzi/mappers"
)

func (h *handler) UziDevicesGet(ctx context.Context) (api.UziDevicesGetRes, error) {
	devices, err := h.services.DeviceService.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return pointer.To(api.UziDevicesGetOKApplicationJSON(mappers.SliceDevice(devices))), nil
}
