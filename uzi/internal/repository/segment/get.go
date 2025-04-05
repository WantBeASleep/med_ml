package segment

import (
	"database/sql"
	"errors"

	daoEntity "uzi/internal/repository/entity"
	"uzi/internal/repository/segment/entity"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

func (q *repo) GetSegmentByID(id uuid.UUID) (entity.Segment, error) {
	query := q.QueryBuilder().
		Select(
			columnID,
			columnNodeID,
			columnImageID,
			columnContor,
			columnAi,
			columnTirads23,
			columnTirads4,
			columnTirads5,
		).
		From(table).
		Where(sq.Eq{
			columnID: id,
		})

	var segment entity.Segment
	if err := q.Runner().Getx(q.Context(), &segment, query); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.Segment{}, daoEntity.ErrNotFound
		}
		return entity.Segment{}, err
	}

	return segment, nil
}

func (q *repo) GetSegmentsByNodeID(id uuid.UUID) ([]entity.Segment, error) {
	query := q.QueryBuilder().
		Select(
			columnID,
			columnNodeID,
			columnImageID,
			columnContor,
			columnAi,
			columnTirads23,
			columnTirads4,
			columnTirads5,
		).
		From(table).
		Where(sq.Eq{
			columnNodeID: id,
		})

	var segments []entity.Segment
	if err := q.Runner().Selectx(q.Context(), &segments, query); err != nil {
		return nil, err
	}

	if len(segments) == 0 {
		return nil, daoEntity.ErrNotFound
	}

	return segments, nil
}

func (q *repo) GetSegmentsByImageID(id uuid.UUID) ([]entity.Segment, error) {
	query := q.QueryBuilder().
		Select(
			columnID,
			columnNodeID,
			columnImageID,
			columnContor,
			columnAi,
			columnTirads23,
			columnTirads4,
			columnTirads5,
		).
		From(table).
		Where(sq.Eq{
			columnImageID: id,
		})

	var segments []entity.Segment
	if err := q.Runner().Selectx(q.Context(), &segments, query); err != nil {
		return nil, err
	}

	if len(segments) == 0 {
		return nil, daoEntity.ErrNotFound
	}

	return segments, nil
}
