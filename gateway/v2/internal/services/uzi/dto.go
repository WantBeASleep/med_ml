package uzi

import (
	"github.com/google/uuid"
	ht "github.com/ogen-go/ogen/http"
)

type CreateUziArg struct {
	File       ht.MultipartFile
	Projection string
	ExternalID uuid.UUID
	DeviceID   int
}

type UpdateUziArg struct {
	Id         uuid.UUID
	Projection *string
	Checked    *bool
}
