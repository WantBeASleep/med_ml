package mappers

import (
	domain "composition-api/internal/domain/uzi"
	api "composition-api/internal/generated/http/api"
)

func Uzi(uzi domain.Uzi) api.Uzi {
	return api.Uzi{
		ID:         uzi.Id,
		Projection: uzi.Projection,
		Checked:    uzi.Checked,
		ExternalID: uzi.ExternalID,
		DeviceID:   uzi.DeviceID,
		Status:     api.UziStatus(uzi.Status),
		CreateAt:   uzi.CreateAt,
	}
}

func SliceUzi(uzis []domain.Uzi) []api.Uzi {
	result := make([]api.Uzi, 0, len(uzis))
	for _, uzi := range uzis {
		result = append(result, Uzi(uzi))
	}
	return result
}
