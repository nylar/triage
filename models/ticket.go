package models

// Ticket holds data on a support ticket
type Ticket struct {
	TicketID int    `json:"ticket_id"`
	Message  string `json:"message"`
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
