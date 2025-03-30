package domain

import (
	"time"

	"github.com/google/uuid"
)

type Patient struct {
	Id          uuid.UUID
	FullName    string
	Email       string
	Policy      string
	Active      bool
	Malignancy  bool
	BirthDate   time.Time
	LastUziDate *time.Time
}
