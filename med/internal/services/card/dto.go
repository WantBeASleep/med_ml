package card

import (
	"med/internal/domain"
)

type UpdateCardArg struct {
	Diagnosis *string
}

func (u UpdateCardArg) Update(d *domain.Card) {
	if u.Diagnosis != nil {
		d.Diagnosis = u.Diagnosis
	}
}
