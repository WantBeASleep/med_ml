package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type TariffPlan struct {
	ID          uuid.UUID
	Name        string
	Description string
	Price       decimal.Decimal
	Duration    time.Duration
}
