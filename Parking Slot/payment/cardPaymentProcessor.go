package payment


// CardPaymentProcessor is a struct that implements the PaymentProcessor interface.
type CardPaymentProcessor struct {
    Amount      float64
    CardDetails CardDetails
}

// NewCardPaymentProcessor is a constructor for CardPaymentProcessor.
func NewCardPaymentProcessor(amount float64, cardDetails CardDetails) *CardPaymentProcessor {
    return &CardPaymentProcessor{
        Amount:      amount,
        CardDetails: cardDetails,
    }
}

// ExecutePayment processes the payment and returns a boolean indicating success or failure.
func (cpp *CardPaymentProcessor) ExecutePayment() bool {
    // Payment processing logic goes here
    return false
}

// GetAmount returns the amount to be paid.
func (cpp *CardPaymentProcessor) GetAmount() float64 {
    return cpp.Amount
}