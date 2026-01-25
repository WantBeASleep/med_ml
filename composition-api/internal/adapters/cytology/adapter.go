package cytology

import (
	"context"

	"github.com/google/uuid"

	domain "composition-api/internal/domain/cytology"
	pb "composition-api/internal/generated/grpc/clients/cytology"
)

type Adapter interface {
	// CYTOLOGY IMAGE
	CreateCytologyImage(ctx context.Context, in CreateCytologyImageIn) (uuid.UUID, error)
	GetCytologyImageById(ctx context.Context, id uuid.UUID) (domain.CytologyImage, *domain.OriginalImage, error)
	GetCytologyImagesByExternalId(ctx context.Context, id uuid.UUID) ([]domain.CytologyImage, error)
	GetCytologyImagesByDoctorIdAndPatientId(ctx context.Context, doctorID, patientID uuid.UUID) ([]domain.CytologyImage, error)
	UpdateCytologyImage(ctx context.Context, in UpdateCytologyImageIn) (domain.CytologyImage, error)
	DeleteCytologyImage(ctx context.Context, id uuid.UUID) error
	CopyCytologyImage(ctx context.Context, id uuid.UUID) (domain.CytologyImage, error)
	GetCytologyImageHistory(ctx context.Context, id uuid.UUID) ([]domain.CytologyImage, error)

	// ORIGINAL IMAGE
	CreateOriginalImage(ctx context.Context, in CreateOriginalImageIn) (uuid.UUID, error)
	GetOriginalImageById(ctx context.Context, id uuid.UUID) (domain.OriginalImage, error)
	GetOriginalImagesByCytologyId(ctx context.Context, id uuid.UUID) ([]domain.OriginalImage, error)
	UpdateOriginalImage(ctx context.Context, in UpdateOriginalImageIn) (domain.OriginalImage, error)

	// SEGMENTATION GROUP
	CreateSegmentationGroup(ctx context.Context, in CreateSegmentationGroupIn) (int, error)
	GetSegmentationGroupsByCytologyId(ctx context.Context, id uuid.UUID, segType *domain.SegType, groupType *domain.GroupType, isAI *bool) ([]domain.SegmentationGroup, error)
	UpdateSegmentationGroup(ctx context.Context, in UpdateSegmentationGroupIn) (domain.SegmentationGroup, error)
	DeleteSegmentationGroup(ctx context.Context, id int) error

	// SEGMENTATION
	CreateSegmentation(ctx context.Context, in CreateSegmentationIn) (int, error)
	GetSegmentationById(ctx context.Context, id int) (domain.Segmentation, error)
	GetSegmentsByGroupId(ctx context.Context, id int) ([]domain.Segmentation, error)
	UpdateSegmentation(ctx context.Context, in UpdateSegmentationIn) (domain.Segmentation, error)
	DeleteSegmentation(ctx context.Context, id int) error
}

type adapter struct {
	client pb.CytologySrvClient
}

func NewAdapter(client pb.CytologySrvClient) Adapter {
	return &adapter{client: client}
}
