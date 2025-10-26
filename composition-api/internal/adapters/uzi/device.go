package uzi

import (
	"context"

	adapter_errors "composition-api/internal/adapters/errors"
	"composition-api/internal/adapters/uzi/mappers"
	domain "composition-api/internal/domain/uzi"
	pb "composition-api/internal/generated/grpc/clients/uzi"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (a *adapter) CreateDevice(ctx context.Context, name string) (int, error) {
	res, err := a.client.CreateDevice(ctx, &pb.CreateDeviceIn{Name: name})
	if err != nil {
		return 0, err
	}

	return int(res.Id), nil
}

func (a *adapter) GetDeviceList(ctx context.Context) ([]domain.Device, error) {
	res, err := a.client.GetDeviceList(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, adapter_errors.HandleGRPCError(err)
	}

	return mappers.Device{}.SliceDomain(res.Devices), nil
}
