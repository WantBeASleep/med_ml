package domain

import (
	"time"

	"github.com/google/uuid"
)

type Uzi struct {
	Id         uuid.UUID
	Projection string
	Checked    bool
	ExternalID uuid.UUID
	DeviceID   int
	Status     UziStatus
	CreateAt   time.Time
}
