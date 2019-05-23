package exponential

import (
	"math/rand"
	"time"

	shared "github.com/bifrostcloud/protoc-gen-httpclient/pkg/go/backoff/shared"
)

// Backoff -
type Backoff struct {
	Interval          int64
	MaxJitterInterval int64
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// New -
func New(interval, maxJitterInterval time.Duration) shared.Backoff {
	result := &Backoff{}
	result.Interval = int64(interval / time.Millisecond)
	result.MaxJitterInterval = int64(maxJitterInterval / time.Millisecond)
	return result
}

// NextInterval -
func (b *Backoff) NextInterval(retry int) time.Duration {
	if retry <= 0 {
		return 0 * time.Millisecond
	}
	rnd := rand.Int63n(b.MaxJitterInterval)
	return (time.Duration(b.Interval) * time.Millisecond) + (time.Duration(rnd) * time.Millisecond)
}
