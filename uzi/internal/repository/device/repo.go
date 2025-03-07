package device

import (
	"uzi/internal/repository/device/entity"

	daolib "github.com/WantBeASleep/med_ml_lib/dao"
)

const (
	table = "device"

	columnId   = "id"
	columnName = "name"
)

type Repository interface {
	CreateDevice(name string) (int, error)

	GetDeviceList() ([]entity.Device, error)
}

type repo struct {
	*daolib.BaseQuery
}

//TODO: костыль, иначе переписывать dao Либу, это не сегодня
func NewR() *repo {
	return &repo{}
}

func (q *repo) SetBaseQuery(baseQuery *daolib.BaseQuery) {
	q.BaseQuery = baseQuery
}
