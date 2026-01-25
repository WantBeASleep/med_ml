package segmentation_group

import (
	sq "github.com/Masterminds/squirrel"

	repoEntity "cytology/internal/repository/entity"
)

func (r *repo) DeleteSegmentationGroup(id int) error {
	query := r.QueryBuilder().
		Delete(table).
		Where(sq.Eq{
			columnID: id,
		})

	_, err := r.Runner().Execx(r.Context(), query)
	if err != nil {
		return repoEntity.WrapDBError(err)
	}

	return nil
}
