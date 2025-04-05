package device

import (
	"context"
)

func (s *service) CreateDevice(ctx context.Context, deviceName string) (int, error) {
	id, err := s.dao.NewDeviceQuery(ctx).CreateDevice(deviceName)
	return id, err
}
