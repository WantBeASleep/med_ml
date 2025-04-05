package node

import (
	"github.com/google/uuid"

	"uzi/internal/domain"
)

// TODO: починить баг при запросе со всеми полями nil
type UpdateNodeArg struct {
	Id         uuid.UUID
	Validation *domain.NodeValidation
	Tirads23   *float64
	Tirads4    *float64
	Tirads5    *float64
}

func (u UpdateNodeArg) UpdateDomain(d *domain.Node) {
	if u.Validation != nil {
		d.Validation = u.Validation
	}
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
