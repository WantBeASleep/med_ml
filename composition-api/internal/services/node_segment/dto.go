package node_segment

import "github.com/google/uuid"

type CreateNodeWithSegmentArg struct {
	Node struct {
		UziID     uuid.UUID
		Ai        bool
		Tirads_23 float64
		Tirads_4  float64
		Tirads_5  float64
	}

	Segments []struct {
		ImageID   uuid.UUID
		Contor    []byte
		Tirads_23 float64
		Tirads_4  float64
		Tirads_5  float64
	}
}
