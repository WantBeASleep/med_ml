package echographic

import (
	"database/sql"
	"errors"

	"uzi/internal/repository/echographic/entity"
	daoEntity "uzi/internal/repository/entity"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

func (q *repo) GetEchographicByID(id uuid.UUID) (entity.Echographic, error) {
	query := q.QueryBuilder().
		Select(
			columnID,
			columnContors,
			columnLeftLobeLength,
			columnLeftLobeWidth,
			columnLeftLobeThick,
			columnLeftLobeVolum,
			columnRightLobeLength,
			columnRightLobeWidth,
			columnRightLobeThick,
			columnRightLobeVolum,
			columnGlandVolum,
			columnIsthmus,
			columnStruct,
			columnEchogenicity,
			columnRegionalLymph,
			columnVascularization,
			columnLocation,
			columnAdditional,
			columnConclusion,
		).
		From(table).
		Where(sq.Eq{
			columnID: id,
		})

	var echographic entity.Echographic
	if err := q.Runner().Getx(q.Context(), &echographic, query); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.Echographic{}, daoEntity.ErrNotFound
		}
		return entity.Echographic{}, err
	}

	return echographic, nil
}
