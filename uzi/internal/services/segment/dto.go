package segment

import (
	"encoding/json"

	"github.com/google/uuid"

	"uzi/internal/domain"
)

type CreateSegmentArg struct {
	ImageID  uuid.UUID
	NodeID   uuid.UUID
	Contor   json.RawMessage
	Tirads23 float64
	Tirads4  float64
	Tirads5  float64
}

// TODO: починить баг при запросе со всеми полями nil
type UpdateSegmentArg struct {
	Id       uuid.UUID
	Tirads23 *float64
	Tirads4  *float64
	Tirads5  *float64
}

func (u UpdateSegmentArg) UpdateDomain(d *domain.Segment) {
	if u.Tirads23 != nil {
		d.Tirads23 = *u.Tirads23
	}
	if u.Tirads4 != nil {
		d.Tirads4 = *u.Tirads4
	}
	if u.Tirads5 != nil {
		d.Tirads5 = *u.Tirads5
	}
}
