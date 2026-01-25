package cytology

import (
	"github.com/google/uuid"
	ht "github.com/ogen-go/ogen/http"

	domain "composition-api/internal/domain/cytology"
)

type CreateCytologyImageArg struct {
	ExternalID        uuid.UUID
	DoctorID          uuid.UUID
	PatientID         uuid.UUID
	DiagnosticNumber  int
	DiagnosticMarking *domain.DiagnosticMarking
	MaterialType      *domain.MaterialType
	Calcitonin        *int
	CalcitoninInFlush *int
	Thyroglobulin     *int
	Details           *string
	PrevID            *uuid.UUID
	ParentPrevID      *uuid.UUID
	File              *ht.MultipartFile
	ContentType       string
}

type UpdateCytologyImageArg struct {
	Id                uuid.UUID
	DiagnosticMarking *domain.DiagnosticMarking
	MaterialType      *domain.MaterialType
	Calcitonin        *int
	CalcitoninInFlush *int
	Thyroglobulin     *int
	Details           *string
	IsLast            *bool
	PrevID            *uuid.UUID
	ParentPrevID      *uuid.UUID
}

type CreateOriginalImageArg struct {
	CytologyID  uuid.UUID
	File        ht.MultipartFile
	ContentType string
	DelayTime   *float64
}

type UpdateOriginalImageArg struct {
	Id         uuid.UUID
	DelayTime  *float64
	ViewedFlag *bool
}

type CreateSegmentationGroupArg struct {
	CytologyID uuid.UUID
	SegType    domain.SegType
	GroupType  domain.GroupType
	IsAI       bool
	Details    *string
}

type UpdateSegmentationGroupArg struct {
	Id      int
	SegType *domain.SegType
	Details *string
}

type CreateSegmentationArg struct {
	SegmentationGroupID int
	Points              []domain.SegmentationPoint
}

type UpdateSegmentationArg struct {
	Id     int
	Points []domain.SegmentationPoint
}
