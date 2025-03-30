package doctor

import (
	"database/sql"
	"errors"

	dentity "med/internal/repository/doctor/entity"
	"med/internal/repository/entity"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

func (r *repo) GetDoctorByID(id uuid.UUID) (dentity.Doctor, error) {
	query := r.QueryBuilder().
		Select(
			columnID,
			columnFullname,
			columnOrg,
			columnJob,
			columnDescription,
		).
		From(table).
		Where(sq.Eq{
			columnID: id,
		})

	var doctor dentity.Doctor
	if err := r.Runner().Getx(r.Context(), &doctor, query); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return dentity.Doctor{}, entity.ErrNotFound
		}
		return dentity.Doctor{}, err
	}

	return doctor, nil
}
