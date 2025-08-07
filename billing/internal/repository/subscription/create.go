package subscription

import (
	"billing/internal/repository/subscription/entity"
)

func (q *repo) InsertSubscription(subscription entity.Subscription) error {
	query := q.QueryBuilder().
		Insert(table).
		Columns(
			columnID,
			columnUserID,
			columnTariffPlanID,
			columnStartDate,
			columnEndDate,
			columnStatus,
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
