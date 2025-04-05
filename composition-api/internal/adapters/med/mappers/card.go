package mappers

import (
	"github.com/google/uuid"

	domain "composition-api/internal/domain/med"
	pb "composition-api/internal/generated/grpc/clients/med"
)

type Card struct{}

func (m Card) Domain(pb *pb.Card) domain.Card {
	return domain.Card{
		DoctorID:  uuid.MustParse(pb.DoctorId),
		PatientID: uuid.MustParse(pb.PatientId),
		Diagnosis: pb.Diagnosis,
	}
}

func (m Card) SliceDomain(pbs []*pb.Card) []domain.Card {
	return slice(pbs, m)
}
