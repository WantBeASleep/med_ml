package domain

import (
	"time"

	"github.com/google/uuid"
)

type Uzi struct {
	Id          uuid.UUID
	Projection  UziProjection
	Checked     bool
	ExternalID  uuid.UUID
	Author      uuid.UUID
	DeviceID    int
	Status      UziStatus
	Description *string
	CreateAt    time.Time
}
