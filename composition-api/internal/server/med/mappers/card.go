package mappers

import (
	domain "composition-api/internal/domain/med"
	api "composition-api/internal/generated/http/api"
	mappers "composition-api/internal/server/mappers"
)

type Card struct{}

func (m Card) Api(d domain.Card) api.Card {
	return api.Card{
		DoctorID:  d.DoctorID,
		PatientID: d.PatientID,
		Diagnosis: mappers.ToOptString(d.Diagnosis),
	}
}

func (m Card) SliceApi(d []domain.Card) []api.Card {
	return slice(d, m)
}
