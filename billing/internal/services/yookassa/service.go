package yookassa

import (
	"context"
	"fmt"

	"github.com/rvinnie/yookassa-sdk-go/yookassa"
	yoocommon "github.com/rvinnie/yookassa-sdk-go/yookassa/common"
	yoopayment "github.com/rvinnie/yookassa-sdk-go/yookassa/payment"
	"github.com/shopspring/decimal"
)

const CurrencyRUB = "RUB"

type Service interface {
	CreateRedirectPayment(ctx context.Context, value decimal.Decimal, description string) (paymentToken string, confirmationURL string, err error)
	CapturePayment(ctx context.Context, paymentId string) error
	CancelPayment(ctx context.Context, paymentId string) error
	IsPaymentPending(ctx context.Context, paymentId string) (bool, error)
	IsPaymentWaitingForCapture(ctx context.Context, paymentId string) (bool, error)
	IsPaymentSucceeded(ctx context.Context, paymentId string) (bool, error)
	IsPaymentCanceled(ctx context.Context, paymentId string) (bool, error)
}

type service struct {
	Yooclient *yookassa.Client
	ReturnURL string
}

func New(
	accountId string,
	secretKey string,
	returnURL string,
) Service {
	return &service{
		Yooclient: yookassa.NewClient(accountId, secretKey),
		ReturnURL: returnURL,
	}
}

func (s *service) CapturePayment(ctx context.Context, paymentId string) error {
	paymentHandler := yookassa.NewPaymentHandler(s.Yooclient)
	payment, err := paymentHandler.FindPayment(paymentId)
	if err != nil {
		return fmt.Errorf("failed to find yookassa payment for capture: %w", err)
	}
	_, err = paymentHandler.CapturePayment(payment)
	if err != nil {
		return fmt.Errorf("failed to capture yookassa payment: %w", err)
	}
	return nil
}

func (s *service) CreateRedirectPayment(ctx context.Context, value decimal.Decimal, description string) (paymentToken string, confirmationURL string, err error) {
	paymentHandler := yookassa.NewPaymentHandler(s.Yooclient)
	payment, err := paymentHandler.CreatePayment(&yoopayment.Payment{
		Amount: &yoocommon.Amount{
			Value:    value.StringFixed(2),
			Currency: CurrencyRUB,
		},
		Confirmation: yoopayment.Redirect{
			Type:      yoopayment.TypeRedirect,
			ReturnURL: s.ReturnURL,
		},
		Description: description,
	})
	if err != nil {
		return "", "", fmt.Errorf("failed to create yookassa payment: %w", err)
	}

	confirmationURL, err = paymentHandler.ParsePaymentLink(payment)
	if err != nil {
		return "", "", fmt.Errorf("failed to parse yookassa payment link: %w", err)
	}

	return payment.ID, confirmationURL, nil
}

func (s *service) CancelPayment(ctx context.Context, paymentId string) error {
	paymentHandler := yookassa.NewPaymentHandler(s.Yooclient)
	_, err := paymentHandler.FindPayment(paymentId)
	if err != nil {
		return fmt.Errorf("failed to find yookassa payment for cancel: %w", err)
	}
	_, err = paymentHandler.CancelPayment(paymentId)
	if err != nil {
		return fmt.Errorf("failed to cancel yookassa payment: %w", err)
	}
	return nil
}

func (s *service) IsPaymentPending(ctx context.Context, paymentId string) (bool, error) {
	return s.checkPaymentStatus(ctx, paymentId, yoopayment.Pending)
}

func (s *service) IsPaymentWaitingForCapture(ctx context.Context, paymentId string) (bool, error) {
	return s.checkPaymentStatus(ctx, paymentId, yoopayment.WaitingForCapture)
}

func (s *service) IsPaymentSucceeded(ctx context.Context, paymentId string) (bool, error) {
	return s.checkPaymentStatus(ctx, paymentId, yoopayment.Succeeded)
}

func (s *service) IsPaymentCanceled(ctx context.Context, paymentId string) (bool, error) {
	return s.checkPaymentStatus(ctx, paymentId, yoopayment.Canceled)
}

func (s *service) checkPaymentStatus(ctx context.Context, paymentId string, expectedStatus yoopayment.Status) (bool, error) {
	paymentHandler := yookassa.NewPaymentHandler(s.Yooclient)
	payment, err := paymentHandler.FindPayment(paymentId)
	if err != nil {
		return false, fmt.Errorf("failed to find yookassa payment: %w", err)
	}

	if payment.Status == expectedStatus {
		return true, nil
	}

	return false, nil
}
