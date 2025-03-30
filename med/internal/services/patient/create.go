package patient

import (
	"context"
	"fmt"

	"med/internal/domain"
	"med/internal/repository/patient/entity"

	"github.com/google/uuid"
)

func (s *service) CreatePatient(ctx context.Context, patient domain.Patient) (uuid.UUID, error) {
	patient.Id = uuid.New()
	if err := s.dao.NewPatientQuery(ctx).InsertPatient(entity.Patient{}.FromDomain(patient)); err != nil {
		return uuid.Nil, fmt.Errorf("insert patient: %w", err)
	}

	return patient.Id, nil
}
