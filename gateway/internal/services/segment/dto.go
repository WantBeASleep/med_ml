package segment

import "github.com/google/uuid"

type CreateSegmentArg struct {
	ImageID   uuid.UUID
	NodeID    uuid.UUID
	Contor    []byte
	Tirads_23 float64
	Tirads_4  float64
	Tirads_5  float64
}

type UpdateSegmentArg struct {
	Id        uuid.UUID
	Tirads_23 *float64
	Tirads_4  *float64
	Tirads_5  *float64
}
