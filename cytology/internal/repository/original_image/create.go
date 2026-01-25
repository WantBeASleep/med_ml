package original_image

import (
	repoEntity "cytology/internal/repository/entity"
	"cytology/internal/repository/original_image/entity"
)

func (q *repo) InsertOriginalImage(img entity.OriginalImage) error {
	query := q.QueryBuilder().
		Insert(table).
		Columns(
			columnID,
			columnCytologyID,
			columnImagePath,
			columnCreateDate,
			columnDelayTime,
			columnViewedFlag,
		).
		Values(
			img.Id,
			img.CytologyID,
			img.ImagePath,
			img.CreateDate,
			img.DelayTime,
			img.ViewedFlag,
		)

	_, err := q.Runner().Execx(q.Context(), query)
	if err != nil {
		return repoEntity.WrapDBError(err)
	}

	return nil
}
