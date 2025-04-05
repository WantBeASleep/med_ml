package med

import (
	"time"

	"github.com/google/uuid"
)

type CreatePatientArg struct {
	Id         uuid.UUID
	FullName   string
	Email      string
	Policy     string
	Active     bool
	Malignancy bool
	BirthDate  time.Time
}

type UpdatePatientIn struct {
	Id          uuid.UUID
	Active      *bool
	Malignancy  *bool
	LastUziDate *time.Time
}
