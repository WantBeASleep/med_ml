package entity

import (
	"fmt"

	"github.com/google/uuid"
)

type Doctor struct {
	ID          uuid.UUID
	Fullname    string
	Org         string
	Job         string
	Description string
}

func (d Doctor) Validate() error {
	if d.ID == uuid.Nil {
		return ErrEmptyDoctorID
	}

	if d.Fullname == "" {
		return ErrEmptyDoctorFullname
	}

	return nil
}

func NewDoctor(id, fullname, org, job, description string) (Doctor, error) {
	doctorUUID, err := uuid.Parse(id)
	if err != nil {
		return Doctor{}, fmt.Errorf("parse uuid: %w", err)
	}

	return Doctor{
		ID:          doctorUUID,
		Fullname:    fullname,
		Org:         org,
		Job:         job,
		Description: description,
	}, nil
}
