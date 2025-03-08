package mappers

import (
	"uzi/internal/domain"
	pb "uzi/internal/generated/grpc/service"
)

func DeviceFromDomain(domain domain.Device) *pb.Device {
	return &pb.Device{
		Id:   int64(domain.Id),
		Name: domain.Name,
	}
}

func SliceDeviceFromDomain(domains []domain.Device) []*pb.Device {
	pbs := make([]*pb.Device, 0, len(domains))
	for _, d := range domains {
		pbs = append(pbs, DeviceFromDomain(d))
	}
	return pbs
}
