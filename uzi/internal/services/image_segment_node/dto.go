package image_segment_node

import (
	"encoding/json"

	"github.com/google/uuid"
)

type CreateNodesWithSegmentsArgNode struct {
	Ai       bool
	UziID    uuid.UUID
	Tirads23 float64
	Tirads4  float64
	Tirads5  float64
}

type CreateNodesWithSegmentsArgSegment struct {
	ImageID  uuid.UUID
	Contor   json.RawMessage
	Tirads23 float64
	Tirads4  float64
	Tirads5  float64
}

type CreateNodesWithSegmentsArg struct {
	Nodes    CreateNodesWithSegmentsArgNode
	Segments []CreateNodesWithSegmentsArgSegment
}

type CreateNodesWithSegmentsID struct {
	NodeID   uuid.UUID
	SegmentID []uuid.UUID
}
