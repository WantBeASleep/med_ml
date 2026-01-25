package segmentation

import (
	sq "github.com/Masterminds/squirrel"

	repoEntity "cytology/internal/repository/entity"
	"cytology/internal/repository/segmentation/entity"
)

func (q *repo) UpdateSegmentation(seg entity.Segmentation) error {
	// Удаляем старые точки
	deleteQuery := q.QueryBuilder().
		Delete(pointTable).
		Where(sq.Eq{
			pointColumnSegmentationID: seg.Id,
		})

	_, err := q.Runner().Execx(q.Context(), deleteQuery)
	if err != nil {
		return repoEntity.WrapDBError(err)
	}

	// Вставляем новые точки (ID будут сгенерированы автоматически БД)
	if len(seg.Points) > 0 {
		pointsQuery := q.QueryBuilder().
			Insert(pointTable).
			Columns(
				pointColumnSegmentationID,
				pointColumnX,
				pointColumnY,
				pointColumnUID,
				pointColumnCreateAt,
			)

		for _, point := range seg.Points {
			pointsQuery = pointsQuery.Values(
				point.SegmentationID,
				point.X,
				point.Y,
				point.UID,
				point.CreateAt,
			)
		}

		_, err = q.Runner().Execx(q.Context(), pointsQuery)
		if err != nil {
			return repoEntity.WrapDBError(err)
		}
	}

	return nil
}
