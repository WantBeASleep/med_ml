package subscription_checker

import (
	"context"
	"log/slog"
	"time"

	"billing/internal/services/subscrption"

	"github.com/google/uuid"
)

type Service struct {
	SubscriptionSrv subscrption.Service
}

func New(SubscriptionSrv subscrption.Service) *Service {
	return &Service{
		SubscriptionSrv: SubscriptionSrv,
	}
}

func (s *Service) Start(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			s.checkAndCancelExpiredSubscriptions(ctx)
		case <-ctx.Done():
			return
		}
	}
}

func (s *Service) checkAndCancelExpiredSubscriptions(ctx context.Context) {
	slog.Info("Starting check for expired subscriptions")
	subscriptions, err := s.SubscriptionSrv.GetAllActiveSubscriptions(ctx)
	if err != nil {
		slog.Error("Failed to get active subscriptions", "err", err)
		return
	}

	var expiredIDs []uuid.UUID
	for _, sub := range subscriptions {
		slog.Info("Checking subscription", "id", sub.ID, "end_date", sub.EndDate, "status", sub.Status, "time Now", time.Now(), "NOW UNIX", time.Now().Unix(), "sub UNIX", sub.EndDate.Unix())
		if time.Now().Unix() > sub.EndDate.Unix() {
			slog.Info("Subscription expired", "id", sub.ID)
			expiredIDs = append(expiredIDs, sub.ID)
		} else {
			slog.Info("Subscription still active", "id", sub.ID)
		}
	}

	if len(expiredIDs) > 0 {
		slog.Info("Cancelling expired subscriptions", "count", len(expiredIDs))
		err := s.SubscriptionSrv.CancelSubscrptions(ctx, expiredIDs)
		if err != nil {
			slog.Error("Failed to set subscriptions status", "err", err)
		} else {
			slog.Info("Subscriptions status updated", "count", len(expiredIDs))
		}
	} else {
		slog.Info("No expired subscriptions found")
	}
}
