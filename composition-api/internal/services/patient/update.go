package patient

import (
	"context"

	"composition-api/internal/adapters/med"
	domain "composition-api/internal/domain/med"

	"github.com/google/uuid"
)

func (s *service) UpdatePatient(
	ctx context.Context,
	id uuid.UUID,
	update UpdatePatientArg,
) (domain.Patient, error) {
	patient, err := s.adapters.Med.UpdatePatient(ctx, med.UpdatePatientIn{
		Id:         id,
		Active:     update.Active,
		Malignancy: update.Malignancy,
	})
	if err != nil {
		return domain.Patient{}, err
	}

	return patient, nil
}
