package segment

import (
	"encoding/json"

	"github.com/google/uuid"
)

type CreateSegmentArg struct {
	ImageID   uuid.UUID
	NodeID    uuid.UUID
	Contor    json.RawMessage
	Tirads_23 float64
	Tirads_4  float64
	Tirads_5  float64
}

type UpdateSegmentArg struct {
	Id        uuid.UUID
	Contor    json.RawMessage
	Tirads_23 *float64
	Tirads_4  *float64
	Tirads_5  *float64
}
