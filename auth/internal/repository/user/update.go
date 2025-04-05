package user

import (
	"github.com/google/uuid"

	sq "github.com/Masterminds/squirrel"
)

func (r *repo) UpdateUserPassword(id uuid.UUID, password string) error {
	query := r.QueryBuilder().
		Update(table).
		SetMap(sq.Eq{
			columnPassword: password,
		}).
		Where(sq.Eq{
			columnID: id,
		})

	_, err := r.Runner().Execx(r.Context(), query)
	return err
}

func (r *repo) UpdateUserRefreshToken(id uuid.UUID, refreshToken string) error {
	query := r.QueryBuilder().
		Update(table).
		SetMap(sq.Eq{
			columnRefreshToken: refreshToken,
		}).
		Where(sq.Eq{
			columnID: id,
		})

	_, err := r.Runner().Execx(r.Context(), query)
	return err
}
