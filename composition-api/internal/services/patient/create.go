package patient

import (
	"context"

	"composition-api/internal/adapters/med"
	"composition-api/internal/domain"
	"composition-api/internal/services/validation"

	"github.com/google/uuid"
)

func (s *service) CreatePatient(ctx context.Context, arg CreatePatientArg) (uuid.UUID, error) {
	// Проверка валидности ОМС перед созданием пользователя
	if !validation.ValidatePolicy(arg.Policy) {
		return uuid.UUID{}, domain.ErrBadRequest
	}

	id, err := s.adapters.Auth.CreateUnRegisteredUser(ctx, arg.Email)
	if err != nil {
		return uuid.UUID{}, err
	}

	err = s.adapters.Med.CreatePatient(ctx, med.CreatePatientArg{
		Id:         id,
		FullName:   arg.Fullname,
		Email:      arg.Email,
		Policy:     arg.Policy,
		Active:     arg.Active,
		Malignancy: arg.Malignancy,
		BirthDate:  arg.BirthDate,
	})
	if err != nil {
		return uuid.UUID{}, err
	}

	return id, nil
}
