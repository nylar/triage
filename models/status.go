package models

type Status struct {
	StatusID int    `json:"status_id"`
	Name     string `json:"name"`
}

func (s *Status) String() string {
	return s.Name
}
