package device

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "uzi/internal/generated/grpc/service"
)

func (h *handler) CreateDevice(ctx context.Context, in *pb.CreateDeviceIn) (*pb.CreateDeviceOut, error) {
	deviceID, err := h.services.Device.CreateDevice(ctx, in.Name)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: create device: %s", err.Error())
	}

	return &pb.CreateDeviceOut{Id: int64(deviceID)}, nil
}
