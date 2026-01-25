package cytology_image

import (
	sq "github.com/Masterminds/squirrel"

	"cytology/internal/repository/cytology_image/entity"
	repoEntity "cytology/internal/repository/entity"
)

func (q *repo) UpdateCytologyImage(img entity.CytologyImage) error {
	updateMap := sq.Eq{
		columnIsLast: img.IsLast,
	}

	if img.DiagnosticMarking.Valid {
		updateMap[columnDiagnosticMarking] = img.DiagnosticMarking
	}
	if img.MaterialType.Valid {
		updateMap[columnMaterialType] = img.MaterialType
	}
	if img.Calcitonin.Valid {
		updateMap[columnCalcitonin] = img.Calcitonin
	}
	if img.CalcitoninInFlush.Valid {
		updateMap[columnCalcitoninInFlush] = img.CalcitoninInFlush
	}
	if img.Thyroglobulin.Valid {
		updateMap[columnThyroglobulin] = img.Thyroglobulin
	}
	if img.Details.Valid {
		updateMap[columnDetails] = img.Details.String
	}

	query := q.QueryBuilder().
		Update(table).
		SetMap(updateMap).
		Where(sq.Eq{
			columnID: img.Id,
		})

	_, err := q.Runner().Execx(q.Context(), query)
	if err != nil {
		return repoEntity.WrapDBError(err)
	}

	return nil
}
