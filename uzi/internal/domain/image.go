package domain

import (
	"github.com/google/uuid"
)

type Image struct {
	Id    uuid.UUID
	UziID uuid.UUID
	Page  int
}	