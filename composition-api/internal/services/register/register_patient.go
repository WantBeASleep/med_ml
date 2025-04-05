package register

import (
	"context"

	"composition-api/internal/adapters/med"
	auth_domain "composition-api/internal/domain/auth"

	"github.com/google/uuid"
)

func (s *service) RegisterPatient(ctx context.Context, arg RegisterPatientArg) (uuid.UUID, error) {
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
