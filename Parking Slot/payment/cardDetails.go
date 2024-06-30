package payment

// CardDetails represents the details of a card.
type CardDetails struct {
    NameOnCard  string
    Pin         int
    CardNumber  string
}

// NewCardDetails is a constructor for CardDetails.
func NewCardDetails(nameOnCard string, pin int, cardNumber string) *CardDetails {
    return &CardDetails{
        NameOnCard: nameOnCard,
        Pin:        pin,
        CardNumber: cardNumber,
    }
}
