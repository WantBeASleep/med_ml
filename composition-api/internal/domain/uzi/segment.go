package domain

import (
	"encoding/json"

	"github.com/google/uuid"
)

type Segment struct {
	Id       uuid.UUID
	ImageID  uuid.UUID
	NodeID   uuid.UUID
	Contor   json.RawMessage
	Tirads23 float64
	Tirads4  float64
	Tirads5  float64
}
