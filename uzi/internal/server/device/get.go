package device

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "uzi/internal/generated/grpc/service"
	"uzi/internal/server/mappers"
)

func (h *handler) GetDeviceList(ctx context.Context, _ *empty.Empty) (*pb.GetDeviceListOut, error) {
	devices, err := h.services.Device.GetDeviceList(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}

	out := new(pb.GetDeviceListOut)
	out.Devices = mappers.SliceDeviceFromDomain(devices)

	return out, nil
}
