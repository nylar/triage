package models

// Ticket holds data on a support ticket
type Ticket struct {
	Message string
}

// NewTicket instantiates a new ticket
func NewTicket(msg string) *Ticket {
	return &Ticket{
		Message: msg,
	}
}

func (t *Ticket) String() string {
	return t.Message
}

type Tickets []Ticket
