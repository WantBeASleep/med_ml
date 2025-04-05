package patient

import (
	"context"
	"fmt"

	"med/internal/domain"
	pentity "med/internal/repository/patient/entity"

	"github.com/google/uuid"
)

func (s *service) UpdatePatient(
	ctx context.Context,
	id uuid.UUID,
	update UpdatePatient,
) (domain.Patient, error) {
	patient, err := s.GetPatient(ctx, id)
	if err != nil {
		return domain.Patient{}, fmt.Errorf("get patient: %w", err)
	}
	update.Update(&patient)

	if err := s.dao.NewPatientQuery(ctx).UpdatePatient(pentity.Patient{}.FromDomain(patient)); err != nil {
		return domain.Patient{}, fmt.Errorf("update patient: %w", err)
	}

	return patient, nil
}
