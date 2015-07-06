package models

// Ticket holds data on a support ticket
type Ticket struct {
	TicketID int    `json:"ticket_id"`
	Message  string `json:"message"`
	Status   Status `json:"status"`
}

// NewTicket instantiates a new ticket
func NewTicket(msg string) *Ticket {
	status := DefaultStatus()
	return &Ticket{
		Message: msg,
		Status:  *status,
	}
}

func (t *Ticket) String() string {
	return t.Message
}

type Tickets []Ticket
