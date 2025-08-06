package tariff_plan

import (
	"billing/internal/repository/tariff_plan/entity"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

func (q *repo) GetTariffPlanByID(id uuid.UUID) (entity.TariffPlan, error) {
	query := q.QueryBuilder().
		Select(
			columnID,
			columnName,
			columnDescription,
			columnPrice,
			columnDuration,
		).
		From(table).
		Where(sq.Eq{
			columnID: id,
		})

	var plan entity.TariffPlan
	if err := q.Runner().Getx(q.Context(), &plan, query); err != nil {
		return entity.TariffPlan{}, err
	}

	return plan, nil
}

func (q *repo) ListTariffPlans() ([]entity.TariffPlan, error) {
	query := q.QueryBuilder().
		Select(
			columnID,
			columnName,
			columnDescription,
			columnPrice,
			columnDuration,
		).
		From(table)

	var plans []entity.TariffPlan
	if err := q.Runner().Selectx(q.Context(), &plans, query); err != nil {
		return nil, err
	}

	return plans, nil
}
