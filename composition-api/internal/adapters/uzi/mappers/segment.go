package mappers

import (
	"github.com/google/uuid"

	domain "composition-api/internal/domain/uzi"
	pb "composition-api/internal/generated/grpc/clients/uzi"
)

type Segment struct{}

func (m Segment) Domain(pb *pb.Segment) domain.Segment {
	return domain.Segment{
		Id:       uuid.MustParse(pb.Id),
		ImageID:  uuid.MustParse(pb.ImageId),
		NodeID:   uuid.MustParse(pb.NodeId),
		Contor:   pb.Contor,
		Ai:       pb.Ai,
		Tirads23: pb.Tirads_23,
		Tirads4:  pb.Tirads_4,
		Tirads5:  pb.Tirads_5,
	}
}

func (m Segment) SliceDomain(pbs []*pb.Segment) []domain.Segment {
	return slice(pbs, m)
}
