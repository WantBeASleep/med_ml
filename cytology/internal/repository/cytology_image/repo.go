package cytology_image

import (
	daolib "github.com/WantBeASleep/med_ml_lib/dao"
	"github.com/google/uuid"

	"cytology/internal/repository/cytology_image/entity"
)

const (
	table = "cytology_image"

	columnID                = "id"
	columnExternalID        = "external_id"
	columnDoctorID          = "doctor_id"
	columnPatientID         = "patient_id"
	columnDiagnosticNumber  = "diagnostic_number"
	columnDiagnosticMarking = "diagnostic_marking"
	columnMaterialType      = "material_type"
	columnDiagnosDate       = "diagnos_date"
	columnIsLast            = "is_last"
	columnCalcitonin        = "calcitonin"
	columnCalcitoninInFlush = "calcitonin_in_flush"
	columnThyroglobulin     = "thyroglobulin"
	columnDetails           = "details"
	columnPrevID            = "prev_id"
	columnParentPrevID      = "parent_prev_id"
	columnCreateAt          = "create_at"
)

type Repository interface {
	CheckExist(id uuid.UUID) (bool, error)
	InsertCytologyImage(img entity.CytologyImage) error
	GetCytologyImageByID(id uuid.UUID) (entity.CytologyImage, error)
	GetCytologyImagesByExternalID(externalID uuid.UUID) ([]entity.CytologyImage, error)
	GetCytologyImagesByDoctorIdAndPatientId(doctorID, patientID uuid.UUID) ([]entity.CytologyImage, error)
	GetCytologyImagesByParentPrevID(parentPrevID uuid.UUID) ([]entity.CytologyImage, error)
	UpdateCytologyImage(img entity.CytologyImage) error
	DeleteCytologyImage(id uuid.UUID) error
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
