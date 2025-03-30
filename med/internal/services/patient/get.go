package patient

import (
	"context"
	"fmt"

	"med/internal/domain"
	pentity "med/internal/repository/patient/entity"

	"github.com/google/uuid"
)

func (s *service) GetPatient(ctx context.Context, id uuid.UUID) (domain.Patient, error) {
	patient, err := s.dao.NewPatientQuery(ctx).GetPatientByID(id)
	if err != nil {
		return domain.Patient{}, fmt.Errorf("get patient: %w", err)
	}

	return patient.ToDomain(), nil
}

func (s *service) GetPatientsByDoctorID(ctx context.Context, doctorID uuid.UUID) ([]domain.Patient, error) {
	patients, err := s.dao.NewPatientQuery(ctx).GetPatientsByDoctorID(doctorID)
	if err != nil {
		return nil, fmt.Errorf("get doctor patients: %w", err)
	}

	return pentity.Patient{}.SliceToDomain(patients), nil
}
