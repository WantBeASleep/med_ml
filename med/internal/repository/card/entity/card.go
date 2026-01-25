package entity

import (
	"database/sql"

	"med/internal/domain"

	"github.com/google/uuid"
)

type Card struct {
	ID        sql.NullInt32  `db:"id"`
	DoctorID  uuid.UUID      `db:"doctor_id"`
	PatientID uuid.UUID      `db:"patient_id"`
	Diagnosis sql.NullString `db:"diagnosis"`
}

func (Card) FromDomain(p domain.Card) Card {
	card := Card{
		DoctorID:  p.DoctorID,
		PatientID: p.PatientID,
	}

	if p.ID != nil {
		card.ID = sql.NullInt32{Int32: int32(*p.ID), Valid: true}
	}

	if p.Diagnosis != nil {
		card.Diagnosis = sql.NullString{String: *p.Diagnosis, Valid: true}
	}

	return card
}

func (p Card) ToDomain() domain.Card {
	card := domain.Card{
		DoctorID:  p.DoctorID,
		PatientID: p.PatientID,
	}

	if p.ID.Valid {
		id := int(p.ID.Int32)
		card.ID = &id
	}

	if p.Diagnosis.Valid {
		card.Diagnosis = &p.Diagnosis.String
	}

	return card
}
