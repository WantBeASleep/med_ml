package mappers

import (
	"github.com/google/uuid"

	domain "composition-api/internal/domain/med"
	pb "composition-api/internal/generated/grpc/clients/med"
)

type Doctor struct{}

func (m Doctor) Domain(pb *pb.Doctor) domain.Doctor {
	return domain.Doctor{
		Id:          uuid.MustParse(pb.Id),
		FullName:    pb.Fullname,
		Org:         pb.Org,
		Job:         pb.Job,
		Description: pb.Description,
	}
}

func (m Doctor) SliceDomain(pbs []*pb.Doctor) []domain.Doctor {
	return slice(pbs, m)
}
