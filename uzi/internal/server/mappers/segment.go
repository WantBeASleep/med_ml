package mappers

import (
	"uzi/internal/domain"
	pb "uzi/internal/generated/grpc/service"
)

func SegmentFromDomain(domain domain.Segment) *pb.Segment {
	return &pb.Segment{
		Id:        domain.Id.String(),
		ImageId:   domain.ImageID.String(),
		NodeId:    domain.NodeID.String(),
		Contor:    domain.Contor,
		Tirads_23: domain.Tirads23,
		Tirads_4:  domain.Tirads4,
		Tirads_5:  domain.Tirads5,
	}
}

func SliceSegmentFromDomain(domains []domain.Segment) []*pb.Segment {
	pbs := make([]*pb.Segment, 0, len(domains))
	for _, d := range domains {
		pbs = append(pbs, SegmentFromDomain(d))
	}
	return pbs
}
