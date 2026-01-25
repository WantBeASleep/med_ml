package original_image

import (
	"database/sql"
	"errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	daoEntity "cytology/internal/repository/entity"
	"cytology/internal/repository/original_image/entity"
)

func (q *repo) GetOriginalImageByID(id uuid.UUID) (entity.OriginalImage, error) {
	query := q.QueryBuilder().
		Select(
			columnID,
			columnCytologyID,
			columnImagePath,
			columnCreateDate,
			columnDelayTime,
			columnViewedFlag,
		).
		From(table).
		Where(sq.Eq{
			columnID: id,
		})

	var img entity.OriginalImage
	if err := q.Runner().Getx(q.Context(), &img, query); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.OriginalImage{}, daoEntity.ErrNotFound
		}
		return entity.OriginalImage{}, err
	}

	return img, nil
}

func (q *repo) GetOriginalImagesByCytologyID(cytologyID uuid.UUID) ([]entity.OriginalImage, error) {
	query := q.QueryBuilder().
		Select(
			columnID,
			columnCytologyID,
			columnImagePath,
			columnCreateDate,
			columnDelayTime,
			columnViewedFlag,
		).
		From(table).
		Where(sq.Eq{
			columnCytologyID: cytologyID,
		})

	var images []entity.OriginalImage
	if err := q.Runner().Selectx(q.Context(), &images, query); err != nil {
		return nil, err
	}

	if len(images) == 0 {
		return nil, daoEntity.ErrNotFound
	}

	return images, nil
}
