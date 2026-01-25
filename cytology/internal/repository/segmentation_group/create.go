package segmentation_group

import (
	repoEntity "cytology/internal/repository/entity"
	"cytology/internal/repository/segmentation_group/entity"
)

func (q *repo) InsertSegmentationGroup(group entity.SegmentationGroup) (int, error) {
	query := q.QueryBuilder().
		Insert(table).
		Columns(
			columnCytologyID,
			columnSegType,
			columnGroupType,
			columnIsAI,
			columnDetails,
			columnCreateAt,
		).
		Values(
			group.CytologyID,
			group.SegType,
			group.GroupType,
			group.IsAI,
			group.Details,
			group.CreateAt,
		).
		Suffix("RETURNING id")

	var id int
	err := q.Runner().Getx(q.Context(), &id, query)
	if err != nil {
		return 0, repoEntity.WrapDBError(err)
	}

	return id, nil
}
