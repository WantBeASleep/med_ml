package payment

import (
	"context"
	"fmt"
	"time"

	"billing/internal/domain"
	"billing/internal/repository"
	"billing/internal/repository/entity"
	"billing/internal/services/yookassa"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

const (
	ProviderYookassa = "Yookassa"
)

type Service interface {
	CreatePayment(ctx context.Context, userID uuid.UUID, subscribtionID uuid.UUID, amount decimal.Decimal, description string, PaymentProviderID uuid.UUID) (paymentID uuid.UUID, confirmationURL string, err error)
	SetPaymentStatus(ctx context.Context, paymentId uuid.UUID, status domain.PaymentStatus) error
	CancelPaymentByProvider(ctx context.Context, paymentId uuid.UUID) error
	CapturePaymentByProvider(ctx context.Context, paymentId uuid.UUID) error
	GetPaymentByID(ctx context.Context, paymentID uuid.UUID) (domain.Payment, error)
	GetPaymentByProviderID(ctx context.Context, providerPaymentID string, paymentProviderID uuid.UUID) (domain.Payment, error)
	GetPaymentByProviderName(ctx context.Context, providerPaymentID string, providerName string) (domain.Payment, error)
	HasPaymentWithStatus(ctx context.Context, userID uuid.UUID, status domain.PaymentStatus) (bool, error)
	IsActualProviderPaymentStatus(ctx context.Context, providerPaymentID string, paymentProviderID uuid.UUID, expectedStatus domain.PaymentStatus) (bool, error)
}

type service struct {
	dao         repository.DAO
	YookassaSrv yookassa.Service
}

func New(
	dao repository.DAO,
	yookassaSrv yookassa.Service,
) Service {
	return &service{
		dao:         dao,
		YookassaSrv: yookassaSrv,
	}
}

func (s *service) CreatePayment(ctx context.Context, userID uuid.UUID, subscriptionID uuid.UUID, amount decimal.Decimal, description string, paymentProviderID uuid.UUID) (paymentID uuid.UUID, confirmationURL string, err error) {
	// Получаем информацию о провайдере из базы данных
	providerDB, err := s.dao.NewPaymentProviderQuery(ctx).GetPaymentProviderByID(paymentProviderID)
	if err != nil {
		return uuid.Nil, "", fmt.Errorf("failed to get payment provider: %w", err)
	}
	provider := providerDB.ToDomain()

	if !provider.IsActive {
		return uuid.Nil, "", fmt.Errorf("payment provider is not active: %s", provider.Name)
	}

	var pspToken string

	switch provider.Name {
	case ProviderYookassa:
		pspToken, confirmationURL, err = s.YookassaSrv.CreateRedirectPayment(ctx, amount, description)
		if err != nil {
			return uuid.Nil, "", fmt.Errorf("failed to create payment with Yookassa: %w", err)
		}
	default:
		return uuid.Nil, "", fmt.Errorf("unsupported payment provider: %s", provider.Name)
	}

	// Создаем запись о платеже в базе данных
	newPayment := domain.Payment{
		ID:                uuid.New(),
		UserID:            userID,
		SubscriptionID:    subscriptionID,
		Amount:            amount,
		Status:            domain.PayPending,
		PaymentProviderID: paymentProviderID,
		PspToken:          pspToken,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}

	err = s.dao.NewPaymentQuery(ctx).InsertPayment(entity.Payment{}.FromDomain(newPayment))
	if err != nil {
		return uuid.Nil, "", fmt.Errorf("failed to insert payment: %w", err)
	}

	return newPayment.ID, confirmationURL, nil
}

func (s *service) SetPaymentStatus(ctx context.Context, paymentId uuid.UUID, status domain.PaymentStatus) error {
	paymentDB, err := s.dao.NewPaymentQuery(ctx).GetPaymentByID(paymentId)
	if err != nil {
		return fmt.Errorf("get payment by id: %w", err)
	}

	payment := paymentDB.ToDomain()
	payment.Status = status
	payment.UpdatedAt = time.Now()

	if _, err = s.dao.NewPaymentQuery(ctx).UpdatePayment(entity.Payment{}.FromDomain(payment)); err != nil {
		return fmt.Errorf("update payment: %w", err)
	}

	return nil
}

func (s *service) CancelPaymentByProvider(ctx context.Context, paymentId uuid.UUID) error {
	paymentDB, err := s.dao.NewPaymentQuery(ctx).GetPaymentByID(paymentId)
	if err != nil {
		return fmt.Errorf("failed to get payment by id: %w", err)
	}

	provider, err := s.dao.NewPaymentProviderQuery(ctx).GetPaymentProviderByID(paymentDB.PaymentProviderID)
	if err != nil {
		return fmt.Errorf("failed to get payment provider: %w", err)
	}

	switch provider.Name {
	case ProviderYookassa:
		err = s.YookassaSrv.CancelPayment(ctx, paymentDB.PspToken)
		if err != nil {
			return fmt.Errorf("failed to cancel payment in Yookassa: %w", err)
		}
	default:
		return fmt.Errorf("unsupported payment provider ID: %s", paymentDB.PaymentProviderID)
	}

	if err := s.SetPaymentStatus(ctx, paymentId, domain.PayWaitingForCancel); err != nil {
		return fmt.Errorf("failed to set payment status to cancelled: %w", err)
	}

	return nil
}

func (s *service) CapturePaymentByProvider(ctx context.Context, paymentId uuid.UUID) error {
	paymentDB, err := s.dao.NewPaymentQuery(ctx).GetPaymentByID(paymentId)
	if err != nil {
		return fmt.Errorf("failed to get payment by id: %w", err)
	}

	provider, err := s.dao.NewPaymentProviderQuery(ctx).GetPaymentProviderByID(paymentDB.PaymentProviderID)
	if err != nil {
		return fmt.Errorf("failed to get payment provider: %w", err)
	}

	switch provider.Name {
	case ProviderYookassa:
		err = s.YookassaSrv.CapturePayment(ctx, paymentDB.PspToken)
		if err != nil {
			return fmt.Errorf("failed to capture payment in Yookassa: %w", err)
		}
	default:
		return fmt.Errorf("unsupported payment provider ID: %s", paymentDB.PaymentProviderID)
	}

	if err := s.SetPaymentStatus(ctx, paymentId, domain.PayWaitingForCapture); err != nil {
		return fmt.Errorf("failed to set payment status to completed: %w", err)
	}

	return nil
}

func (s *service) GetPaymentByID(ctx context.Context, paymentID uuid.UUID) (domain.Payment, error) {
	paymentDB, err := s.dao.NewPaymentQuery(ctx).GetPaymentByID(paymentID)
	if err != nil {
		return domain.Payment{}, fmt.Errorf("failed to get payment by id: %w", err)
	}

	payment := paymentDB.ToDomain()

	return payment, nil
}

func (s *service) GetPaymentByProviderID(ctx context.Context, providerPaymentID string, paymentProviderID uuid.UUID) (domain.Payment, error) {
	paymentDB, err := s.dao.NewPaymentQuery(ctx).GetPaymentByProviderID(providerPaymentID, paymentProviderID)
	if err != nil {
		return domain.Payment{}, fmt.Errorf("failed to get payment by provider ID: %w", err)
	}

	return paymentDB.ToDomain(), nil
}

func (s *service) GetPaymentByProviderName(ctx context.Context, providerPaymentID string, providerName string) (domain.Payment, error) {
	providerDB, err := s.dao.NewPaymentProviderQuery(ctx).GetPaymentProviderByName(providerName)
	if err != nil {
		return domain.Payment{}, fmt.Errorf("failed to get payment provider by name: %w", err)
	}
	providerID := providerDB.ID

	return s.GetPaymentByProviderID(ctx, providerPaymentID, providerID)
}

func (s *service) HasPaymentWithStatus(ctx context.Context, userID uuid.UUID, status domain.PaymentStatus) (bool, error) {
	exists, err := s.dao.NewPaymentQuery(ctx).CheckExistPaymentByStatusAndUserID(status, userID)
	if err != nil {
		return false, fmt.Errorf("failed to check existence of payments with status: %w", err)
	}

	return exists, nil
}

func (s *service) IsActualProviderPaymentStatus(ctx context.Context, providerPaymentID string, paymentProviderID uuid.UUID, expectedStatus domain.PaymentStatus) (bool, error) {
	provider, err := s.dao.NewPaymentProviderQuery(ctx).GetPaymentProviderByID(paymentProviderID)
	if err != nil {
		return false, fmt.Errorf("failed to get payment provider: %w", err)
	}

	switch provider.Name {
	case ProviderYookassa:
		return s.checkYookassaPaymentStatus(ctx, providerPaymentID, expectedStatus)
	default:
		return false, fmt.Errorf("unsupported payment provider: %s", provider.Name)
	}
}

func (s *service) checkYookassaPaymentStatus(ctx context.Context, providerPaymentID string, expectedStatus domain.PaymentStatus) (bool, error) {
	switch expectedStatus {
	case domain.PayPending:
		return s.YookassaSrv.IsPaymentPending(ctx, providerPaymentID)
	case domain.PayWaitingForCapture:
		return s.YookassaSrv.IsPaymentWaitingForCapture(ctx, providerPaymentID)
	case domain.PayCompleted:
		return s.YookassaSrv.IsPaymentSucceeded(ctx, providerPaymentID)
	case domain.PayCancelled:
		return s.YookassaSrv.IsPaymentCanceled(ctx, providerPaymentID)
	default:
		return false, fmt.Errorf("unsupported payment status: %s", expectedStatus)
	}
}
