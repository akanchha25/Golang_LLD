package payment

type Payment interface {
	ProcessPayment(amount float64) error
}
