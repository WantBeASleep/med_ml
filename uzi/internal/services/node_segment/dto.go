package node_segment

import (
	"encoding/json"

	"github.com/google/uuid"
)

type CreateNodesWithSegmentsArgNode struct {
	Tirads23    float64
	Tirads4     float64
	Tirads5     float64
	Description *string
}

type CreateNodesWithSegmentsArgSegment struct {
	ImageID  uuid.UUID
	Contor   json.RawMessage
	Tirads23 float64
	Tirads4  float64
	Tirads5  float64
}

type CreateNodesWithSegmentsArg struct {
	Node     CreateNodesWithSegmentsArgNode
	Segments []CreateNodesWithSegmentsArgSegment
}

type CreateNodesWithSegmentsID struct {
	NodeID     uuid.UUID
	SegmentsID []uuid.UUID
}
