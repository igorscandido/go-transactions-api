package client

import (
	"context"
	"fmt"
	"math"

	"github.com/igorscandido/go-transactions-api/internal/domain"
	"github.com/igorscandido/go-transactions-api/internal/ports"
	"github.com/igorscandido/go-transactions-api/pkg/configs"
	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/paymentintent"
	//"github.com/stripe/stripe-go/v81/paymentmethod"
)

type stripeClient struct {
}

func NewStripePaymentClient(configs *configs.Configs) ports.StripeClient {
	stripe.Key = configs.ExternalKeys.Stripe
	return &stripeClient{}
}

func (s *stripeClient) CreatePaymentMethod(ctx context.Context, cardDetails *domain.CardDetails) (*string, error) {
	// cardParams := &stripe.PaymentMethodParams{
	// 	Type: stripe.String(string(stripe.PaymentMethodTypeCard)),
	// 	Card: &stripe.PaymentMethodCardParams{
	// 		Number:   stripe.String(cardDetails.Number),
	// 		CVC:      stripe.String(cardDetails.CVV),
	// 		ExpMonth: stripe.Int64(int64(cardDetails.ExpiryNumber)),
	// 		ExpYear:  stripe.Int64(int64(cardDetails.ExpiryYear)),
	// 	},
	// }

	// paymentMethod, err := paymentmethod.New(cardParams)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to create payment method: %v", err)
	// }

	// return &paymentMethod.ID, nil

	return &[]string{"pm_card_visa"}[0], nil
}

func (s *stripeClient) ProcessPayment(ctx context.Context, payment *domain.Payment, paymentMethodID *string) (*string, error) {
	piParams := &stripe.PaymentIntentParams{
		Amount:        stripe.Int64(convertFloatToInt64Decimal(*payment.GatewayTransactionatedAmount)),
		Currency:      stripe.String(payment.GatewayTransactionatedCurrency),
		PaymentMethod: stripe.String(*paymentMethodID),
		Confirm:       stripe.Bool(true),
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			AllowRedirects: stripe.String("never"),
			Enabled:        stripe.Bool(true),
		},
	}

	paymentIntent, err := paymentintent.New(piParams)
	if err != nil {
		return nil, fmt.Errorf("failed to process payment: %v", err)
	}

	return &paymentIntent.ID, nil
}

func (s *stripeClient) GetPaymentStatus(ctx context.Context, paymentID *string) (*string, error) {
	paymentIntent, err := paymentintent.Get(*paymentID, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve payment status: %w", err)
	}

	paymentStatus := string(paymentIntent.Status)
	return &paymentStatus, nil
}

func convertFloatToInt64Decimal(amount float64) int64 {
	scaledAmount := amount * 100
	return int64(math.Trunc(scaledAmount))
}

func convertInt64DecimalToFloat(amountInCents int64) float64 {
	return float64(amountInCents) / 100
}
