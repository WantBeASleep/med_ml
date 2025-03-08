package device

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"

	pb "uzi/internal/generated/grpc/service"
	"uzi/internal/services"
)

type DeviceHandler interface {
	CreateDevice(ctx context.Context, in *pb.CreateDeviceIn) (*pb.CreateDeviceOut, error)
	GetDeviceList(ctx context.Context, _ *empty.Empty) (*pb.GetDeviceListOut, error)
}

type handler struct {
	services *services.Services
}

func New(
	services *services.Services,
) DeviceHandler {
	return &handler{
		services: services,
	}
}
