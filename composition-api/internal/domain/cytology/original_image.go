package cytology

import (
	"time"

	"github.com/google/uuid"
)

type OriginalImage struct {
	Id         uuid.UUID
	CytologyID uuid.UUID
	ImagePath  string
	CreateDate time.Time
	DelayTime  *float64
	ViewedFlag bool
}
