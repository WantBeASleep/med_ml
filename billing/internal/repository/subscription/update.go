package subscription

import (
	"billing/internal/domain"
	"billing/internal/repository/subscription/entity"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

func (q *repo) UpdateSubscription(subscription entity.Subscription) (int64, error) {
	query := q.QueryBuilder().
		Update(table).
		SetMap(sq.Eq{
			columnStartDate: subscription.StartDate,
			columnEndDate:   subscription.EndDate,
			columnStatus:    subscription.Status,
		}).
		Where(sq.Eq{
			columnID: subscription.ID,
		})

	res, err := q.Runner().Execx(q.Context(), query)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func (q *repo) SetSubscriptionsStatusBatch(ids []uuid.UUID, status domain.SubscriptionStatus) error {
	query := q.QueryBuilder().
		Update(table).
		Set(columnStatus, status).
		Where(sq.Eq{columnID: ids})

	_, err := q.Runner().Execx(q.Context(), query)
	return err
}
