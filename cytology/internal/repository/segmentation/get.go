package segmentation

import (
	"database/sql"
	"errors"

	sq "github.com/Masterminds/squirrel"

	daoEntity "cytology/internal/repository/entity"
	"cytology/internal/repository/segmentation/entity"
)

func (q *repo) GetSegmentationByID(id int) (entity.Segmentation, error) {
	// Получаем сегментацию
	query := q.QueryBuilder().
		Select(
			columnID,
			columnSegmentationGroupID,
			columnCreateAt,
		).
		From(table).
		Where(sq.Eq{
			columnID: id,
		})

	var seg entity.Segmentation
	if err := q.Runner().Getx(q.Context(), &seg, query); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.Segmentation{}, daoEntity.ErrNotFound
		}
		return entity.Segmentation{}, err
	}

	// Получаем точки
	pointsQuery := q.QueryBuilder().
		Select(
			pointColumnID,
			pointColumnSegmentationID,
			pointColumnX,
			pointColumnY,
			pointColumnUID,
			pointColumnCreateAt,
		).
		From(pointTable).
		Where(sq.Eq{
			pointColumnSegmentationID: id,
		})

	var points []entity.SegmentationPoint
	if err := q.Runner().Selectx(q.Context(), &points, pointsQuery); err != nil {
		return entity.Segmentation{}, err
	}

	seg.Points = points
	return seg, nil
}

func (q *repo) GetSegmentsByGroupID(groupID int) ([]entity.Segmentation, error) {
	// Получаем все сегментации группы
	query := q.QueryBuilder().
		Select(
			columnID,
			columnSegmentationGroupID,
			columnCreateAt,
		).
		From(table).
		Where(sq.Eq{
			columnSegmentationGroupID: groupID,
		})

	var segs []entity.Segmentation
	if err := q.Runner().Selectx(q.Context(), &segs, query); err != nil {
		return nil, err
	}

	if len(segs) == 0 {
		return nil, daoEntity.ErrNotFound
	}

	// Получаем точки для каждой сегментации
	for i := range segs {
		pointsQuery := q.QueryBuilder().
			Select(
				pointColumnID,
				pointColumnSegmentationID,
				pointColumnX,
				pointColumnY,
				pointColumnUID,
				pointColumnCreateAt,
			).
			From(pointTable).
			Where(sq.Eq{
				pointColumnSegmentationID: segs[i].Id,
			})

		var points []entity.SegmentationPoint
		if err := q.Runner().Selectx(q.Context(), &points, pointsQuery); err != nil {
			return nil, err
		}

		segs[i].Points = points
	}

	return segs, nil
}
