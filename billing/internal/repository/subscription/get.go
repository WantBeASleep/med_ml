package subscription

import (
	"billing/internal/domain"
	"billing/internal/repository/subscription/entity"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

func (q *repo) GetSubscriptionByID(id uuid.UUID) (entity.Subscription, error) {
	query := q.QueryBuilder().
		Select(
			columnID,
			columnUserID,
			columnTariffPlanID,
			columnStartDate,
			columnEndDate,
			columnStatus,
		).
		From(table).
		Where(sq.Eq{
			columnID: id,
		})

	var subscription entity.Subscription
	if err := q.Runner().Getx(q.Context(), &subscription, query); err != nil {
		return entity.Subscription{}, err
	}

	return subscription, nil
}

func (q *repo) GetSubscriptionsByUserID(userID uuid.UUID) ([]entity.Subscription, error) {
	query := q.QueryBuilder().
		Select(
			columnID,
			columnUserID,
			columnTariffPlanID,
			columnStartDate,
			columnEndDate,
			columnStatus,
		).
		From(table).
		Where(sq.Eq{
			columnUserID: userID,
		})

	var subscriptions []entity.Subscription
	if err := q.Runner().Selectx(q.Context(), &subscriptions, query); err != nil {
		return nil, err
	}

	return subscriptions, nil
}

func (q *repo) GetSubscrptionsByStatusAndUserID(status string, userID uuid.UUID) ([]entity.Subscription, error) {
	query := q.QueryBuilder().
		Select(
			columnID,
			columnUserID,
			columnTariffPlanID,
			columnStartDate,
			columnEndDate,
			columnStatus,
		).
		From(table).
		Where(sq.Eq{
			columnUserID: userID,
			columnStatus: status,
		})

	var subscriptions []entity.Subscription
	if err := q.Runner().Selectx(q.Context(), &subscriptions, query); err != nil {
		return nil, err
	}

	return subscriptions, nil
}

func (q *repo) CheckExistSubscrptionByStatusAndUserID(status string, userID uuid.UUID) (bool, error) {
	query := q.QueryBuilder().
		Select(columnID).
		Prefix("SELECT EXISTS (").
		From(table).
		Where(sq.Eq{
			columnUserID: userID,
			columnStatus: status,
		}).
		Suffix(")")

	var exists bool
	if err := q.Runner().Getx(q.Context(), &exists, query); err != nil {
		return false, err
	}

	return exists, nil
}

func (q *repo) GetAllActiveSubscriptions() ([]entity.Subscription, error) {
	query := q.QueryBuilder().
		Select(
			columnID,
			columnUserID,
			columnTariffPlanID,
			columnStartDate,
			columnEndDate,
			columnStatus,
		).
		From(table).
		Where(sq.Eq{
			columnStatus: domain.SubActive,
		})

	var subscriptions []entity.Subscription
	if err := q.Runner().Selectx(q.Context(), &subscriptions, query); err != nil {
		return nil, err
	}

	return subscriptions, nil
}
