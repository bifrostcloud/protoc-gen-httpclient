package circuitbreaker

import (
	"net/http"
	"time"

	"github.com/afex/hystrix-go/hystrix"
	client "github.com/bifrostcloud/protoc-gen-httpclient/pkg/go/client/shared"
	retry "github.com/bifrostcloud/protoc-gen-httpclient/pkg/go/retry/shared"
	utils "github.com/bifrostcloud/protoc-gen-httpclient/pkg/go/utils"
)

type fallbackFunc func(error) error

// Client -  is the hystrix client implementation
type Client struct {
	http.Client
	HystrixTimeout         time.Duration
	HystrixCommandName     string
	MaxConcurrentRequests  int
	RequestVolumeThreshold int
	SleepWindow            int
	ErrorPercentThreshold  int
	RetryCount             int
	Retrier                retry.Retriable
	FallbackFunc           func(err error) error
}

// Option -
type Option func(*Client)

// New -
func New(opts ...Option) client.ClientInterface {
	result := &Client{}
	result.Transport = http.DefaultTransport
	result.Timeout = client.DefaultTimeout
	result.HystrixTimeout = client.DefaultTimeout
	result.MaxConcurrentRequests = client.DefaultMaxConcurrentRequests
	result.ErrorPercentThreshold = client.DefaultErrorPercentThreshold
	result.SleepWindow = client.DefaultSleepWindow
	result.RequestVolumeThreshold = client.DefaultRequestVolumeThreshold
	result.RetryCount = client.DefaultRetryCount
	result.Retrier = nil
	for _, opt := range opts {
		opt(result)
	}

	hystrix.ConfigureCommand(result.HystrixCommandName, hystrix.CommandConfig{
		Timeout:                utils.DurationToInt(client.DefaultTimeout, time.Millisecond),
		MaxConcurrentRequests:  client.DefaultMaxConcurrentRequests,
		RequestVolumeThreshold: client.DefaultRequestVolumeThreshold,
		SleepWindow:            client.DefaultSleepWindow,
		ErrorPercentThreshold:  client.DefaultErrorPercentThreshold,
	})

	return result
}

// CommandName -
func CommandName(name string) Option {
	return func(c *Client) {
		c.HystrixCommandName = name
	}
}

// Timeout sets hystrix timeout
func Timeout(timeout time.Duration) Option {
	return func(c *Client) {
		c.Timeout = timeout
	}
}

// HystrixTimeout -
func HystrixTimeout(timeout time.Duration) Option {
	return func(c *Client) {
		c.HystrixTimeout = timeout
	}
}

// MaxConcurrentRequests -
func MaxConcurrentRequests(maxConcurrentRequests int) Option {
	return func(c *Client) {
		c.MaxConcurrentRequests = maxConcurrentRequests
	}
}

// RequestVolumeThreshold -
func RequestVolumeThreshold(requestVolumeThreshold int) Option {
	return func(c *Client) {
		c.RequestVolumeThreshold = requestVolumeThreshold
	}
}

// SleepWindow -
func SleepWindow(sleepWindow int) Option {
	return func(c *Client) {
		c.SleepWindow = sleepWindow
	}
}

// ErrorPercentThreshold -
func ErrorPercentThreshold(errorPercentThreshold int) Option {
	return func(c *Client) {
		c.ErrorPercentThreshold = errorPercentThreshold
	}
}

// FallbackFunc -
func FallbackFunc(fn fallbackFunc) Option {
	return func(c *Client) {
		c.FallbackFunc = fn
	}
}

// RetryCount -
func RetryCount(retryCount int) Option {
	return func(c *Client) {
		c.RetryCount = retryCount
	}
}

// Retrier -
func Retrier(retrier retry.Retriable) Option {
	return func(c *Client) {
		c.Retrier = retrier
	}
}
