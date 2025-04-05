package user

import (
	"database/sql"
	"errors"

	"auth/internal/repository/entity"
	uentity "auth/internal/repository/user/entity"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

func (r *repo) GetUserByID(id uuid.UUID) (uentity.User, error) {
	query := r.QueryBuilder().
		Select(
			columnID,
			columnEmail,
			columnPassword,
			columnRefreshToken,
			columnRole,
		).
		From(table).
		Where(sq.Eq{
			columnID: id,
		})

	var user uentity.User
	if err := r.Runner().Getx(r.Context(), &user, query); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return uentity.User{}, entity.ErrNotFound
		}

		return uentity.User{}, err
	}

	return user, nil
}

func (r *repo) GetUserByEmail(email string) (uentity.User, error) {
	query := r.QueryBuilder().
		Select(
			columnID,
			columnEmail,
			columnPassword,
			columnRefreshToken,
			columnRole,
		).
		From(table).
		Where(sq.Eq{
			columnEmail: email,
		})

	var user uentity.User
	if err := r.Runner().Getx(r.Context(), &user, query); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return uentity.User{}, entity.ErrNotFound
		}

		return uentity.User{}, err
	}

	return user, nil
}
