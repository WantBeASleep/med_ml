package mappers

import (
	domain "composition-api/internal/domain/med"
	api "composition-api/internal/generated/http/api"
	"composition-api/internal/server/mappers"
)

type Patient struct{}

func (m Patient) Api(d domain.Patient) api.Patient {
	return api.Patient{
		ID:          d.Id,
		Fullname:    d.FullName,
		Email:       d.Email,
		Policy:      d.Policy,
		Active:      d.Active,
		Malignancy:  d.Malignancy,
		BirthDate:   d.BirthDate,
		LastUziDate: mappers.ToOptDate(d.LastUziDate),
	}
}

func (m Patient) SliceApi(d []domain.Patient) []api.Patient {
	return slice(d, m)
}
