package repository

import (
	"billing/internal/repository/entity"

	sq "github.com/Masterminds/squirrel"
	"github.com/WantBeASleep/goooool/daolib"
	"github.com/google/uuid"
)

const tariffPlanTable = "tariff_plan"

type TariffPlanQuery interface {
	InsertTariffPlan(plan entity.TariffPlan) error
	GetTariffPlanByID(id uuid.UUID) (entity.TariffPlan, error)
	UpdateTariffPlan(plan entity.TariffPlan) (int64, error)
	DeleteTariffPlan(id uuid.UUID) error
	ListTariffPlans() ([]entity.TariffPlan, error)
}

type tariffPlanQuery struct {
	*daolib.BaseQuery
}

func (q *tariffPlanQuery) SetBaseQuery(baseQuery *daolib.BaseQuery) {
	q.BaseQuery = baseQuery
}

func (q *tariffPlanQuery) InsertTariffPlan(plan entity.TariffPlan) error {
	query := q.QueryBuilder().
		Insert(tariffPlanTable).
		Columns(
			"id",
			"name",
			"description",
			"price",
			"duration",
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

func (q *tariffPlanQuery) GetTariffPlanByID(id uuid.UUID) (entity.TariffPlan, error) {
	query := q.QueryBuilder().
		Select(
			"id",
			"name",
			"description",
			"price",
			"duration",
		).
		From(tariffPlanTable).
		Where(sq.Eq{
			"id": id,
		})

	var plan entity.TariffPlan
	if err := q.Runner().Getx(q.Context(), &plan, query); err != nil {
		return entity.TariffPlan{}, err
	}

	return plan, nil
}

func (q *tariffPlanQuery) UpdateTariffPlan(plan entity.TariffPlan) (int64, error) {
	query := q.QueryBuilder().
		Update(tariffPlanTable).
		SetMap(sq.Eq{
			"name":        plan.Name,
			"description": plan.Description,
			"price":       plan.Price,
			"duration":    plan.Duration,
		}).
		Where(sq.Eq{
			"id": plan.ID,
		})

	res, err := q.Runner().Execx(q.Context(), query)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func (q *tariffPlanQuery) DeleteTariffPlan(id uuid.UUID) error {
	query := q.QueryBuilder().
		Delete(tariffPlanTable).
		Where(sq.Eq{
			"id": id,
		})

	_, err := q.Runner().Execx(q.Context(), query)
	if err != nil {
		return err
	}

	return nil
}

func (q *tariffPlanQuery) ListTariffPlans() ([]entity.TariffPlan, error) {
	query := q.QueryBuilder().
		Select(
			"id",
			"name",
			"description",
			"price",
			"duration",
		).
		From(tariffPlanTable)

	var plans []entity.TariffPlan
	if err := q.Runner().Selectx(q.Context(), &plans, query); err != nil {
		return nil, err
	}

	return plans, nil
}
