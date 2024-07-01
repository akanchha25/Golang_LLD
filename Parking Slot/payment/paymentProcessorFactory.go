package payment


// PaymentProcessorFactory provides methods to get different payment processors
type PaymentProcessorFactory struct{}

func NewPaymentProcessorFactory() *PaymentProcessorFactory {
	return &PaymentProcessorFactory{}
}

func (f *PaymentProcessorFactory) GetCardBasedPaymentProcessor(amount float64, cardDetails CardDetails) PaymentProcessor {
	return &CardPaymentProcessor{
		Amount:      amount,
		CardDetails: cardDetails,
	}
}

func (f *PaymentProcessorFactory) GetCashBasedPaymentProcessor(amount float64) PaymentProcessor {
	return &CashPaymentProcessor{
		Amount: amount,
	}
}