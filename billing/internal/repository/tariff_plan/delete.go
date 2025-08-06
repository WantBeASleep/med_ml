package tariff_plan

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

func (q *repo) DeleteTariffPlan(id uuid.UUID) error {
	query := q.QueryBuilder().
		Delete(table).
		Where(sq.Eq{
			columnID: id,
		})

	_, err := q.Runner().Execx(q.Context(), query)
	if err != nil {
		return err
	}

	return nil
}
