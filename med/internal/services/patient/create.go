package patient

import (
	"context"
	"strings"

	"med/internal/domain"
	"med/internal/repository/patient/entity"
	"med/internal/services/validation"
)

func (s *service) InsertPatient(ctx context.Context, patient domain.Patient) error {
	// Проверка валидности ОМС
	if !validation.ValidatePolicy(patient.Policy) {
		return domain.ErrBadRequest
	}

	err := s.dao.NewPatientQuery(ctx).InsertPatient(entity.Patient{}.FromDomain(patient))
	if err != nil {
		if strings.Contains(err.Error(), "validation") || strings.Contains(err.Error(), "constraint") || strings.Contains(err.Error(), "check") {
			return domain.ErrUnprocessableEntity
		}
		return err
	}
	return nil
}
