package cytology_image

import (
	"github.com/google/uuid"

	"cytology/internal/domain"
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
	Details           []byte
	PrevID            *uuid.UUID
	ParentPrevID      *uuid.UUID
	File              []byte
	ContentType       string
}

type UpdateCytologyImageArg struct {
	Id                uuid.UUID
	DiagnosticMarking *domain.DiagnosticMarking
	MaterialType      *domain.MaterialType
	Calcitonin        *int
	CalcitoninInFlush *int
	Thyroglobulin     *int
	Details           []byte
	IsLast            *bool
	PrevID            *uuid.UUID
	ParentPrevID      *uuid.UUID
}

func (u UpdateCytologyImageArg) UpdateDomain(d *domain.CytologyImage) {
	if u.DiagnosticMarking != nil {
		d.DiagnosticMarking = u.DiagnosticMarking
	}
	if u.MaterialType != nil {
		d.MaterialType = u.MaterialType
	}
	if u.Calcitonin != nil {
		d.Calcitonin = u.Calcitonin
	}
	if u.CalcitoninInFlush != nil {
		d.CalcitoninInFlush = u.CalcitoninInFlush
	}
	if u.Thyroglobulin != nil {
		d.Thyroglobulin = u.Thyroglobulin
	}
	if u.Details != nil {
		d.Details = u.Details
	}
	if u.IsLast != nil {
		d.IsLast = *u.IsLast
	}
	if u.PrevID != nil {
		d.PrevID = u.PrevID
	}
	if u.ParentPrevID != nil {
		d.ParentPrevID = u.ParentPrevID
	}
}
