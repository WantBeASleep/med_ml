package image

import (
	daoEntity "uzi/internal/repository/entity"
	"uzi/internal/repository/image/entity"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

func (q *repo) GetImagesByUziID(uziID uuid.UUID) ([]entity.Image, error) {
	query := q.QueryBuilder().
		Select(
			columnId,
			columnPage,
		).
		From(table).
		Where(sq.Eq{
			columnUziId: uziID,
		})

	var images []entity.Image
	if err := q.Runner().Selectx(q.Context(), &images, query); err != nil {
		return nil, err
	}

	if len(images) == 0 {
		return nil, daoEntity.ErrNotFound
	}

	return images, nil
}
