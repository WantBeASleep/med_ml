package repository

import (
	"database/sql"
	"errors"

	"billing/internal/domain"
	"billing/internal/repository/entity"

	sq "github.com/Masterminds/squirrel"
	"github.com/WantBeASleep/goooool/daolib"
	"github.com/google/uuid"
)

const paymentTable = "payment"

type PaymentQuery interface {
	InsertPayment(payment entity.Payment) error
	GetPaymentByID(id uuid.UUID) (entity.Payment, error)
	UpdatePayment(payment entity.Payment) (int64, error)
	GetPaymentByProviderID(providerPaymentID string, paymentProviderID uuid.UUID) (entity.Payment, error)
	CheckExistPaymentByStatusAndUserID(status domain.PaymentStatus, userID uuid.UUID) (bool, error)
}

type paymentQuery struct {
	*daolib.BaseQuery
}

func (q *paymentQuery) SetBaseQuery(baseQuery *daolib.BaseQuery) {
	q.BaseQuery = baseQuery
}

func (q *paymentQuery) InsertPayment(payment entity.Payment) error {
	query := q.QueryBuilder().
		Insert(paymentTable).
		Columns(
			"id",
			"user_id",
			"subscription_id",
			"amount",
			"status",
			"payment_provider_id",
			"psp_token",
			"created_at",
			"updated_at",
		).
		Values(
			payment.ID,
			payment.UserID,
			payment.SubscriptionID,
			payment.Amount,
			payment.Status,
			payment.PaymentProviderID,
			payment.PspToken,
			payment.CreatedAt,
			payment.UpdatedAt,
		)

	_, err := q.Runner().Execx(q.Context(), query)
	if err != nil {
		return err
	}

	return nil
}

func (q *paymentQuery) GetPaymentByID(id uuid.UUID) (entity.Payment, error) {
	query := q.QueryBuilder().
		Select(
			"id",
			"user_id",
			"subscription_id",
			"amount",
			"status",
			"payment_provider_id",
			"psp_token",
			"created_at",
			"updated_at",
		).
		From(paymentTable).
		Where(sq.Eq{
			"id": id,
		})

	var payment entity.Payment
	if err := q.Runner().Getx(q.Context(), &payment, query); err != nil {
		return entity.Payment{}, err
	}

	return payment, nil
}

func (q *paymentQuery) UpdatePayment(payment entity.Payment) (int64, error) {
	query := q.QueryBuilder().
		Update(paymentTable).
		SetMap(sq.Eq{
			"amount":              payment.Amount,
			"status":              payment.Status,
			"payment_provider_id": payment.PaymentProviderID,
			"psp_token":           payment.PspToken,
			"updated_at":          payment.UpdatedAt,
		}).
		Where(sq.Eq{
			"id": payment.ID,
		})

	res, err := q.Runner().Execx(q.Context(), query)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func (q *paymentQuery) GetPaymentByProviderID(providerPaymentID string, paymentProviderID uuid.UUID) (entity.Payment, error) {
	query := q.QueryBuilder().
		Select(
			"id",
			"user_id",
			"subscription_id",
			"amount",
			"status",
			"payment_provider_id",
			"psp_token",
			"created_at",
			"updated_at",
		).
		From(paymentTable).
		Where(sq.Eq{
			"psp_token":           providerPaymentID,
			"payment_provider_id": paymentProviderID,
		})

	var payment entity.Payment
	if err := q.Runner().Getx(q.Context(), &payment, query); err != nil {
		return entity.Payment{}, err
	}

	return payment, nil
}

func (q *paymentQuery) CheckExistPaymentByStatusAndUserID(status domain.PaymentStatus, userID uuid.UUID) (bool, error) {
	query := q.QueryBuilder().
		Select("1").
		From(paymentTable).
		Where(sq.Eq{
			"status":  status,
			"user_id": userID,
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
