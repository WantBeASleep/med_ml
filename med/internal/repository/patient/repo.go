package patient

import (
	entity "med/internal/repository/patient/entity"

	daolib "github.com/WantBeASleep/med_ml_lib/dao"
	"github.com/google/uuid"
)

const (
	table             = "patient"
	columnID          = "id"
	columnFullName    = "fullname"
	columnEmail       = "email"
	columnPolicy      = "policy"
	columnActive      = "active"
	columnMalignancy  = "malignancy"
	columnBirthDate   = "birth_date"
	columnLastUziDate = "last_uzi_date"
)

type Repository interface {
	InsertPatient(patient entity.Patient) error

	GetPatientByID(id uuid.UUID) (entity.Patient, error)
	GetPatientsByDoctorID(id uuid.UUID) ([]entity.Patient, error)

	UpdatePatient(patient entity.Patient) error
}

type repo struct {
	*daolib.BaseQuery
}

func NewR() *repo {
	return &repo{}
}

func (r *repo) SetBaseQuery(baseQuery *daolib.BaseQuery) {
	r.BaseQuery = baseQuery
}
