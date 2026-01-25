package segmentation_group

import (
	daolib "github.com/WantBeASleep/med_ml_lib/dao"
	"github.com/google/uuid"

	"cytology/internal/repository/segmentation_group/entity"
)

const (
	table = "segmentation_group"

	columnID         = "id"
	columnCytologyID = "cytology_id"
	columnSegType    = "seg_type"
	columnGroupType  = "group_type"
	columnIsAI       = "is_ai"
	columnDetails    = "details"
	columnCreateAt   = "create_at"
)

type Repository interface {
	InsertSegmentationGroup(group entity.SegmentationGroup) (int, error)
	GetSegmentationGroupByID(id int) (entity.SegmentationGroup, error)
	GetSegmentationGroupsByCytologyID(cytologyID uuid.UUID) ([]entity.SegmentationGroup, error)
	UpdateSegmentationGroup(group entity.SegmentationGroup) error
	DeleteSegmentationGroup(id int) error
}

type repo struct {
	*daolib.BaseQuery
}

func NewR() *repo {
	return &repo{}
}

func (q *repo) SetBaseQuery(baseQuery *daolib.BaseQuery) {
	q.BaseQuery = baseQuery
}
