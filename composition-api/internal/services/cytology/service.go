package cytology

import (
	"context"

	"github.com/google/uuid"

	"composition-api/internal/adapters"
	dbus "composition-api/internal/dbus/producers"
	domain "composition-api/internal/domain/cytology"
	"composition-api/internal/repository"
)

type Service interface {
	CreateCytologyImage(ctx context.Context, arg CreateCytologyImageArg) (uuid.UUID, error)
	GetCytologyImageById(ctx context.Context, id uuid.UUID) (domain.CytologyImage, *domain.OriginalImage, error)
	GetCytologyImagesByExternalId(ctx context.Context, externalID uuid.UUID) ([]domain.CytologyImage, error)
	GetCytologyImagesByDoctorIdAndPatientId(ctx context.Context, doctorID, patientID uuid.UUID) ([]domain.CytologyImage, error)
	UpdateCytologyImage(ctx context.Context, arg UpdateCytologyImageArg) (domain.CytologyImage, error)
	DeleteCytologyImage(ctx context.Context, id uuid.UUID) error

	CreateOriginalImage(ctx context.Context, arg CreateOriginalImageArg) (uuid.UUID, error)
	GetOriginalImageById(ctx context.Context, id uuid.UUID) (domain.OriginalImage, error)
	GetOriginalImagesByCytologyId(ctx context.Context, id uuid.UUID) ([]domain.OriginalImage, error)
	UpdateOriginalImage(ctx context.Context, arg UpdateOriginalImageArg) (domain.OriginalImage, error)

	CreateSegmentationGroup(ctx context.Context, arg CreateSegmentationGroupArg) (int, error)
	GetSegmentationGroupsByCytologyId(ctx context.Context, id uuid.UUID, segType *domain.SegType, groupType *domain.GroupType, isAI *bool) ([]domain.SegmentationGroup, error)
	UpdateSegmentationGroup(ctx context.Context, arg UpdateSegmentationGroupArg) (domain.SegmentationGroup, error)
	DeleteSegmentationGroup(ctx context.Context, id int) error

	CreateSegmentation(ctx context.Context, arg CreateSegmentationArg) (int, error)
	GetSegmentationById(ctx context.Context, id int) (domain.Segmentation, error)
	GetSegmentsByGroupId(ctx context.Context, id int) ([]domain.Segmentation, error)
	UpdateSegmentation(ctx context.Context, arg UpdateSegmentationArg) (domain.Segmentation, error)
	DeleteSegmentation(ctx context.Context, id int) error

	CopyCytologyImage(ctx context.Context, id uuid.UUID) (domain.CytologyImage, error)
	GetCytologyImageHistory(ctx context.Context, id uuid.UUID) ([]domain.CytologyImage, error)
}

type service struct {
	adapters *adapters.Adapters
	dao      repository.DAO
	dbus     dbus.Producer
}

func New(adapters *adapters.Adapters, dao repository.DAO, dbus dbus.Producer) Service {
	return &service{
		adapters: adapters,
		dao:      dao,
		dbus:     dbus,
	}
}
