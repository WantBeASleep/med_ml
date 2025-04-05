package mappers

import (
	"time"

	"github.com/google/uuid"

	domain "composition-api/internal/domain/med"
	pb "composition-api/internal/generated/grpc/clients/med"
)

type Patient struct{}

// TODO: обрабатывать ошибки
func (m Patient) Domain(pb *pb.Patient) domain.Patient {
	birthDate, _ := time.Parse(time.RFC3339, pb.BirthDate)
	var lastUziDate *time.Time
	if pb.LastUziDate != nil {
		lastUziDateParsed, _ := time.Parse(time.RFC3339, *pb.LastUziDate)
		lastUziDate = &lastUziDateParsed
	}
	return domain.Patient{
		Id:          uuid.MustParse(pb.Id),
		FullName:    pb.Fullname,
		Email:       pb.Email,
		Policy:      pb.Policy,
		Active:      pb.Active,
		Malignancy:  pb.Malignancy,
		BirthDate:   birthDate,
		LastUziDate: lastUziDate,
	}
}

func (m Patient) SliceDomain(pbs []*pb.Patient) []domain.Patient {
	return slice(pbs, m)
}
