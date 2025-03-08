package mappers

import (
	"uzi/internal/domain"
	pb "uzi/internal/generated/grpc/service"
)

func NodeFromDomain(domain domain.Node) *pb.Node {
	return &pb.Node{
		Id:        domain.Id.String(),
		Ai:        domain.Ai,
		UziId:     domain.UziID.String(),
		Tirads_23: domain.Tirads23,
		Tirads_4:  domain.Tirads4,
		Tirads_5:  domain.Tirads5,
	}
}

func SliceNodeFromDomain(domains []domain.Node) []*pb.Node {
	pbs := make([]*pb.Node, 0, len(domains))
	for _, d := range domains {
		pbs = append(pbs, NodeFromDomain(d))
	}
	return pbs
}
