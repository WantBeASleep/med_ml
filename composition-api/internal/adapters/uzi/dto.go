package uzi

import (
	"encoding/json"

	"github.com/google/uuid"

	domain "composition-api/internal/domain/uzi"
)

type CreateUziIn struct {
	Projection  domain.UziProjection
	ExternalID  uuid.UUID
	Author      uuid.UUID
	DeviceID    int
	Description *string
}

type UpdateUziIn struct {
	Id         uuid.UUID
	Projection *domain.UziProjection
	Checked    *bool
}

type UpdateNodeIn struct {
	Id         uuid.UUID
	Validation *domain.NodeValidation
	Tirads_23  *float64
	Tirads_4   *float64
	Tirads_5   *float64
}

type CreateSegmentIn struct {
	ImageID   uuid.UUID
	NodeID    uuid.UUID
	Contor    json.RawMessage
	Tirads_23 float64
	Tirads_4  float64
	Tirads_5  float64
}

type UpdateSegmentIn struct {
	Id        uuid.UUID
	Contor    json.RawMessage
	Tirads_23 *float64
	Tirads_4  *float64
	Tirads_5  *float64
}

type CreateNodeWithSegmentsIn_Node struct {
	Tirads_23   float64
	Tirads_4    float64
	Tirads_5    float64
	Description *string
}

type CreateNodeWithSegmentsIn_Segment struct {
	ImageID   uuid.UUID
	Contor    json.RawMessage
	Tirads_23 float64
	Tirads_4  float64
	Tirads_5  float64
}

type CreateNodeWithSegmentsIn struct {
	UziID    uuid.UUID
	Node     CreateNodeWithSegmentsIn_Node
	Segments []CreateNodeWithSegmentsIn_Segment
}
