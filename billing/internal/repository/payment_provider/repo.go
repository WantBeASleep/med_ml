package payment_provider

import (
	"billing/internal/repository/payment_provider/entity"

	daolib "github.com/WantBeASleep/goooool/daolib"
	"github.com/google/uuid"
)

const (
	table = "payment_provider"

	columnID       = "id"
	columnName     = "name"
	columnIsActive = "is_active"
)

type Repository interface {
	InsertPaymentProvider(provider entity.PaymentProvider) error

	GetPaymentProviderByID(id uuid.UUID) (entity.PaymentProvider, error)
	GetPaymentProviderByName(name string) (entity.PaymentProvider, error)
	ListPaymentProviders() ([]entity.PaymentProvider, error)

	UpdatePaymentProvider(provider entity.PaymentProvider) (int64, error)
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
