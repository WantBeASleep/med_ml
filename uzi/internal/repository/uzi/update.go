package uzi

import (
	"uzi/internal/repository/uzi/entity"

	sq "github.com/Masterminds/squirrel"
)

func (q *repo) UpdateUzi(uzi entity.Uzi) error {
	query := q.QueryBuilder().
		Update(uziTable).
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
		return err
	}

	return nil
}
