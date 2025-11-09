package doctor

import (
	"med/internal/repository/doctor/entity"
	repoEntity "med/internal/repository/entity"

	sq "github.com/Masterminds/squirrel"
)

func (r *repo) UpdateDoctor(doctor entity.Doctor) error {
	query := r.QueryBuilder().
		Update(table).
		SetMap(sq.Eq{
			columnOrg:         doctor.Org,
			columnJob:         doctor.Job,
			columnDescription: doctor.Description,
		}).
		Where(sq.Eq{
			columnID: doctor.Id,
		})

	_, err := r.Runner().Execx(r.Context(), query)
	if err != nil {
		return repoEntity.WrapDBError(err)
	}
	return nil
}
