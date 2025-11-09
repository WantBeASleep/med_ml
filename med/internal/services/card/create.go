package card

import (
	"context"
	"errors"
	"fmt"

	"med/internal/domain"
	centity "med/internal/repository/card/entity"
	"med/internal/repository/entity"
	"med/internal/services/validation"
)

func (s *service) CreateCard(ctx context.Context, card domain.Card) error {
	// Проверяем, что пациент существует и имеет валидный ОМС
	patient, err := s.dao.NewPatientQuery(ctx).GetPatientByID(card.PatientID)
	if err != nil {
		if errors.Is(err, entity.ErrNotFound) {
			return domain.ErrNotFound
		}
		return fmt.Errorf("get patient: %w", err)
	}

	// Проверка валидности ОМС пациента
	if !validation.ValidatePolicy(patient.Policy) {
		return domain.ErrBadRequest
	}

	if err := s.dao.NewCardQuery(ctx).InsertCard(centity.Card{}.FromDomain(card)); err != nil {
		var dbErr *entity.DBConflictError
		if errors.As(err, &dbErr) {
			return domain.ErrConflict
		}
		var valErr *entity.DBValidationError
		if errors.As(err, &valErr) {
			return domain.ErrUnprocessableEntity
		}
		return fmt.Errorf("insert card: %w", err)
	}

	return nil
}
