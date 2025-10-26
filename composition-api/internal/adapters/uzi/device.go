package uzi

import (
	"context"
	"fmt"

	adapter_errors "composition-api/internal/adapters/errors"
	"composition-api/internal/adapters/uzi/mappers"
	domain "composition-api/internal/domain/uzi"
	pb "composition-api/internal/generated/grpc/clients/uzi"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
		st, ok := status.FromError(err)
		if !ok {
			return nil, fmt.Errorf("unknown error: %w", err)
		}

		switch st.Code() {
		case codes.NotFound:
			return nil, adapter_errors.ErrNotFound
		default:
			return nil, err
		}
	}

	return mappers.Device{}.SliceDomain(res.Devices), nil
}
