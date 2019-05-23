package shared

import (
	"net/http"
	"time"

	"github.com/palantir/stacktrace"
)

// Errors
var (
	Err5xx = stacktrace.NewError("server returned 5xx status code")
)

// Constants
const (
	DefaultRetryCount             = 0
	DefaultTimeout                = 30 * time.Second
	DefaultMaxConcurrentRequests  = 100
	DefaultErrorPercentThreshold  = 25
	DefaultSleepWindow            = 10
	DefaultRequestVolumeThreshold = 10
)

// ClientInterface -
type ClientInterface interface {
	Do(req *http.Request) (*http.Response, error)
}
