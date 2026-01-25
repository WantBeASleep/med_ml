package segmentation_group

import (
	"database/sql"
	"errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	daoEntity "cytology/internal/repository/entity"
	"cytology/internal/repository/segmentation_group/entity"
)

func (q *repo) GetSegmentationGroupByID(id int) (entity.SegmentationGroup, error) {
	query := q.QueryBuilder().
		Select(
			columnID,
			columnCytologyID,
			columnSegType,
			columnGroupType,
			columnIsAI,
			columnDetails,
			columnCreateAt,
		).
		From(table).
		Where(sq.Eq{
			columnID: id,
		})

	var group entity.SegmentationGroup
	if err := q.Runner().Getx(q.Context(), &group, query); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.SegmentationGroup{}, daoEntity.ErrNotFound
		}
		return entity.SegmentationGroup{}, err
	}

	return group, nil
}

func (q *repo) GetSegmentationGroupsByCytologyID(cytologyID uuid.UUID) ([]entity.SegmentationGroup, error) {
	query := q.QueryBuilder().
		Select(
			columnID,
			columnCytologyID,
			columnSegType,
			columnGroupType,
			columnIsAI,
			columnDetails,
			columnCreateAt,
		).
		From(table).
		Where(sq.Eq{
			columnCytologyID: cytologyID,
		})

	var groups []entity.SegmentationGroup
	if err := q.Runner().Selectx(q.Context(), &groups, query); err != nil {
		return nil, err
	}

	if len(groups) == 0 {
		return nil, daoEntity.ErrNotFound
	}

	return groups, nil
}
