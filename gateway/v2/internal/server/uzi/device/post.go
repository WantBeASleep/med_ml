package device

import (
	"context"

	api "gateway/internal/generated/http/api"
)

func (h *handler) UziDevicePost(ctx context.Context, req *api.UziDevicePostReq) (api.UziDevicePostRes, error) {
	id, err := h.services.DeviceService.Create(ctx, req.Name)
	if err != nil {
		return nil, err
	}

	return &api.UziDevicePostOK{ID: id}, nil
}
