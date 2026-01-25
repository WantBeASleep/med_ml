package cytology

import (
	"time"

	"github.com/google/uuid"
)

type DiagnosticMarking string

const (
	DiagnosticMarkingP11 DiagnosticMarking = "П11"
	DiagnosticMarkingL23 DiagnosticMarking = "Л23"
)

type MaterialType string

const (
	MaterialTypeGS  MaterialType = "GS"
	MaterialTypeBP  MaterialType = "BP"
	MaterialTypeTP  MaterialType = "TP"
	MaterialTypePTP MaterialType = "PTP"
	MaterialTypeLNP MaterialType = "LNP"
)

type CytologyImage struct {
	Id                uuid.UUID
	ExternalID        uuid.UUID
	DoctorID          uuid.UUID
	PatientID         uuid.UUID
	DiagnosticNumber  int
	DiagnosticMarking *DiagnosticMarking
	MaterialType      *MaterialType
	DiagnosDate       time.Time
	IsLast            bool
	Calcitonin        *int
	CalcitoninInFlush *int
	Thyroglobulin     *int
	Details           *string
	PrevID            *uuid.UUID
	ParentPrevID      *uuid.UUID
	CreateAt          time.Time
}
