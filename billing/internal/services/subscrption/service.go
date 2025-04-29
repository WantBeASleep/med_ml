package subscrption

import (
	"context"
	"fmt"
	"time"

	"billing/internal/domain"
	"billing/internal/repository"
	"billing/internal/repository/entity"
	"billing/internal/services/payment"

	"github.com/google/uuid"
)

type Service interface {
	PurchaseSubscrption(ctx context.Context, tariffPlanID uuid.UUID, paymentProviderID uuid.UUID, userId uuid.UUID) (subscribtionID uuid.UUID, confirmationURL string, err error)
	SetSubscrptionStatus(ctx context.Context, subscriptionId uuid.UUID, status domain.SubscriptionStatus) error
	IsUserHasActiveSubscrption(ctx context.Context, userId uuid.UUID) (bool, error)
	GetUserActiveSubscrption(ctx context.Context, userId uuid.UUID) (*domain.Subscription, error)
	GetAllActiveSubscriptions(ctx context.Context) ([]domain.Subscription, error)
	CancelSubscrptions(ctx context.Context, subscriptionIDs []uuid.UUID) error
}

type service struct {
	dao        repository.DAO
	PaymentSrv payment.Service
}

func New(
	dao repository.DAO,
	PaymentSrv payment.Service,
) Service {
	return &service{
		PaymentSrv: PaymentSrv,
		dao:        dao,
	}
}

func (s *service) PurchaseSubscrption(ctx context.Context, tariffPlanID uuid.UUID, paymentProviderID uuid.UUID, userId uuid.UUID) (subscribtionID uuid.UUID, confirmationURL string, err error) {
	// проверяем тариф план
	tariffPlanDB, err := s.dao.NewTariffPlanQuery(ctx).GetTariffPlanByID(tariffPlanID)
	if err != nil {
		return uuid.Nil, "", fmt.Errorf("get tariff plan by id: %w", err)
	}
	tariffPlan := tariffPlanDB.ToDomain()
	// не даем создать пользователю больше одной незавершенной подписки
	ctx, err = s.dao.BeginTx(ctx)
	if err != nil {
		return uuid.Nil, "", fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer func() {
		if err != nil {
			_ = s.dao.RollbackTx(ctx)
		}
	}()
	userSubscribtionsDB, err := s.dao.NewSubscriptionQuery(ctx).GetSubscriptionsByUserID(userId)
	if err != nil {
		return uuid.Nil, "", fmt.Errorf("get subscribtion by userID: %w", err)
	}
	for _, subscribtionDB := range userSubscribtionsDB {
		subscribtion := subscribtionDB.ToDomain()
		if subscribtion.Status != domain.SubCancelled {
			return uuid.Nil, "", fmt.Errorf("already have active or pending_payment subscribtion")
		}
	}
	// создаем подписку
	newSubscribtion := domain.Subscription{
		ID:           uuid.New(),
		UserID:       userId,
		TariffPlanID: tariffPlanID,
		StartDate:    time.Now(),
		EndDate:      time.Now().Add(tariffPlan.Duration),
		Status:       domain.SubPendingPayment,
	}

	err = s.dao.NewSubscriptionQuery(ctx).InsertSubscription(entity.Subscription{}.FromDomain(newSubscribtion))
	if err != nil {
		return uuid.Nil, "", fmt.Errorf("insert subscription: %w", err)
	}
	// создаем платеж
	_, confirmationURL, err = s.PaymentSrv.CreatePayment(ctx, userId, newSubscribtion.ID, tariffPlan.Price, tariffPlan.Description, paymentProviderID)
	if err != nil {
		return uuid.Nil, "", fmt.Errorf("create payment: %w", err)
	}

	if err := s.dao.CommitTx(ctx); err != nil {
		return uuid.Nil, "", fmt.Errorf("commit transaction: %w", err)
	}
	return newSubscribtion.ID, confirmationURL, nil
}

func (s *service) SetSubscrptionStatus(ctx context.Context, subscriptionId uuid.UUID, status domain.SubscriptionStatus) error {
	subscriptionDB, err := s.dao.NewSubscriptionQuery(ctx).GetSubscriptionByID(subscriptionId)
	if err != nil {
		return fmt.Errorf("get subscription by id: %w", err)
	}
	subscription := subscriptionDB.ToDomain()
	subscription.Status = status
	if _, err = s.dao.NewSubscriptionQuery(ctx).UpdateSubscription(entity.Subscription{}.FromDomain(subscription)); err != nil {
		return fmt.Errorf("update subscription: %w", err)
	}
	return nil
}

func (s *service) IsUserHasActiveSubscrption(ctx context.Context, userId uuid.UUID) (bool, error) {
	exists, err := s.dao.NewSubscriptionQuery(ctx).CheckExistSubscrptionByStatusAndUserID(string(domain.SubActive), userId)
	if err != nil {
		return false, fmt.Errorf("check exist subscrption: %w", err)
	}
	return exists, nil
}

func (s *service) GetUserActiveSubscrption(ctx context.Context, userId uuid.UUID) (*domain.Subscription, error) {
	subscriptionsDB, err := s.dao.NewSubscriptionQuery(ctx).GetSubscrptionsByStatusAndUserID(string(domain.SubActive), userId)
	if err != nil {
		return nil, fmt.Errorf("get subscrptions by status and user_id: %w", err)
	}
	if len(subscriptionsDB) == 0 {
		return nil, nil
	}
	subscription := subscriptionsDB[0].ToDomain()
	return &subscription, nil
}

func (s *service) GetAllActiveSubscriptions(ctx context.Context) ([]domain.Subscription, error) {
	activeSubscriptionDB, err := s.dao.NewSubscriptionQuery(ctx).GetAllActiveSubscriptions()
	if err != nil {
		return nil, fmt.Errorf("get active subscriptions: %w", err)
	}

	var activeSubscriptions []domain.Subscription
	for _, subscriptionDB := range activeSubscriptionDB {
		subscription := subscriptionDB.ToDomain()
		activeSubscriptions = append(activeSubscriptions, subscription)
	}

	return activeSubscriptions, nil
}

func (s *service) CancelSubscrptions(ctx context.Context, subscriptionIDs []uuid.UUID) error {
	err := s.dao.NewSubscriptionQuery(ctx).SetSubscriptionsStatusBatch(subscriptionIDs, domain.SubCancelled)
	if err != nil {
		return fmt.Errorf("cansel subscrptions: %w", err)
	}
	return nil
}
