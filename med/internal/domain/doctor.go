package domain

import "github.com/google/uuid"

type Doctor struct {
	Id          uuid.UUID
	FullName    string
	Org         string
	Job         string
	Description *string
}
