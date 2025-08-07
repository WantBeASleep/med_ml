package tariff_plan

import (
	"billing/internal/repository/tariff_plan/entity"
)

func (q *repo) InsertTariffPlan(plan entity.TariffPlan) error {
	query := q.QueryBuilder().
		Insert(table).
		Columns(
			columnID,
			columnName,
			columnDescription,
			columnPrice,
			columnDuration,
		).
		Values(
			plan.ID,
			plan.Name,
			plan.Description,
			plan.Price,
			plan.Duration,
		)

	_, err := q.Runner().Execx(q.Context(), query)
	if err != nil {
		return err
	}

	return nil
}
