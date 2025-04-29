package repository

import (
	"billing/internal/domain"
	"billing/internal/repository/entity"

	sq "github.com/Masterminds/squirrel"
	"github.com/WantBeASleep/goooool/daolib"
	"github.com/google/uuid"
)

const subscriptionTable = "subscription"

type SubscriptionQuery interface {
	InsertSubscription(subscription entity.Subscription) error
	GetSubscriptionByID(id uuid.UUID) (entity.Subscription, error)
	UpdateSubscription(subscription entity.Subscription) (int64, error)
	GetSubscriptionsByUserID(userID uuid.UUID) ([]entity.Subscription, error)
	GetSubscrptionsByStatusAndUserID(status string, userID uuid.UUID) ([]entity.Subscription, error)
	CheckExistSubscrptionByStatusAndUserID(status string, userID uuid.UUID) (bool, error)
	GetAllActiveSubscriptions() ([]entity.Subscription, error)
	SetSubscriptionsStatusBatch(ids []uuid.UUID, status domain.SubscriptionStatus) error
}

type subscriptionQuery struct {
	*daolib.BaseQuery
}

func (q *subscriptionQuery) SetBaseQuery(baseQuery *daolib.BaseQuery) {
	q.BaseQuery = baseQuery
}

func (q *subscriptionQuery) InsertSubscription(subscription entity.Subscription) error {
	query := q.QueryBuilder().
		Insert(subscriptionTable).
		Columns(
			"id",
			"user_id",
			"tariff_plan_id",
			"start_date",
			"end_date",
			"status",
		).
		Values(
			subscription.ID,
			subscription.UserID,
			subscription.TariffPlanID,
			subscription.StartDate,
			subscription.EndDate,
			subscription.Status,
		)

	_, err := q.Runner().Execx(q.Context(), query)
	if err != nil {
		return err
	}

	return nil
}

func (q *subscriptionQuery) GetSubscriptionByID(id uuid.UUID) (entity.Subscription, error) {
	query := q.QueryBuilder().
		Select(
			"id",
			"user_id",
			"tariff_plan_id",
			"start_date",
			"end_date",
			"status",
		).
		From(subscriptionTable).
		Where(sq.Eq{
			"id": id,
		})

	var subscription entity.Subscription
	if err := q.Runner().Getx(q.Context(), &subscription, query); err != nil {
		return entity.Subscription{}, err
	}

	return subscription, nil
}

func (q *subscriptionQuery) GetSubscriptionsByUserID(userID uuid.UUID) ([]entity.Subscription, error) {
	query := q.QueryBuilder().
		Select(
			"id",
			"user_id",
			"tariff_plan_id",
			"start_date",
			"end_date",
			"status",
		).
		From(subscriptionTable).
		Where(sq.Eq{
			"user_id": userID,
		})

	var subscription []entity.Subscription
	if err := q.Runner().Selectx(q.Context(), &subscription, query); err != nil {
		return nil, err
	}

	return subscription, nil
}

func (q *subscriptionQuery) GetSubscrptionsByStatusAndUserID(status string, userID uuid.UUID) ([]entity.Subscription, error) {
	query := q.QueryBuilder().
		Select(
			"id",
			"user_id",
			"tariff_plan_id",
			"start_date",
			"end_date",
			"status",
		).
		From(subscriptionTable).
		Where(sq.Eq{
			"user_id": userID,
			"status":  status,
		})

	var subscription []entity.Subscription
	if err := q.Runner().Selectx(q.Context(), &subscription, query); err != nil {
		return nil, err
	}

	return subscription, nil
}

func (q *subscriptionQuery) UpdateSubscription(subscription entity.Subscription) (int64, error) {
	query := q.QueryBuilder().
		Update(subscriptionTable).
		SetMap(sq.Eq{
			"start_date": subscription.StartDate,
			"end_date":   subscription.EndDate,
			"status":     subscription.Status,
		}).
		Where(sq.Eq{
			"id": subscription.ID,
		})

	res, err := q.Runner().Execx(q.Context(), query)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func (q *subscriptionQuery) CheckExistSubscrptionByStatusAndUserID(status string, userID uuid.UUID) (bool, error) {
	query := q.QueryBuilder().
		Select(
			"id",
			"user_id",
			"tariff_plan_id",
			"start_date",
			"end_date",
			"status",
		).
		Prefix("SELECT EXISTS (").
		From(subscriptionTable).
		Where(sq.Eq{
			"user_id": userID,
			"status":  status,
		}).
		Suffix(")")

	var exists bool
	if err := q.Runner().Getx(q.Context(), &exists, query); err != nil {
		return false, err
	}

	return exists, nil
}

func (q *subscriptionQuery) GetAllActiveSubscriptions() ([]entity.Subscription, error) {
	query := q.QueryBuilder().
		Select(
			"id",
			"user_id",
			"tariff_plan_id",
			"start_date",
			"end_date",
			"status",
		).
		From(subscriptionTable).
		Where(sq.Eq{
			"status": domain.SubActive,
		})

	var subscriptions []entity.Subscription
	if err := q.Runner().Selectx(q.Context(), &subscriptions, query); err != nil {
		return nil, err
	}

	return subscriptions, nil
}

func (q *subscriptionQuery) SetSubscriptionsStatusBatch(ids []uuid.UUID, status domain.SubscriptionStatus) error {
	query := q.QueryBuilder().
		Update(subscriptionTable).
		Set("status", status).
		Where(sq.Eq{"id": ids})

	_, err := q.Runner().Execx(q.Context(), query)
	return err
}
