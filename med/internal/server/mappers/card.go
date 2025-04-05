package mappers

import (
	"med/internal/domain"
	pb "med/internal/generated/grpc/service"
)

func CardFromDomain(domain domain.Card) *pb.Card {
	return &pb.Card{
		DoctorId:  domain.DoctorID.String(),
		PatientId: domain.PatientID.String(),
		Diagnosis: domain.Diagnosis,
	}
}
