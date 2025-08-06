package payment

import (
	"database/sql"
	"errors"

	"billing/internal/domain"
	"billing/internal/repository/payment/entity"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

func (q *repo) GetPaymentByID(id uuid.UUID) (entity.Payment, error) {
	query := q.QueryBuilder().
		Select(
			columnID,
			columnUserID,
			columnSubscriptionID,
			columnAmount,
			columnStatus,
			columnPaymentProviderID,
			columnPspToken,
			columnCreatedAt,
			columnUpdatedAt,
		).
		From(paymentTable).
		Where(sq.Eq{
			columnID: id,
		})

	var payment entity.Payment
	if err := q.Runner().Getx(q.Context(), &payment, query); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.Payment{}, sql.ErrNoRows
		}
		return entity.Payment{}, err
	}

	return payment, nil
}

func (q *repo) GetPaymentByProviderID(providerPaymentID string, paymentProviderID uuid.UUID) (entity.Payment, error) {
	query := q.QueryBuilder().
		Select(
			columnID,
			columnUserID,
			columnSubscriptionID,
			columnAmount,
			columnStatus,
			columnPaymentProviderID,
			columnPspToken,
			columnCreatedAt,
			columnUpdatedAt,
		).
		From(paymentTable).
		Where(sq.Eq{
			columnPspToken:          providerPaymentID,
			columnPaymentProviderID: paymentProviderID,
		})

	var payment entity.Payment
	if err := q.Runner().Getx(q.Context(), &payment, query); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.Payment{}, sql.ErrNoRows
		}
		return entity.Payment{}, err
	}

	return payment, nil
}

func (q *repo) CheckExistPaymentByStatusAndUserID(status domain.PaymentStatus, userID uuid.UUID) (bool, error) {
	query := q.QueryBuilder().
		Select("1").
		From(paymentTable).
		Where(sq.Eq{
			columnStatus: status,
			columnUserID: userID,
		}).
		Limit(1)

	var exists bool
	err := q.Runner().Getx(q.Context(), &exists, query)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}

	return exists, nil
}
