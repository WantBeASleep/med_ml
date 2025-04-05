package patient

import (
	"context"

	domain "composition-api/internal/domain/med"

	"github.com/google/uuid"
)

func (s *service) GetPatient(ctx context.Context, id uuid.UUID) (domain.Patient, error) {
	patient, err := s.adapters.Med.GetPatient(ctx, id)
	if err != nil {
		return domain.Patient{}, err
	}

	return patient, nil
}

func (s *service) GetPatientsByDoctorID(ctx context.Context, doctorID uuid.UUID) ([]domain.Patient, error) {
	patients, err := s.adapters.Med.GetPatientsByDoctorID(ctx, doctorID)
	if err != nil {
		return nil, err
	}

	return patients, nil
}
