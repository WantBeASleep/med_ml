package entity

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Patient struct {
	ID          uuid.UUID
	Fullname    string
	Email       string
	Policy      string
	Active      bool
	Malignancy  bool
	BirthDate   time.Time
	LastUziDate time.Time
}

func (p Patient) Validate() error {
	if p.ID == uuid.Nil {
		return ErrEmptyPatientID
	}

	if p.Fullname == "" {
		return ErrEmptyPatientFullname
	}

	if p.Email == "" {
		return ErrEmptyPatientEmail
	}

	if p.Policy == "" {
		return ErrEmptyPatientPolicy
	}

	return nil
}

func NewPatient(id, fullname, email, policy string, active, malignancy bool, birthDate, lastUziDate time.Time) (Patient, error) {
	patientUUID, err := uuid.Parse(id)
	if err != nil {
		return Patient{}, fmt.Errorf("parse uuid: %w", err)
	}

	return Patient{
		ID:          patientUUID,
		Fullname:    fullname,
		Email:       email,
		Policy:      policy,
		Active:      active,
		Malignancy:  malignancy,
		BirthDate:   birthDate,
		LastUziDate: lastUziDate,
	}, nil
}
