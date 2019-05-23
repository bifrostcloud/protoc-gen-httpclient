package basic

import (
	"net/http"
	"time"

	client "github.com/bifrostcloud/protoc-gen-httpclient/pkg/go/client/shared"
	retry "github.com/bifrostcloud/protoc-gen-httpclient/pkg/go/retry/shared"
)

// Client -
type Client struct {
	http.Client
	RetryCount int
	Retrier    retry.Retriable
}

// Option -
type Option func(*Client)

// New -
func New(opts ...Option) client.ClientInterface {
	result := &Client{}
	result.Transport = http.DefaultTransport
	result.Timeout = client.DefaultTimeout
	result.RetryCount = client.DefaultRetryCount
	result.Retrier = nil
	for _, opt := range opts {
		opt(result)
	}
	return result
}

// Timeout -
func Timeout(timeout time.Duration) Option {
	return func(c *Client) {
		c.Timeout = timeout
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
