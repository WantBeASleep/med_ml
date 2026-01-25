package original_image

import (
	daolib "github.com/WantBeASleep/med_ml_lib/dao"
	"github.com/google/uuid"

	"cytology/internal/repository/original_image/entity"
)

const (
	table = "original_image"

	columnID         = "id"
	columnCytologyID = "cytology_id"
	columnImagePath  = "image_path"
	columnCreateDate = "create_date"
	columnDelayTime  = "delay_time"
	columnViewedFlag = "viewed_flag"
)

type Repository interface {
	InsertOriginalImage(img entity.OriginalImage) error
	GetOriginalImageByID(id uuid.UUID) (entity.OriginalImage, error)
	GetOriginalImagesByCytologyID(cytologyID uuid.UUID) ([]entity.OriginalImage, error)
	UpdateOriginalImage(img entity.OriginalImage) error
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
