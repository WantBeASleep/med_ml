package patient

import (
	"context"
	"errors"
	"fmt"

	"med/internal/domain"
	repoEntity "med/internal/repository/entity"
	"med/internal/repository/patient/entity"
	"med/internal/services/validation"
)

func (s *service) InsertPatient(ctx context.Context, patient domain.Patient) error {
	// Проверка валидности ОМС
	if !validation.ValidatePolicy(patient.Policy) {
		return fmt.Errorf("%w: неверный формат ОМС", domain.ErrBadRequest)
	}

	err := s.dao.NewPatientQuery(ctx).InsertPatient(entity.Patient{}.FromDomain(patient))
	if err != nil {
		switch {
		case errors.Is(err, repoEntity.ErrConflict):
			return domain.ErrConflict
		case errors.Is(err, repoEntity.ErrValidation):
			return domain.ErrUnprocessableEntity
		default:
			return err
		}
	}
	return nil
}
