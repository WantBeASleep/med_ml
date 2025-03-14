package uzi

import (
	"context"

	"github.com/google/uuid"

	domain "gateway/internal/domain/uzi"
	pb "gateway/internal/generated/grpc/clients/uzi"
)

type Adapter interface {
	// DEVICE
	CreateDevice(ctx context.Context, name string) (int, error)
	GetDeviceList(ctx context.Context) ([]domain.Device, error)
	// UZI
	CreateUzi(ctx context.Context, in CreateUziIn) (uuid.UUID, error)
	GetUziById(ctx context.Context, id uuid.UUID) (domain.Uzi, error)
	GetUzisByExternalId(ctx context.Context, id uuid.UUID) ([]domain.Uzi, error)
	GetEchographicByUziId(ctx context.Context, id uuid.UUID) (domain.Echographic, error)
	UpdateUzi(ctx context.Context, in UpdateUziIn) (domain.Uzi, error)
	UpdateEchographic(ctx context.Context, in domain.Echographic) (domain.Echographic, error)
	// IMAGE
	GetImagesByUziId(ctx context.Context, id uuid.UUID) ([]domain.Image, error)
	// NODE
	GetNodesByUziId(ctx context.Context, id uuid.UUID) ([]domain.Node, error)
	UpdateNode(ctx context.Context, in UpdateNodeIn) (domain.Node, error)
	// SEGMENT
	CreateSegment(ctx context.Context, in CreateSegmentIn) (uuid.UUID, error)
	GetSegmentsByNodeId(ctx context.Context, id uuid.UUID) ([]domain.Segment, error)
	UpdateSegment(ctx context.Context, in UpdateSegmentIn) (domain.Segment, error)
	// доменные области слишком сильно пересекаются, вынесено в одну надобласть
	// NODE-SEGMENT
	CreateNodeWithSegments(ctx context.Context, in CreateNodeWithSegmentsIn) (uuid.UUID, []uuid.UUID, error)
	GetNodesWithSegmentsByImageId(ctx context.Context, id uuid.UUID) ([]domain.Node, []domain.Segment, error)
	DeleteNode(ctx context.Context, id uuid.UUID) error
	DeleteSegment(ctx context.Context, id uuid.UUID) error
}

type adapter struct {
	client pb.UziSrvClient
}

func NewAdapter(client pb.UziSrvClient) Adapter {
	return &adapter{client: client}
}
