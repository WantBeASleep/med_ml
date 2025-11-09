package uzi

import (
	"github.com/google/uuid"

	sq "github.com/Masterminds/squirrel"

	repoEntity "uzi/internal/repository/entity"
)

func (r *repo) DeleteUzi(id uuid.UUID) error {
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
