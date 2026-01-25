package domain

import "github.com/google/uuid"

type Card struct {
	ID        *int
	DoctorID  uuid.UUID
	PatientID uuid.UUID
	Diagnosis *string
}
