package repository

import (
	"context"

	"github.com/WantBeASleep/goooool/daolib"
	"github.com/jmoiron/sqlx"

	"billing/internal/repository/payment"
	"billing/internal/repository/payment_notification"
	"billing/internal/repository/payment_provider"
	"billing/internal/repository/subscription"
	"billing/internal/repository/tariff_plan"
)

type DAO interface {
	daolib.DAO
	NewTariffPlanQuery(ctx context.Context) tariff_plan.Repository
	NewSubscriptionQuery(ctx context.Context) subscription.Repository
	NewPaymentQuery(ctx context.Context) payment.Repository
	NewPaymentNotificationQuery(ctx context.Context) payment_notification.Repository
	NewPaymentProviderQuery(ctx context.Context) payment_provider.Repository
}

type dao struct {
	daolib.DAO
}

func NewRepository(psql *sqlx.DB) DAO {
	return &dao{DAO: daolib.NewDao(psql)}
}

func (d *dao) NewTariffPlanQuery(ctx context.Context) tariff_plan.Repository {
	tariffPlanQuery := tariff_plan.NewRepo()
	d.NewRepo(ctx, tariffPlanQuery)

	return tariffPlanQuery
}

func (d *dao) NewSubscriptionQuery(ctx context.Context) subscription.Repository {
	subscriptionQuery := subscription.NewRepo()
	d.NewRepo(ctx, subscriptionQuery)

	return subscriptionQuery
}

func (d *dao) NewPaymentQuery(ctx context.Context) payment.Repository {
	paymentQuery := payment.NewRepo()
	d.NewRepo(ctx, paymentQuery)

	return paymentQuery
}

func (d *dao) NewPaymentNotificationQuery(ctx context.Context) payment_notification.Repository {
	paymentNotificationQuery := payment_notification.NewRepo()
	d.NewRepo(ctx, paymentNotificationQuery)

	return paymentNotificationQuery
}

func (d *dao) NewPaymentProviderQuery(ctx context.Context) payment_provider.Repository {
	paymentProviderQuery := payment_provider.NewRepo()
	d.NewRepo(ctx, paymentProviderQuery)

	return paymentProviderQuery
}
