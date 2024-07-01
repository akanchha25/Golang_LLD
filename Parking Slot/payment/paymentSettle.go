package payment

import (
	"errors"
	"strconv"
)

type PaymentSettle struct{
	paymentMode PaymentMode
	paymentDetails map[string]string
	amount float64
}

func NewPaymentSettle(paymentMode PaymentMode, paymentDetails map[string]string, amount float64) *PaymentSettle {
	return &PaymentSettle{
		paymentMode:    paymentMode,
		paymentDetails: paymentDetails,
		amount:         amount,
	}
}

func (s *PaymentSettle) GetPaymentSettle(paymentMode PaymentMode, paymentDetails map[string]string, amount float64) (PaymentProcessor, error) {
	paymentProcessorFactory := NewPaymentProcessorFactory()

	var paymentProcessor PaymentProcessor
	switch paymentMode {
	case CARD:
		cardDetails := CardDetails{
			NameOnCard: paymentDetails["NAME_ON_CARD"],
			Pin:        parsePin(paymentDetails["PIN"]),
			CardNumber: paymentDetails["CARD_NUMBER"],
		}
		paymentProcessor = paymentProcessorFactory.GetCardBasedPaymentProcessor(amount, cardDetails)
	case CASH:
		paymentProcessor = paymentProcessorFactory.GetCashBasedPaymentProcessor(amount)
	default:
		return nil, errors.New("unsupported payment mode") // Correctly return nil for the first value
	}

	return paymentProcessor, nil
}

func parsePin(pinStr string) int {
	pin, _ := strconv.Atoi(pinStr)
	return pin
}