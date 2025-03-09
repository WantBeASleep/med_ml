package mappers

import (
	"github.com/google/uuid"

	domain "gateway/internal/domain/uzi"
	pb "gateway/internal/generated/grpc/clients/uzi"
)

func Node(pb *pb.Node) domain.Node {
	return domain.Node{
		Id:        uuid.MustParse(pb.Id),
		Ai:        pb.Ai,
		UziID:     uuid.MustParse(pb.UziId),
		Tirads23:  pb.Tirads_23,
		Tirads4:   pb.Tirads_4,
		Tirads5:   pb.Tirads_5,
	}
}

func SliceNode(pbs []*pb.Node) []domain.Node {
	domains := make([]domain.Node, 0, len(pbs))
	for _, pb := range pbs {
		domains = append(domains, Node(pb))
	}
	return domains
}
