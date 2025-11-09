package segment

import (
	repoEntity "uzi/internal/repository/entity"
	"uzi/internal/repository/segment/entity"

	sq "github.com/Masterminds/squirrel"
)

func (q *repo) UpdateSegment(segment entity.Segment) error {
	query := q.QueryBuilder().
		Update(table).
		SetMap(sq.Eq{
			columnContor:   segment.Contor,
			columnTirads23: segment.Tirads23,
			columnTirads4:  segment.Tirads4,
			columnTirads5:  segment.Tirads5,
		}).
		Where(sq.Eq{
			columnID: segment.Id,
		})

	_, err := q.Runner().Execx(q.Context(), query)
	if err != nil {
		return repoEntity.WrapDBError(err)
	}

	return nil
}
