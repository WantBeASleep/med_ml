package domain

import "github.com/google/uuid"

type Card struct {
	DoctorID  uuid.UUID
	PatientID uuid.UUID
	Diagnosis *string
}
