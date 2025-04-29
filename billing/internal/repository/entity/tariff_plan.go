package entity

import (
	"time"

	"github.com/shopspring/decimal"

	"billing/internal/domain"

	"github.com/google/uuid"
)

type TariffPlan struct {
	ID          uuid.UUID       `db:"id"`
	Name        string          `db:"name"`
	Description string          `db:"description"`
	Price       decimal.Decimal `db:"price"`
	Duration    Duration        `db:"duration"`
}

func (TariffPlan) FromDomain(p domain.TariffPlan) TariffPlan {
	return TariffPlan{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		Duration:    Duration(p.Duration),
	}
}

func (p TariffPlan) ToDomain() domain.TariffPlan {
	return domain.TariffPlan{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		Duration:    time.Duration(p.Duration),
	}
}
