package shared

import "time"

// Retriable -
type Retriable interface {
	NextInterval(retry int) time.Duration
}
