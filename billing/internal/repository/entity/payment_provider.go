package entity

import (
	"billing/internal/domain"

	"github.com/google/uuid"
)

type PaymentProvider struct {
	ID       uuid.UUID `db:"id"`
	Name     string    `db:"name"`
	IsActive bool      `db:"is_active"`
}

func (PaymentProvider) FromDomain(p domain.PaymentProvider) PaymentProvider {
	return PaymentProvider{
		ID:       p.ID,
		Name:     p.Name,
		IsActive: p.IsActive,
	}
}

func (p PaymentProvider) ToDomain() domain.PaymentProvider {
	return domain.PaymentProvider{
		ID:       p.ID,
		Name:     p.Name,
		IsActive: p.IsActive,
	}
}
