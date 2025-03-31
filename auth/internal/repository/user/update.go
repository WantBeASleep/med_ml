package user

import (
	uentity "auth/internal/repository/user/entity"

	sq "github.com/Masterminds/squirrel"
)

func (r *repo) UpdateUser(user uentity.User) error {
	query := r.QueryBuilder().
		Update(table).
		SetMap(sq.Eq{
			columnPassword:     user.Password,
			columnRefreshToken: user.RefreshToken,
		}).
		Where(sq.Eq{
			columnID: user.Id,
		})

	_, err := r.Runner().Execx(r.Context(), query)
	return err
}
