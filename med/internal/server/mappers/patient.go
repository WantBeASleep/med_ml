package mappers

import (
	"time"

	"med/internal/domain"
	pb "med/internal/generated/grpc/service"

	"github.com/AlekSi/pointer"
)

func PatientFromDomain(domain domain.Patient) *pb.Patient {
	var lastUziDate *string
	if domain.LastUziDate != nil {
		lastUziDate = pointer.ToString(domain.LastUziDate.Format(time.RFC3339))
	}

	return &pb.Patient{
		Id:          domain.Id.String(),
		Fullname:    domain.FullName,
		Email:       domain.Email,
		Policy:      domain.Policy,
		Active:      domain.Active,
		Malignancy:  domain.Malignancy,
		BirthDate:   domain.BirthDate.Format(time.RFC3339),
		LastUziDate: lastUziDate,
	}
}

func SlicePatientFromDomain(domains []domain.Patient) []*pb.Patient {
	pbs := make([]*pb.Patient, 0, len(domains))
	for _, d := range domains {
		pbs = append(pbs, PatientFromDomain(d))
	}
	return pbs
}
