package segmentation

import (
	daolib "github.com/WantBeASleep/med_ml_lib/dao"

	"cytology/internal/repository/segmentation/entity"
)

const (
	table = "segmentation"

	columnID                  = "id"
	columnSegmentationGroupID = "segmentation_group_id"
	columnCreateAt            = "create_at"
)

const (
	pointTable = "segmentation_point"

	pointColumnID             = "id"
	pointColumnSegmentationID = "segmentation_id"
	pointColumnX              = "x"
	pointColumnY              = "y"
	pointColumnUID            = "uid"
	pointColumnCreateAt       = "create_at"
)

type Repository interface {
	InsertSegmentation(seg entity.Segmentation) (int, error)
	GetSegmentationByID(id int) (entity.Segmentation, error)
	GetSegmentsByGroupID(groupID int) ([]entity.Segmentation, error)
	UpdateSegmentation(seg entity.Segmentation) error
	DeleteSegmentation(id int) error
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
