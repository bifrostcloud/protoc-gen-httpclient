package shared

import "time"

// Backoff -
type Backoff interface {
	NextInterval(retry int) time.Duration
}
