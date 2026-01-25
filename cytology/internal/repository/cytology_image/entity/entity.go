package entity

import (
	"database/sql"
	"encoding/json"
	"time"

	"cytology/internal/domain"

	"github.com/google/uuid"
)

type CytologyImage struct {
	Id                uuid.UUID      `db:"id"`
	ExternalID        uuid.UUID      `db:"external_id"`
	DoctorID          uuid.UUID      `db:"doctor_id"`
	PatientID         uuid.UUID      `db:"patient_id"`
	DiagnosticNumber  int            `db:"diagnostic_number"`
	DiagnosticMarking sql.NullString `db:"diagnostic_marking"`
	MaterialType      sql.NullString `db:"material_type"`
	DiagnosDate       time.Time      `db:"diagnos_date"`
	IsLast            bool           `db:"is_last"`
	Calcitonin        sql.NullInt32  `db:"calcitonin"`
	CalcitoninInFlush sql.NullInt32  `db:"calcitonin_in_flush"`
	Thyroglobulin     sql.NullInt32  `db:"thyroglobulin"`
	Details           sql.NullString `db:"details"`
	PrevID            uuid.NullUUID  `db:"prev_id"`
	ParentPrevID      uuid.NullUUID  `db:"parent_prev_id"`
	CreateAt          time.Time      `db:"create_at"`
}

func (CytologyImage) FromDomain(d domain.CytologyImage) CytologyImage {
	var diagnosticMarking sql.NullString
	if d.DiagnosticMarking != nil {
		diagnosticMarking = sql.NullString{String: d.DiagnosticMarking.String(), Valid: true}
	}

	var materialType sql.NullString
	if d.MaterialType != nil {
		materialType = sql.NullString{String: d.MaterialType.String(), Valid: true}
	}

	var calcitonin sql.NullInt32
	if d.Calcitonin != nil {
		calcitonin = sql.NullInt32{Int32: int32(*d.Calcitonin), Valid: true}
	}

	var calcitoninInFlush sql.NullInt32
	if d.CalcitoninInFlush != nil {
		calcitoninInFlush = sql.NullInt32{Int32: int32(*d.CalcitoninInFlush), Valid: true}
	}

	var thyroglobulin sql.NullInt32
	if d.Thyroglobulin != nil {
		thyroglobulin = sql.NullInt32{Int32: int32(*d.Thyroglobulin), Valid: true}
	}

	var prevID uuid.NullUUID
	if d.PrevID != nil {
		prevID = uuid.NullUUID{UUID: *d.PrevID, Valid: true}
	}

	var parentPrevID uuid.NullUUID
	if d.ParentPrevID != nil {
		parentPrevID = uuid.NullUUID{UUID: *d.ParentPrevID, Valid: true}
	}

	// Конвертируем Details в sql.NullString для правильной обработки NULL
	var details sql.NullString
	if len(d.Details) > 0 {
		details = sql.NullString{String: string(d.Details), Valid: true}
	}

	return CytologyImage{
		Id:                d.Id,
		ExternalID:        d.ExternalID,
		DoctorID:          d.DoctorID,
		PatientID:         d.PatientID,
		DiagnosticNumber:  d.DiagnosticNumber,
		DiagnosticMarking: diagnosticMarking,
		MaterialType:      materialType,
		DiagnosDate:       d.DiagnosDate,
		IsLast:            d.IsLast,
		Calcitonin:        calcitonin,
		CalcitoninInFlush: calcitoninInFlush,
		Thyroglobulin:     thyroglobulin,
		Details:           details,
		PrevID:            prevID,
		ParentPrevID:      parentPrevID,
		CreateAt:          d.CreateAt,
	}
}

func (d CytologyImage) ToDomain() domain.CytologyImage {
	var diagnosticMarking *domain.DiagnosticMarking
	if d.DiagnosticMarking.Valid {
		dm := domain.DiagnosticMarking(d.DiagnosticMarking.String)
		diagnosticMarking = &dm
	}

	var materialType *domain.MaterialType
	if d.MaterialType.Valid {
		mt := domain.MaterialType(d.MaterialType.String)
		materialType = &mt
	}

	var calcitonin *int
	if d.Calcitonin.Valid {
		val := int(d.Calcitonin.Int32)
		calcitonin = &val
	}

	var calcitoninInFlush *int
	if d.CalcitoninInFlush.Valid {
		val := int(d.CalcitoninInFlush.Int32)
		calcitoninInFlush = &val
	}

	var thyroglobulin *int
	if d.Thyroglobulin.Valid {
		val := int(d.Thyroglobulin.Int32)
		thyroglobulin = &val
	}

	var prevID *uuid.UUID
	if d.PrevID.Valid {
		prevID = &d.PrevID.UUID
	}

	var parentPrevID *uuid.UUID
	if d.ParentPrevID.Valid {
		parentPrevID = &d.ParentPrevID.UUID
	}

	// Конвертируем Details обратно в json.RawMessage
	var details json.RawMessage
	if d.Details.Valid {
		details = json.RawMessage(d.Details.String)
	}

	return domain.CytologyImage{
		Id:                d.Id,
		ExternalID:        d.ExternalID,
		DoctorID:          d.DoctorID,
		PatientID:         d.PatientID,
		DiagnosticNumber:  d.DiagnosticNumber,
		DiagnosticMarking: diagnosticMarking,
		MaterialType:      materialType,
		DiagnosDate:       d.DiagnosDate,
		IsLast:            d.IsLast,
		Calcitonin:        calcitonin,
		CalcitoninInFlush: calcitoninInFlush,
		Thyroglobulin:     thyroglobulin,
		Details:           details,
		PrevID:            prevID,
		ParentPrevID:      parentPrevID,
		CreateAt:          d.CreateAt,
	}
}

func (CytologyImage) SliceToDomain(images []CytologyImage) []domain.CytologyImage {
	domainImages := make([]domain.CytologyImage, 0, len(images))
	for _, v := range images {
		domainImages = append(domainImages, v.ToDomain())
	}
	return domainImages
}
