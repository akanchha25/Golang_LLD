package payment

// PaymentProcessor is an interface for processing payments.
type PaymentProcessor interface {
	GetAmount() float64
	ExecutePayment() bool
}
