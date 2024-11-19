package pkg

import "time"

// Module represents a single, individually updateable status module.
// For example, a module could be a clock, a battery status, or network traffic.
type Module struct {
	Name     string
	Interval time.Duration
	Update   func() string
	Ticker   chan struct{}
}
