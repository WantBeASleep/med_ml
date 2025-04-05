package patient

import (
	"context"

	"med/internal/domain"
	"med/internal/repository/patient/entity"
)

func (s *service) InsertPatient(ctx context.Context, patient domain.Patient) error {
	err := s.dao.NewPatientQuery(ctx).InsertPatient(entity.Patient{}.FromDomain(patient))
	return err
}
