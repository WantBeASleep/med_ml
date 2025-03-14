package mappers

import (
	"github.com/google/uuid"

	domain "gateway/internal/domain/uzi"
	pb "gateway/internal/generated/grpc/clients/uzi"
)

func Segment(pb *pb.Segment) domain.Segment {
	return domain.Segment{
		Id:       uuid.MustParse(pb.Id),
		ImageID:  uuid.MustParse(pb.ImageId),
		NodeID:   uuid.MustParse(pb.NodeId),
		Contor:   pb.Contor,
		Tirads23: pb.Tirads_23,
		Tirads4:  pb.Tirads_4,
		Tirads5:  pb.Tirads_5,
	}
}

func SliceSegment(pbs []*pb.Segment) []domain.Segment {
	domains := make([]domain.Segment, 0, len(pbs))
	for _, pb := range pbs {
		domains = append(domains, Segment(pb))
	}
	return domains
}
