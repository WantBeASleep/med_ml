package notification_manager

import (
	"context"
	"fmt"
	"time"

	"billing/internal/repository/entity"

	"github.com/google/uuid"

	"billing/internal/domain"
	"billing/internal/repository"
	"billing/internal/services/payment"
	"billing/internal/services/subscrption"
	"billing/internal/services/yookassa"
)

type (
	UniversalNotification struct {
		ProviderPaymentID string
		Event             string
		ProviderName      string
		ReceivedAt        time.Time
		ExtraInfo         map[string]interface{}
	}

	SuccessfulNotification struct {
		UniversalNotification
	}

	CanceledNotification struct {
		UniversalNotification
		Reason string
	}

	WaitingForCaptureNotification struct {
		UniversalNotification
	}
)

type Service interface {
	HandlSuccessfulNotification(ctx context.Context, notification SuccessfulNotification) error
	HandlCanceledNotification(ctx context.Context, notification CanceledNotification) error
	HandlWaitingForCaptureNotification(ctx context.Context, notification WaitingForCaptureNotification) error
}

type service struct {
	dao             repository.DAO
	SubscriptionSrv subscrption.Service
	PaymentSrv      payment.Service
	YookassaSrv     yookassa.Service
}

func New(
	dao repository.DAO,
	SubscriptionSrv subscrption.Service,
	PaymentSrv payment.Service,
	yookassaSrv yookassa.Service,
) Service {
	return &service{
		dao:             dao,
		SubscriptionSrv: SubscriptionSrv,
		PaymentSrv:      PaymentSrv,
		YookassaSrv:     yookassaSrv,
	}
}

func (s *service) HandlSuccessfulNotification(ctx context.Context, notification SuccessfulNotification) error {
	// на всякий случай еще проверить, что все подписки и платежи перешли в нужный статус
	// а так все должно произойти при capture
	ctx, err := s.dao.BeginTx(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer func() {
		if err != nil {
			_ = s.dao.RollbackTx(ctx)
		}
	}()
	internalPayment, err := s.PaymentSrv.GetPaymentByProviderName(ctx, notification.ProviderPaymentID, notification.ProviderName)
	if err != nil {
		return fmt.Errorf("failed to get payment by provider name: %w", err)
	}

	// Проверяем, что уведомление настоящее
	isValid, err := s.PaymentSrv.IsActualProviderPaymentStatus(ctx, notification.ProviderPaymentID, internalPayment.PaymentProviderID, domain.PayCompleted)
	if err != nil {
		return fmt.Errorf("failed to verify payment status: %w", err)
	}
	if !isValid {
		return fmt.Errorf("payment status does not match notification")
	}

	// Проверяем текущий статус платежа
	if internalPayment.Status != domain.PayWaitingForCapture {
		return fmt.Errorf("payment is not in waiting for capture status")
	}

	if err := s.PaymentSrv.SetPaymentStatus(ctx, internalPayment.ID, domain.PayCompleted); err != nil {
		return fmt.Errorf("failed to set payment status to completed: %w", err)
	}

	if err := s.SubscriptionSrv.SetSubscrptionStatus(ctx, internalPayment.SubscriptionID, domain.SubActive); err != nil {
		return fmt.Errorf("failed to set subscription status to active: %w", err)
	}

	newNotification := domain.PaymentNotification{
		ID:                uuid.New(),
		ProviderPaymentID: notification.ProviderPaymentID,
		Event:             notification.Event,
		PaymentProviderID: internalPayment.PaymentProviderID,
		ReceivedAt:        notification.ReceivedAt,
		NotificationData:  notification.ExtraInfo,
		IsValid:           true,
	}

	err = s.dao.NewPaymentNotificationQuery(ctx).InsertPaymentNotification(entity.PaymentNotification{}.FromDomain(newNotification))
	if err != nil {
		return fmt.Errorf("failed to insert payment notification: %w", err)
	}

	if err = s.dao.CommitTx(ctx); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (s *service) HandlCanceledNotification(ctx context.Context, notification CanceledNotification) error {
	// на всякий случай еще проверить, что все подписки и платежи перешли в нужный статус
	// а так все должно произойти при canceled

	ctx, err := s.dao.BeginTx(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer func() {
		if err != nil {
			_ = s.dao.RollbackTx(ctx)
		}
	}()

	internalPayment, err := s.PaymentSrv.GetPaymentByProviderName(ctx, notification.ProviderPaymentID, notification.ProviderName)
	if err != nil {
		return fmt.Errorf("failed to get payment by provider name: %w", err)
	}

	// Проверяем, что уведомление настоящее
	isValid, err := s.PaymentSrv.IsActualProviderPaymentStatus(ctx, notification.ProviderPaymentID, internalPayment.PaymentProviderID, domain.PayCancelled)
	if err != nil {
		return fmt.Errorf("failed to verify payment status: %w", err)
	}
	if !isValid {
		return fmt.Errorf("payment status does not match notification")
	}

	// Проверяем текущий статус платежа
	if internalPayment.Status != domain.PayWaitingForCancel {
		return fmt.Errorf("payment is not in waiting for cancel status")
	}

	if err := s.PaymentSrv.SetPaymentStatus(ctx, internalPayment.ID, domain.PayCancelled); err != nil {
		return fmt.Errorf("failed to set payment status to cancelled: %w", err)
	}

	if err := s.SubscriptionSrv.SetSubscrptionStatus(ctx, internalPayment.SubscriptionID, domain.SubCancelled); err != nil {
		return fmt.Errorf("failed to set subscription status to cancelled: %w", err)
	}

	newNotification := domain.PaymentNotification{
		ID:                uuid.New(),
		ProviderPaymentID: notification.ProviderPaymentID,
		Event:             notification.Event,
		PaymentProviderID: internalPayment.PaymentProviderID,
		ReceivedAt:        notification.ReceivedAt,
		NotificationData:  notification.ExtraInfo,
		IsValid:           true,
	}

	err = s.dao.NewPaymentNotificationQuery(ctx).InsertPaymentNotification(entity.PaymentNotification{}.FromDomain(newNotification))
	if err != nil {
		return fmt.Errorf("failed to insert payment notification: %w", err)
	}

	if err = s.dao.CommitTx(ctx); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (s *service) HandlWaitingForCaptureNotification(ctx context.Context, notification WaitingForCaptureNotification) error {
	// проверяем нет ли у юзера активной подписки или платежей в статусе WaitingForCapture
	// если все ок, то делаем capture
	// переводим в нужные статусы у платежа и подписки
	// если не ок, то делаем отмену платежа
	// и переводим статусы подписки и платежа в отмену

	ctx, err := s.dao.BeginTx(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer func() {
		if err != nil {
			_ = s.dao.RollbackTx(ctx)
		}
	}()

	internalPayment, err := s.PaymentSrv.GetPaymentByProviderName(ctx, notification.ProviderPaymentID, notification.ProviderName)
	if err != nil {
		return fmt.Errorf("failed to get payment by provider name: %w", err)
	}

	// Проверяем, что уведомление настоящее
	isValid, err := s.PaymentSrv.IsActualProviderPaymentStatus(ctx, notification.ProviderPaymentID, internalPayment.PaymentProviderID, domain.PayWaitingForCapture)
	if err != nil {
		return fmt.Errorf("failed to verify payment status: %w", err)
	}
	if !isValid {
		return fmt.Errorf("payment status does not match notification")
	}

	if internalPayment.Status != domain.PayPending {
		return fmt.Errorf("payment is not in waiting for capture status")
	}

	// Проверяем, есть ли у пользователя активная подписка
	hasActive, err := s.SubscriptionSrv.IsUserHasActiveSubscrption(ctx, internalPayment.UserID)
	if err != nil {
		return fmt.Errorf("failed to check if user has active subscription: %w", err)
	}

	// Проверяем, есть ли у пользователя платежи в статусе WaitingForCapture
	hasWaitingForCapture, err := s.PaymentSrv.HasPaymentWithStatus(ctx, internalPayment.UserID, domain.PayWaitingForCapture)
	if err != nil {
		return fmt.Errorf("failed to check if user has payments waiting for capture: %w", err)
	}

	if hasActive || hasWaitingForCapture {
		// Если есть активная подписка или платежи в статусе WaitingForCapture, отменяем платеж
		if err := s.PaymentSrv.CancelPaymentByProvider(ctx, internalPayment.ID); err != nil {
			return fmt.Errorf("failed to cancel payment: %w", err)
		}
	} else {
		// Если нет активной подписки и платежей в статусе WaitingForCapture, захватываем платеж
		if err := s.PaymentSrv.CapturePaymentByProvider(ctx, internalPayment.ID); err != nil {
			return fmt.Errorf("failed to capture payment: %w", err)
		}
	}

	newNotification := domain.PaymentNotification{
		ID:                uuid.New(),
		ProviderPaymentID: notification.ProviderPaymentID,
		Event:             notification.Event,
		PaymentProviderID: internalPayment.PaymentProviderID,
		ReceivedAt:        notification.ReceivedAt,
		NotificationData:  notification.ExtraInfo,
		IsValid:           true,
	}

	err = s.dao.NewPaymentNotificationQuery(ctx).InsertPaymentNotification(entity.PaymentNotification{}.FromDomain(newNotification))
	if err != nil {
		return fmt.Errorf("failed to insert payment notification: %w", err)
	}

	if err = s.dao.CommitTx(ctx); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
