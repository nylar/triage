package models

// Ticket holds data on a support ticket
type Ticket struct {
	TicketID int    `json:"ticket_id"`
	Subject  string `json:"subject"`
	Status   Status `json:"status"`
}

// NewTicket instantiates a new ticket
func NewTicket(subject string) *Ticket {
	status := DefaultStatus()
	return &Ticket{
		Subject: subject,
		Status:  *status,
	}
}

func (t *Ticket) String() string {
	return t.Subject
}

type Tickets []Ticket
