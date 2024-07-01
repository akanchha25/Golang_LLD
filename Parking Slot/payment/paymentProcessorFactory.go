package payment

import "sync"

var (
	paymentFactory *PaymentProcessorFactory
	once sync.Once
)

type PaymentProcessorFactory struct{}

func GetPaymentProcessorFactory() *PaymentProcessorFactory {
	once.Do(func() {
		paymentFactory = &PaymentProcessorFactory{}
	})
	return paymentFactory
}

func (p *PaymentProcessorFactory) GetPaymentProcessor(paymentMode PaymentMode) *PaymentProcessor {
	var paymentProcessor PaymentProcessor

	switch paymentMode {
	case CASH:
		paymentProcessor = &CashPaymentProcessor{}
	case CARD:
		paymentProcessor = &CardPaymentProcessor{}	
	}

	return &paymentProcessor
}