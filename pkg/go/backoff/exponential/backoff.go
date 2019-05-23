package exponential

import (
	"math"
	"math/rand"
	"time"

	shared "github.com/bifrostcloud/protoc-gen-httpclient/pkg/go/backoff/shared"
)

// Backoff -
type Backoff struct {
	InitialTimeout    float64
	MaxTimeout        float64
	MaxJitterInterval int64
	Exponent          float64
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// New -
func New(initialTimeout, maxTimeout, maximumJitterInterval time.Duration, exponent float64) shared.Backoff {
	result := &Backoff{}
	result.InitialTimeout = float64(initialTimeout / time.Millisecond)
	result.MaxTimeout = float64(maxTimeout / time.Millisecond)
	result.MaxJitterInterval = int64(maximumJitterInterval / time.Millisecond)
	result.Exponent = exponent
	return result
}

// NextInterval -
func (b *Backoff) NextInterval(retry int) time.Duration {
	if retry <= 0 {
		return 0 * time.Millisecond
	}
	to := b.InitialTimeout + math.Pow(b.Exponent, float64(retry))
	if to > b.MaxTimeout {
		to = b.MaxTimeout
	}
	res := to + float64(rand.Int63n(b.MaxJitterInterval))
	return time.Duration(res) * time.Millisecond
}
