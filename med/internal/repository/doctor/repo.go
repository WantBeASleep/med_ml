package doctor

import (
	daolib "github.com/WantBeASleep/med_ml_lib/dao"
	"github.com/google/uuid"

	dentity "med/internal/repository/doctor/entity"
)

const (
	table             = "doctor"
	columnID          = "id"
	columnFullname    = "fullname"
	columnOrg         = "org"
	columnJob         = "job"
	columnDescription = "description"
)

type Repository interface {
	InsertDoctor(doctor dentity.Doctor) error

	GetDoctorByID(id uuid.UUID) (dentity.Doctor, error)

	UpdateDoctor(doctor dentity.Doctor) error
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
