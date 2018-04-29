package clock

import "time"

// Clock mocks the time package
type Clock interface {
	Now() time.Time
}

// Real implements the Clock interface
type Real struct{}

// Now returns the current time
func (Real) Now() time.Time {
	return time.Now()
}
