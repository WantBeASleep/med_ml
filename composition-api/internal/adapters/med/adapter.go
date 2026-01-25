package med

import (
	"context"

	domain "composition-api/internal/domain/med"
	pb "composition-api/internal/generated/grpc/clients/med"

	"github.com/google/uuid"
)

type Adapter interface {
	RegisterDoctor(ctx context.Context, doctor domain.Doctor) error
	GetDoctor(ctx context.Context, id uuid.UUID) (domain.Doctor, error)

	CreatePatient(ctx context.Context, arg CreatePatientArg) error
	GetPatient(ctx context.Context, id uuid.UUID) (domain.Patient, error)
	GetPatientsByDoctorID(ctx context.Context, id uuid.UUID, status *bool) ([]domain.Patient, error)
	UpdatePatient(ctx context.Context, arg UpdatePatientIn) (domain.Patient, error)

	CreateCard(ctx context.Context, card domain.Card) (domain.Card, error)
	GetCard(ctx context.Context, doctorID, patientID uuid.UUID) (domain.Card, error)
	UpdateCard(ctx context.Context, card domain.Card) (domain.Card, error)
}

type adapter struct {
	client pb.MedSrvClient
}

func NewAdapter(client pb.MedSrvClient) Adapter {
	return &adapter{client: client}
}
