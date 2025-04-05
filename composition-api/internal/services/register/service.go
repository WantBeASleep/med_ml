package register

import (
	"context"

	"composition-api/internal/adapters"

	"github.com/google/uuid"
)

type Service interface {
	RegisterDoctor(ctx context.Context, arg RegisterDoctorArg) (uuid.UUID, error)
	RegisterPatient(ctx context.Context, arg RegisterPatientArg) (uuid.UUID, error)
}

type service struct {
	adapters *adapters.Adapters
}

func New(
	adapters *adapters.Adapters,
) Service {
	return &service{
		adapters: adapters,
	}
}
