package mappers

import (
	domain "composition-api/internal/domain/uzi"
	api "composition-api/internal/generated/http/api"
)

type Uzi struct{}

func (Uzi) Domain(uzi domain.Uzi) api.Uzi {
	return api.Uzi{
		ID:         uzi.Id,
		Projection: api.UziProjection(uzi.Projection),
		Checked:    uzi.Checked,
		ExternalID: uzi.ExternalID,
		AuthorID:   uzi.Author,
		DeviceID:   uzi.DeviceID,
		Status:     api.UziStatus(uzi.Status),
		CreateAt:   uzi.CreateAt,
	}
}

func (Uzi) SliceDomain(uzis []domain.Uzi) []api.Uzi {
	return slice(uzis, Uzi{})
}
