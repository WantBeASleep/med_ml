package segmentation_group

import (
	sq "github.com/Masterminds/squirrel"

	repoEntity "cytology/internal/repository/entity"
	"cytology/internal/repository/segmentation_group/entity"
)

func (q *repo) UpdateSegmentationGroup(group entity.SegmentationGroup) error {
	updateMap := sq.Eq{}

	if group.SegType != "" {
		updateMap[columnSegType] = group.SegType
	}
	if group.Details.Valid {
		updateMap[columnDetails] = group.Details
	}

	query := q.QueryBuilder().
		Update(table).
		SetMap(updateMap).
		Where(sq.Eq{
			columnID: group.Id,
		})

	_, err := q.Runner().Execx(q.Context(), query)
	if err != nil {
		return repoEntity.WrapDBError(err)
	}

	return nil
}
