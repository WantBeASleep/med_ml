package cytology_image

import (
	"context"

	"cytology/internal/domain"
	"cytology/internal/repository"

	"github.com/google/uuid"
)

type Service interface {
	CreateCytologyImage(ctx context.Context, arg CreateCytologyImageArg) (uuid.UUID, error)
	GetCytologyImageByID(ctx context.Context, id uuid.UUID) (domain.CytologyImage, error)
	GetCytologyImagesByExternalID(ctx context.Context, externalID uuid.UUID) ([]domain.CytologyImage, error)
	GetCytologyImagesByDoctorIdAndPatientId(ctx context.Context, doctorID, patientID uuid.UUID) ([]domain.CytologyImage, error)
	UpdateCytologyImage(ctx context.Context, arg UpdateCytologyImageArg) (domain.CytologyImage, error)
	DeleteCytologyImage(ctx context.Context, id uuid.UUID) error
	CopyCytologyImage(ctx context.Context, id uuid.UUID) (domain.CytologyImage, error)
	GetCytologyImageHistory(ctx context.Context, id uuid.UUID) ([]domain.CytologyImage, error)
}

type service struct {
	dao repository.DAO
}

func New(dao repository.DAO) Service {
	return &service{
		dao: dao,
	}
}
