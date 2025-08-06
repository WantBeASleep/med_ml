package subscription

import (
	"billing/internal/domain"
	"billing/internal/repository/subscription/entity"

	daolib "github.com/WantBeASleep/goooool/daolib"
	"github.com/google/uuid"
)

const (
	table = "subscription"

	columnID           = "id"
	columnUserID       = "user_id"
	columnTariffPlanID = "tariff_plan_id"
	columnStartDate    = "start_date"
	columnEndDate      = "end_date"
	columnStatus       = "status"
)

type Repository interface {
	InsertSubscription(subscription entity.Subscription) error

	GetSubscriptionByID(id uuid.UUID) (entity.Subscription, error)
	GetSubscriptionsByUserID(userID uuid.UUID) ([]entity.Subscription, error)
	GetSubscrptionsByStatusAndUserID(status string, userID uuid.UUID) ([]entity.Subscription, error)
	CheckExistSubscrptionByStatusAndUserID(status string, userID uuid.UUID) (bool, error)
	GetAllActiveSubscriptions() ([]entity.Subscription, error)

	UpdateSubscription(subscription entity.Subscription) (int64, error)
	SetSubscriptionsStatusBatch(ids []uuid.UUID, status domain.SubscriptionStatus) error
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
