package user

import (
	"auth/internal/repository/user/entity"
)

func (r *repo) InsertUser(user entity.User) error {
	query := r.QueryBuilder().
		Insert(table).
		Columns(
			columnID,
			columnEmail,
			columnPassword,
			columnRefreshToken,
			columnRole,
		).
		Values(
			user.Id,
			user.Email,
			user.Password,
			user.RefreshToken,
			user.Role,
		)

	_, err := r.Runner().Execx(r.Context(), query)
	if err != nil {
		return err
	}

	return nil
}
