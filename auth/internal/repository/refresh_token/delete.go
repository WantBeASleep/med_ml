package refresh_token

import (
	rtentity "auth/internal/repository/refresh_token/entity"

	sq "github.com/Masterminds/squirrel"
)

func (r *repo) DeleteRefreshTokens(refreshTokens []rtentity.RefreshToken) error {
	cond := sq.Or{}
	for _, pk := range refreshTokens {
		cond = append(cond, sq.Eq{
			columnID:           pk.Id,
			columnRefreshToken: pk.RefreshToken,
		})
	}

	query := r.QueryBuilder().
		Delete(table).
		Where(cond)

	_, err := r.Runner().Execx(r.Context(), query)
	return err
}
