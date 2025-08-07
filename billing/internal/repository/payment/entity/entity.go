package entity

import (
	"time"

	"billing/internal/domain"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Payment struct {
	ID                uuid.UUID       `db:"id"`
	UserID            uuid.UUID       `db:"user_id"`
	SubscriptionID    uuid.UUID       `db:"subscription_id"`
	Amount            decimal.Decimal `db:"amount"`
	Status            string          `db:"status"`
	PaymentProviderID uuid.UUID       `db:"payment_provider_id"`
	PspToken          string          `db:"psp_token"`
	CreatedAt         time.Time       `db:"created_at"`
	UpdatedAt         time.Time       `db:"updated_at"`
}

func (Payment) FromDomain(p domain.Payment) Payment {
	return Payment{
		ID:                p.ID,
		UserID:            p.UserID,
		SubscriptionID:    p.SubscriptionID,
		Amount:            p.Amount,
		Status:            string(p.Status),
		PaymentProviderID: p.PaymentProviderID,
		PspToken:          p.PspToken,
		CreatedAt:         p.CreatedAt,
		UpdatedAt:         p.UpdatedAt,
	}
}

func (p Payment) ToDomain() domain.Payment {
	return domain.Payment{
		ID:                p.ID,
		UserID:            p.UserID,
		SubscriptionID:    p.SubscriptionID,
		Amount:            p.Amount,
		Status:            domain.PaymentStatus(p.Status),
		PaymentProviderID: p.PaymentProviderID,
		PspToken:          p.PspToken,
		CreatedAt:         p.CreatedAt,
		UpdatedAt:         p.UpdatedAt,
	}
}
