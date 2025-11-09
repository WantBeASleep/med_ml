package uzi

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	repoEntity "uzi/internal/repository/entity"
	"uzi/internal/repository/uzi/entity"
)

func (q *repo) UpdateUzi(uzi entity.Uzi) error {
	query := q.QueryBuilder().
		Update(table).
		SetMap(sq.Eq{
			columnProjection: uzi.Projection,
			columnChecked:    uzi.Checked,
			columnStatus:     uzi.Status,
		}).
		Where(sq.Eq{
			columnID: uzi.Id,
		})

	_, err := q.Runner().Execx(q.Context(), query)
	if err != nil {
		return repoEntity.WrapDBError(err)
	}

	return nil
}

func (q *repo) UpdateUziStatus(id uuid.UUID, status string) error {
	query := q.QueryBuilder().
		Update(table).
		SetMap(sq.Eq{
			columnStatus: status,
		}).
		Where(sq.Eq{
			columnID: id,
		})

	_, err := q.Runner().Execx(q.Context(), query)
	if err != nil {
		return repoEntity.WrapDBError(err)
	}

	return nil
}
