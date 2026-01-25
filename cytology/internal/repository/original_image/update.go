package original_image

import (
	sq "github.com/Masterminds/squirrel"

	repoEntity "cytology/internal/repository/entity"
	"cytology/internal/repository/original_image/entity"
)

func (q *repo) UpdateOriginalImage(img entity.OriginalImage) error {
	updateMap := sq.Eq{
		columnViewedFlag: img.ViewedFlag,
	}

	if img.DelayTime.Valid {
		updateMap[columnDelayTime] = img.DelayTime
	}

	query := q.QueryBuilder().
		Update(table).
		SetMap(updateMap).
		Where(sq.Eq{
			columnID: img.Id,
		})

	_, err := q.Runner().Execx(q.Context(), query)
	if err != nil {
		return repoEntity.WrapDBError(err)
	}

	return nil
}
