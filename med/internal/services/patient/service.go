package patient

import (
	"context"

	"med/internal/domain"
	"med/internal/repository"

	"github.com/google/uuid"
)

type Service interface {
	InsertPatient(ctx context.Context, patient domain.Patient) error

	GetPatient(ctx context.Context, id uuid.UUID) (domain.Patient, error)
	GetPatientsByDoctorID(ctx context.Context, doctorID uuid.UUID) ([]domain.Patient, error)

	UpdatePatient(ctx context.Context, id uuid.UUID, update UpdatePatient) (domain.Patient, error)
}

type service struct {
	dao repository.DAO
}

func New(
	dao repository.DAO,
) Service {
	return &service{
		dao: dao,
	}
}
