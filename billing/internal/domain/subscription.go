package domain

import (
	"time"

	"github.com/google/uuid"
)

type SubscriptionStatus string

const (
	SubPendingPayment SubscriptionStatus = "pending_payment"
	SubActive         SubscriptionStatus = "active"
	SubCancelled      SubscriptionStatus = "cancelled"
)

type Subscription struct {
	ID           uuid.UUID
	UserID       uuid.UUID
	TariffPlanID uuid.UUID
	StartDate    time.Time
	EndDate      time.Time
	Status       SubscriptionStatus
}
