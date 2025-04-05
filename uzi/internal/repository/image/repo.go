package image

import (
	"uzi/internal/repository/image/entity"

	daolib "github.com/WantBeASleep/med_ml_lib/dao"
	"github.com/google/uuid"
)

const (
	table = "image"

	columnId    = "id"
	columnUziId = "uzi_id"
	columnPage  = "page"
)

type Repository interface {
	InsertImages(images ...entity.Image) error

	GetImagesByUziID(uziID uuid.UUID) ([]entity.Image, error)
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
