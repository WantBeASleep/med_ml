package mappers

import (
	domain "composition-api/internal/domain/med"
	api "composition-api/internal/generated/http/api"
	"composition-api/internal/server/mappers"
)

type Doctor struct{}

func (m Doctor) Api(d domain.Doctor) api.Doctor {
	return api.Doctor{
		ID:          d.Id,
		Fullname:    d.FullName,
		Org:         d.Org,
		Job:         d.Job,
		Description: mappers.ToOptString(d.Description),
	}
}

func (m Doctor) SliceApi(d []domain.Doctor) []api.Doctor {
	return slice(d, m)
}
