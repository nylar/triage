package models

type StatusFlag int

const (
	Open StatusFlag = iota + 1
	Closed
	OnHold

	defaultStatus = Open
)

type Status struct {
	StatusID int    `json:"status_id"`
	Name     string `json:"name"`
}

func (s *Status) String() string {
	return s.Name
}

func DefaultStatus() *Status {
	return &Status{StatusID: int(defaultStatus), Name: "open"}
}
