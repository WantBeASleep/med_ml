package mappers

import (
	"github.com/google/uuid"

	domain "gateway/internal/domain/uzi"
	pb "gateway/internal/generated/grpc/clients/uzi"
)

func Image(pb *pb.Image) domain.Image {
	return domain.Image{
		Id:    uuid.MustParse(pb.Id),
		UziID: uuid.MustParse(pb.UziId),
		Page:  int(pb.Page),
	}
}

func SliceImage(pbs []*pb.Image) []domain.Image {
	domains := make([]domain.Image, 0, len(pbs))
	for _, pb := range pbs {
		domains = append(domains, Image(pb))
	}
	return domains
}
