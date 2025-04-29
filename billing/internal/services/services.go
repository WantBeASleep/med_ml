package services

import (
	"billing/internal/config"
	"billing/internal/repository"
	"billing/internal/services/notification_manager"
	"billing/internal/services/payment"
	"billing/internal/services/payment_provider"
	"billing/internal/services/subscrption"
	"billing/internal/services/tariff_plan"
	"billing/internal/services/yookassa"
)

type Services struct {
	NotificationManager notification_manager.Service
	Payment             payment.Service
	Subscription        subscrption.Service
	TariffPlan          tariff_plan.Service
	Yookassa            yookassa.Service
	PaymentProvider     payment_provider.Service
}

func New(
	dao repository.DAO,
	cfg *config.Config,
) *Services {
	yookassaService := yookassa.New(cfg.Yookassa.AccountID, cfg.Yookassa.SecretKey, cfg.Yookassa.ReturnURL)
	paymentService := payment.New(dao, yookassaService)
	subscriptionService := subscrption.New(dao, paymentService)
	tariffPlanService := tariff_plan.New(dao)
	notificationManagerService := notification_manager.New(dao, subscriptionService, paymentService, yookassaService)
	paymentProviderService := payment_provider.New(dao)

	return &Services{
		NotificationManager: notificationManagerService,
		Payment:             paymentService,
		Subscription:        subscriptionService,
		TariffPlan:          tariffPlanService,
		Yookassa:            yookassaService,
		PaymentProvider:     paymentProviderService,
	}
}
