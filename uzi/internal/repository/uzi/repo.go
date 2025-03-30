package uzi

import (
	"uzi/internal/repository/uzi/entity"

	daolib "github.com/WantBeASleep/med_ml_lib/dao"
	"github.com/google/uuid"
)

const (
	uziTable = "uzi"

	columnID         = "id"
	columnProjection = "projection"
	columnChecked    = "checked"
	columnExternalID = "external_id"
	columnAuthor     = "author"
	columnDeviceID   = "device_id"
	columnStatus     = "status"
	columnCreateAt   = "create_at"
)

type Repository interface {
	CheckExist(id uuid.UUID) (bool, error)

	InsertUzi(uzi entity.Uzi) error

	GetUziByID(id uuid.UUID) (entity.Uzi, error)
	GetUzisByExternalID(externalID uuid.UUID) ([]entity.Uzi, error)
	GetUzisByAuthor(author uuid.UUID) ([]entity.Uzi, error)

	UpdateUzi(uzi entity.Uzi) error
	UpdateUziStatus(id uuid.UUID, status string) error
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
