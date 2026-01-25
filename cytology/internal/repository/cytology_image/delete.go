package cytology_image

import (
	"github.com/google/uuid"

	sq "github.com/Masterminds/squirrel"

	repoEntity "cytology/internal/repository/entity"
)

func (r *repo) DeleteCytologyImage(id uuid.UUID) error {
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
