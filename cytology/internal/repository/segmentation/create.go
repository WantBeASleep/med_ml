package segmentation

import (
	repoEntity "cytology/internal/repository/entity"
	"cytology/internal/repository/segmentation/entity"
)

func (q *repo) InsertSegmentation(seg entity.Segmentation) (int, error) {
	// Вставляем сегментацию (ID будет сгенерирован автоматически, так как это SERIAL)
	query := q.QueryBuilder().
		Insert(table).
		Columns(
			columnSegmentationGroupID,
			columnCreateAt,
		).
		Values(
			seg.SegmentationGroupID,
			seg.CreateAt,
		).
		Suffix("RETURNING id")

	var id int
	err := q.Runner().Getx(q.Context(), &id, query)
	if err != nil {
		return 0, repoEntity.WrapDBError(err)
	}

	// Вставляем точки (ID будут сгенерированы автоматически)
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
				id, // Используем созданный ID сегментации
				point.X,
				point.Y,
				point.UID,
				point.CreateAt,
			)
		}

		_, err = q.Runner().Execx(q.Context(), pointsQuery)
		if err != nil {
			return 0, repoEntity.WrapDBError(err)
		}
	}

	return id, nil
}
