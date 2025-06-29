package entity

import (
	"github.com/google/uuid"
)

type RefreshToken struct {
	Id           uuid.UUID `db:"id"`
	RefreshToken string    `db:"refresh_token"`
}

// конвентера нет, состовная сущность
