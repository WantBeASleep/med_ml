package entity

import (
	"time"

	"billing/internal/domain"

	"github.com/google/uuid"
)

type Subscription struct {
	ID           uuid.UUID `db:"id"`
	UserID       uuid.UUID `db:"user_id"`
	TariffPlanID uuid.UUID `db:"tariff_plan_id"`
	StartDate    time.Time `db:"start_date"`
	EndDate      time.Time `db:"end_date"`
	Status       string    `db:"status"`
}

func (Subscription) FromDomain(p domain.Subscription) Subscription {
	return Subscription{
		ID:           p.ID,
		UserID:       p.UserID,
		TariffPlanID: p.TariffPlanID,
		StartDate:    p.StartDate,
		EndDate:      p.EndDate,
		Status:       string(p.Status),
	}
}

func (p Subscription) ToDomain() domain.Subscription {
	return domain.Subscription{
		ID:           p.ID,
		UserID:       p.UserID,
		TariffPlanID: p.TariffPlanID,
		StartDate:    p.StartDate,
		EndDate:      p.EndDate,
		Status:       domain.SubscriptionStatus(p.Status),
	}
}
