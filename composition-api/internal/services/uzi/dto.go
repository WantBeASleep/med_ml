package uzi

import (
	domain "composition-api/internal/domain/uzi"

	"github.com/google/uuid"
	ht "github.com/ogen-go/ogen/http"
)

type CreateUziArg struct {
	File        ht.MultipartFile
	Projection  domain.UziProjection
	ExternalID  uuid.UUID
	Author      uuid.UUID
	DeviceID    int
	Description *string
}

type UpdateUziArg struct {
	Id         uuid.UUID
	Projection *domain.UziProjection
	Checked    *bool
}
