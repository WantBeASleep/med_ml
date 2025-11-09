package register

import (
	"context"

	"composition-api/internal/adapters/med"
	"composition-api/internal/domain"
	auth_domain "composition-api/internal/domain/auth"
	"composition-api/internal/services/validation"

	"github.com/google/uuid"
)

func (s *service) RegisterPatient(ctx context.Context, arg RegisterPatientArg) (uuid.UUID, error) {
	// Проверка валидности ОМС перед созданием пользователя
	if !validation.ValidatePolicy(arg.Policy) {
		return uuid.UUID{}, domain.ErrBadRequest
	}

	id, err := s.adapters.Auth.RegisterUser(ctx, arg.Email, arg.Password, auth_domain.RolePatient)
	if err != nil {
		return uuid.UUID{}, err
	}

	if err := s.adapters.Med.CreatePatient(ctx, med.CreatePatientArg{
		Id:         id,
		FullName:   arg.FullName,
		Email:      arg.Email,
		Policy:     arg.Policy,
		Active:     true,
		Malignancy: false,
		BirthDate:  arg.BirthDate,
	}); err != nil {
		return uuid.UUID{}, err
	}

	return id, nil
}
