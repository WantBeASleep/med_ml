package yookassa_webhook

import (
	"context"

	"composition-api/internal/adapters"
	domain "composition-api/internal/domain/billing"
)

type Service interface {
	HandleYookassaWebhook(ctx context.Context, req domain.YookassaWebhookRequest) error
}

type service struct {
	adapters *adapters.Adapters
}

func New(adapters *adapters.Adapters) Service {
	return &service{
		adapters: adapters,
	}
}
