package domain

import (
	"github.com/google/uuid"
)

type PaymentProvider struct {
	ID       uuid.UUID
	Name     string
	IsActive bool
}
