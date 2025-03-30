package uzi

import (
	"database/sql"
	"errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	daoEntity "uzi/internal/repository/entity"
	"uzi/internal/repository/uzi/entity"
)

func (q *repo) GetUziByID(id uuid.UUID) (entity.Uzi, error) {
	query := q.QueryBuilder().
		Select(
			columnID,
			columnProjection,
			columnChecked,
			columnExternalID,
			columnAuthor,
			columnDeviceID,
			columnStatus,
			columnDescription,
			columnCreateAt,
		).
		From(uziTable).
		Where(sq.Eq{
			columnID: id,
		})

	var uzi entity.Uzi
	if err := q.Runner().Getx(q.Context(), &uzi, query); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.Uzi{}, daoEntity.ErrNotFound
		}
		return entity.Uzi{}, err
	}

	return uzi, nil
}

func (q *repo) GetUzisByExternalID(externalID uuid.UUID) ([]entity.Uzi, error) {
	query := q.QueryBuilder().
		Select(
			columnID,
			columnProjection,
			columnChecked,
			columnExternalID,
			columnAuthor,
			columnDeviceID,
			columnStatus,
			columnDescription,
			columnCreateAt,
		).
		From(uziTable).
		Where(sq.Eq{
			columnExternalID: externalID,
		})

	var uzi []entity.Uzi
	if err := q.Runner().Selectx(q.Context(), &uzi, query); err != nil {
		return nil, err
	}

	if len(uzi) == 0 {
		return nil, daoEntity.ErrNotFound
	}

	return uzi, nil
}

func (q *repo) GetUzisByAuthor(author uuid.UUID) ([]entity.Uzi, error) {
	query := q.QueryBuilder().
		Select(
			columnID,
			columnProjection,
			columnChecked,
			columnExternalID,
			columnAuthor,
			columnDeviceID,
			columnStatus,
			columnDescription,
			columnCreateAt,
		).
		From(uziTable).
		Where(sq.Eq{
			columnAuthor: author,
		})

	var uzi []entity.Uzi
	if err := q.Runner().Selectx(q.Context(), &uzi, query); err != nil {
		return nil, err
	}

	if len(uzi) == 0 {
		return nil, daoEntity.ErrNotFound
	}

	return uzi, nil
}

func (q *repo) CheckExist(id uuid.UUID) (bool, error) {
	query := q.QueryBuilder().
		Select(
			columnID,
			columnProjection,
			columnChecked,
			columnExternalID,
			columnAuthor,
			columnDeviceID,
			columnStatus,
			columnDescription,
			columnCreateAt,
		).
		Prefix("SELECT EXISTS (").
		From(uziTable).
		Where(sq.Eq{
			columnID: id,
		}).
		Suffix(")")

	var exists bool
	if err := q.Runner().Getx(q.Context(), &exists, query); err != nil {
		return false, err
	}

	return exists, nil
}
