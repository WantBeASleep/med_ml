package mappers

import (
	"med/internal/domain"
	pb "med/internal/generated/grpc/service"
)

func CardFromDomain(domain domain.Card) *pb.Card {
	card := &pb.Card{
		DoctorId:  domain.DoctorID.String(),
		PatientId: domain.PatientID.String(),
		Diagnosis: domain.Diagnosis,
	}
	if domain.ID != nil {
		id := int32(*domain.ID)
		card.Id = &id
	}
	return card
}
