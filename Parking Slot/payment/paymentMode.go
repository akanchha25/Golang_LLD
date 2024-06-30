package payment

type PaymentMode int

const (
    // Define your vehicle types here.
    CASH PaymentMode = iota
    CARD
)
