package mappers

import (
	domain "composition-api/internal/domain/uzi"
	pb "composition-api/internal/generated/grpc/clients/uzi"
)

func Device(pb *pb.Device) domain.Device {
	return domain.Device{
		Id:   int(pb.Id),
		Name: pb.Name,
	}
}

func SliceDevice(pbs []*pb.Device) []domain.Device {
	domains := make([]domain.Device, 0, len(pbs))
	for _, p := range pbs {
		domains = append(domains, Device(p))
	}
	return domains
}
