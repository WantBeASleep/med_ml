package device

import (
	"context"

	"github.com/AlekSi/pointer"

	api "composition-api/internal/generated/http/api"
	mappers "composition-api/internal/server/uzi/mappers"
)

func (h *handler) UziDevicesGet(ctx context.Context) (api.UziDevicesGetRes, error) {
	devices, err := h.services.DeviceService.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return pointer.To(api.UziDevicesGetOKApplicationJSON(mappers.SliceDevice(devices))), nil
}
