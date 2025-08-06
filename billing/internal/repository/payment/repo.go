package payment

import (
	"billing/internal/repository/payment/entity"

	"billing/internal/domain"

	daolib "github.com/WantBeASleep/goooool/daolib"
	"github.com/google/uuid"
)

const (
	paymentTable = "payment"

	columnID                = "id"
	columnUserID            = "user_id"
	columnSubscriptionID    = "subscription_id"
	columnAmount            = "amount"
	columnStatus            = "status"
	columnPaymentProviderID = "payment_provider_id"
	columnPspToken          = "psp_token"
	columnCreatedAt         = "created_at"
	columnUpdatedAt         = "updated_at"
)

type Repository interface {
	InsertPayment(payment entity.Payment) error

	GetPaymentByID(id uuid.UUID) (entity.Payment, error)
	GetPaymentByProviderID(providerPaymentID string, paymentProviderID uuid.UUID) (entity.Payment, error)
	CheckExistPaymentByStatusAndUserID(status domain.PaymentStatus, userID uuid.UUID) (bool, error)

	UpdatePayment(payment entity.Payment) (int64, error)
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
