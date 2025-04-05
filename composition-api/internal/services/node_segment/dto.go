package node_segment

import (
	"encoding/json"

	"github.com/google/uuid"
)

type CreateNodeWithSegmentArg_Node struct {
	Tirads_23   float64
	Tirads_4    float64
	Tirads_5    float64
	Description *string
}

type CreateNodeWithSegmentArg_Segment struct {
	ImageID   uuid.UUID
	Contor    json.RawMessage
	Tirads_23 float64
	Tirads_4  float64
	Tirads_5  float64
}

type CreateNodeWithSegmentArg struct {
	UziID    uuid.UUID
	Node     CreateNodeWithSegmentArg_Node
	Segments []CreateNodeWithSegmentArg_Segment
}
