package mappers

import (
	domain "composition-api/internal/domain/uzi"
	pb "composition-api/internal/generated/grpc/clients/uzi"
)

type Device struct{}

func (m Device) Domain(pb *pb.Device) domain.Device {
	return domain.Device{
		Id:   int(pb.Id),
		Name: pb.Name,
	}
}

func (m Device) SliceDomain(pbs []*pb.Device) []domain.Device {
	return slice(pbs, m)
}
