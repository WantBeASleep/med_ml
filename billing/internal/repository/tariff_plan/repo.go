package tariff_plan

import (
	"billing/internal/repository/tariff_plan/entity"

	daolib "github.com/WantBeASleep/goooool/daolib"
	"github.com/google/uuid"
)

const (
	table = "tariff_plan"

	columnID          = "id"
	columnName        = "name"
	columnDescription = "description"
	columnPrice       = "price"
	columnDuration    = "duration"
)

type Repository interface {
	InsertTariffPlan(plan entity.TariffPlan) error

	GetTariffPlanByID(id uuid.UUID) (entity.TariffPlan, error)
	ListTariffPlans() ([]entity.TariffPlan, error)

	UpdateTariffPlan(plan entity.TariffPlan) (int64, error)

	DeleteTariffPlan(id uuid.UUID) error
}

type repo struct {
	*daolib.BaseQuery
}

func NewRepo() *repo {
	return &repo{}
}

func (q *repo) SetBaseQuery(baseQuery *daolib.BaseQuery) {
	q.BaseQuery = baseQuery
}
