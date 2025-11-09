package refresh_token

import (
	rtentity "auth/internal/repository/refresh_token/entity"
	repoEntity "auth/internal/repository/entity"
)

func (r *repo) InsertRefreshToken(refreshToken rtentity.RefreshToken) error {
	query := r.QueryBuilder().
		Insert(table).
		Columns(
			columnID,
			columnRefreshToken,
		).
		Values(
			refreshToken.Id,
			refreshToken.RefreshToken,
		)

	_, err := r.Runner().Execx(r.Context(), query)
	if err != nil {
		return repoEntity.WrapDBError(err)
	}

	return nil
}
