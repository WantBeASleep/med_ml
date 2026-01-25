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

func (s *service) CreateCard(ctx context.Context, card domain.Card) (domain.Card, error) {
	// Проверяем, что пациент существует и имеет валидный ОМС
	patient, err := s.dao.NewPatientQuery(ctx).GetPatientByID(card.PatientID)
	if err != nil {
		if errors.Is(err, entity.ErrNotFound) {
			return domain.Card{}, domain.ErrNotFound
		}
		return domain.Card{}, fmt.Errorf("get patient: %w", err)
	}

	// Проверка валидности ОМС пациента
	if !validation.ValidatePolicy(patient.Policy) {
		return domain.Card{}, domain.ErrBadRequest
	}

	id, err := s.dao.NewCardQuery(ctx).InsertCard(centity.Card{}.FromDomain(card))
	if err != nil {
		var dbErr *entity.DBConflictError
		if errors.As(err, &dbErr) {
			return domain.Card{}, domain.ErrConflict
		}
		var valErr *entity.DBValidationError
		if errors.As(err, &valErr) {
			return domain.Card{}, domain.ErrUnprocessableEntity
		}
		return domain.Card{}, fmt.Errorf("insert card: %w", err)
	}

	// Возвращаем созданную карту с ID
	card.ID = &id
	return card, nil
}
