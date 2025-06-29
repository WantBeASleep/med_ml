package refresh_token

import (
	"database/sql"
	"errors"

	rtentity "auth/internal/repository/refresh_token/entity"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

func (r *repo) ExistsRefreshToken(id uuid.UUID, refreshToken string) (bool, error) {
	query := r.QueryBuilder().
		Select(
			columnID,
			columnRefreshToken,
		).
		From(table).
		Where(sq.Eq{
			columnID:           id,
			columnRefreshToken: refreshToken,
		})

	var entity rtentity.RefreshToken
	if err := r.Runner().Getx(r.Context(), &entity, query); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}
