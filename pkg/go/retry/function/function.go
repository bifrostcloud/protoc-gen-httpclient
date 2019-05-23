package function

import (
	"time"

	shared "github.com/bifrostcloud/protoc-gen-httpclient/pkg/go/retry/shared"
)

// Func -
type Func func(retry int) time.Duration

// New -
func New(f Func) shared.Retriable {
	return f
}

// NextInterval -
func (f Func) NextInterval(retry int) time.Duration {
	return f(retry)
}
