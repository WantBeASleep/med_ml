package cytology

import (
	"github.com/google/uuid"
	ht "github.com/ogen-go/ogen/http"

	domain "composition-api/internal/domain/cytology"
)

type CreateCytologyImageIn struct {
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

type UpdateCytologyImageIn struct {
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

type CreateOriginalImageIn struct {
	CytologyID  uuid.UUID
	File        ht.MultipartFile // Используется только если ImagePath не указан
	ContentType string
	DelayTime   *float64
	ImagePath   *string // Путь к файлу в S3 (если файл уже загружен)
}

type UpdateOriginalImageIn struct {
	Id         uuid.UUID
	DelayTime  *float64
	ViewedFlag *bool
}

type CreateSegmentationGroupIn struct {
	CytologyID uuid.UUID
	SegType    domain.SegType
	GroupType  domain.GroupType
	IsAI       bool
	Details    *string
}

type UpdateSegmentationGroupIn struct {
	Id      int
	SegType *domain.SegType
	Details *string
}

type CreateSegmentationIn struct {
	SegmentationGroupID int
	Points              []domain.SegmentationPoint
}

type UpdateSegmentationIn struct {
	Id     int
	Points []domain.SegmentationPoint
}
