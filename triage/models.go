package triage

type Ticket struct {
	Message string
}

func NewTicket(msg string) *Ticket {
	return &Ticket{
		Message: msg,
	}
}

func (t *Ticket) String() string {
	return t.Message
}
