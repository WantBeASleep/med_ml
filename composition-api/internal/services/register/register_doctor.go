package register

import (
	"context"

	auth_domain "composition-api/internal/domain/auth"
	med_domain "composition-api/internal/domain/med"

	"github.com/google/uuid"
)

func (s *service) RegisterDoctor(ctx context.Context, arg RegisterDoctorArg) (uuid.UUID, error) {
	id, err := s.adapters.Auth.RegisterUser(ctx, arg.Email, arg.Password, auth_domain.RoleDoctor)
	if err != nil {
		return uuid.UUID{}, err
	}

	if err := s.adapters.Med.RegisterDoctor(ctx, med_domain.Doctor{
		Id:          id,
		FullName:    arg.FullName,
		Org:         arg.Org,
		Job:         arg.Job,
		Description: arg.Description,
	}); err != nil {
		return uuid.UUID{}, err
	}

	return id, nil
}
