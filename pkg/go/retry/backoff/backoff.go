package backoff

import (
	"time"

	bf "github.com/bifrostcloud/protoc-gen-httpclient/pkg/go/backoff/shared"
)

// Backoff -
type Backoff struct {
	Backoff bf.Backoff
}

// New -
func New(b bf.Backoff) Backoff {
	result := Backoff{}
	result.Backoff = b
	return result
}

// NextInterval -
func (b *Backoff) NextInterval(retry int) time.Duration {
	return b.Backoff.NextInterval(retry)
}
