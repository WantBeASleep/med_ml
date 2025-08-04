package redis

import (
	"time"
)

type DoctorDTO struct {
	ID          string `redis:"id"`
	Fullname    string `redis:"name"`
	Org         string `redis:"org"`
	Job         string `redis:"job"`
	Description string `redis:"desc"`
}

type PatientDTO struct {
	ID         string    `redis:"id"`
	Fullname   string    `redis:"name"`
	Email      string    `redis:"email"`
	Policy     string    `redis:"policy"`
	Active     bool      `redis:"active"`
	Malignancy bool      `redis:"malignancy"`
	BirthDate  time.Time `redis:"birth_date"`
	LastUSG    time.Time `redis:"last_usg"`
}
