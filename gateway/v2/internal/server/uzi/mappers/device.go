package mappers

import (
	domain "gateway/internal/domain/uzi"
	api "gateway/internal/generated/http/api"
)

func Device(device domain.Device) api.Device {
	return api.Device{
		ID:   device.Id,
		Name: device.Name,
	}
}

func SliceDevice(devices []domain.Device) []api.Device {
	result := make([]api.Device, 0, len(devices))
	for _, device := range devices {
		result = append(result, Device(device))
	}
	return result
}
