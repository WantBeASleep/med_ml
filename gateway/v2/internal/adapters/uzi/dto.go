package uzi

import "github.com/google/uuid"

type CreateUziIn struct {
	Projection string
	ExternalID string
	DeviceID   int
}

type UpdateUziIn struct {
	Id         uuid.UUID
	Projection *string
	Checked    *bool
}

type UpdateNodeIn struct {
	Id        uuid.UUID
	Tirads_23 *float64
	Tirads_4  *float64
	Tirads_5  *float64
}

type CreateSegmentIn struct {
	ImageID   string
	NodeID    string
	Contor    []byte
	Tirads_23 float64
	Tirads_4  float64
	Tirads_5  float64
}

type UpdateSegmentIn struct {
	Id        uuid.UUID
	Tirads_23 *float64
	Tirads_4  *float64
	Tirads_5  *float64
}

type CreateNodeWithSegmentsIn struct {
	Node struct {
		UziID     string
		Ai        bool
		Tirads_23 float64
		Tirads_4  float64
		Tirads_5  float64
	}

	Segments []struct {
		ImageID   string
		Contor    []byte
		Tirads_23 float64
		Tirads_4  float64
		Tirads_5  float64
	}
}
