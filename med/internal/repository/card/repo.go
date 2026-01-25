package card

import (
	daolib "github.com/WantBeASleep/med_ml_lib/dao"
	"github.com/google/uuid"

	centity "med/internal/repository/card/entity"
)

const (
	table           = "card"
	columnID        = "id"
	columnDoctorID  = "doctor_id"
	columnPatientID = "patient_id"
	columnDiagnosis = "diagnosis"
)

type Repository interface {
	InsertCard(card centity.Card) (int, error)

	GetCardByPK(doctorID uuid.UUID, patientID uuid.UUID) (centity.Card, error)

	UpdateCard(card centity.Card) error
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
