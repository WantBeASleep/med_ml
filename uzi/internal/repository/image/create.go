package image

import (
	"uzi/internal/repository/image/entity"
)

func (q *repo) InsertImages(images ...entity.Image) error {
	query := q.QueryBuilder().
		Insert(table).
		Columns(
			columnId,
			columnUziId,
			columnPage,
		)

	for _, v := range images {
		query = query.Values(
			v.Id,
			v.UziID,
			v.Page,
		)
	}

	_, err := q.Runner().Execx(q.Context(), query)
	if err != nil {
		return err
	}

	return nil
}
