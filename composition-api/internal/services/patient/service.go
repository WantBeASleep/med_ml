package patient

import (
	"context"

	"composition-api/internal/adapters"
	domain "composition-api/internal/domain/med"

	"github.com/google/uuid"
)

type Service interface {
	CreatePatient(ctx context.Context, arg CreatePatientArg) (uuid.UUID, error)

	GetPatient(ctx context.Context, id uuid.UUID) (domain.Patient, error)
	GetPatientsByDoctorID(ctx context.Context, doctorID uuid.UUID) ([]domain.Patient, error)

	UpdatePatient(ctx context.Context, id uuid.UUID, update UpdatePatientArg) (domain.Patient, error)
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
