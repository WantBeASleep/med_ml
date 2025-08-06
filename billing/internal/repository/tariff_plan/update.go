package tariff_plan

import (
	"billing/internal/repository/tariff_plan/entity"

	sq "github.com/Masterminds/squirrel"
)

func (q *repo) UpdateTariffPlan(plan entity.TariffPlan) (int64, error) {
	query := q.QueryBuilder().
		Update(table).
		SetMap(sq.Eq{
			columnName:        plan.Name,
			columnDescription: plan.Description,
			columnPrice:       plan.Price,
			columnDuration:    plan.Duration,
		}).
		Where(sq.Eq{
			columnID: plan.ID,
		})

	res, err := q.Runner().Execx(q.Context(), query)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
