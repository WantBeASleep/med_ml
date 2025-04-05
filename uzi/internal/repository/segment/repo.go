package segment

import (
	"uzi/internal/repository/segment/entity"

	daolib "github.com/WantBeASleep/med_ml_lib/dao"
	"github.com/google/uuid"
)

const (
	table = "segment"

	columnID       = "id"
	columnNodeID   = "node_id"
	columnImageID  = "image_id"
	columnContor   = "contor"
	columnAi       = "ai"
	columnTirads23 = "tirads_23"
	columnTirads4  = "tirads_4"
	columnTirads5  = "tirads_5"
)

type Repository interface {
	InsertSegments(segments ...entity.Segment) error

	GetSegmentByID(id uuid.UUID) (entity.Segment, error)
	GetSegmentsByNodeID(id uuid.UUID) ([]entity.Segment, error)
	GetSegmentsByImageID(id uuid.UUID) ([]entity.Segment, error)

	UpdateSegment(segment entity.Segment) error

	DeleteSegmentByID(id uuid.UUID) error
	DeleteSegmentsByUziID(id uuid.UUID) error
}

type repo struct {
	*daolib.BaseQuery
}

func NewRepo() *repo {
	return &repo{}
}

func (q *repo) SetBaseQuery(baseQuery *daolib.BaseQuery) {
	q.BaseQuery = baseQuery
}
