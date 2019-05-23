package client

import (
	basic "github.com/bifrostcloud/protoc-gen-httpclient/pkg/go/client/basic"
	circuitbreaker "github.com/bifrostcloud/protoc-gen-httpclient/pkg/go/client/circuit-breaker"
	shared "github.com/bifrostcloud/protoc-gen-httpclient/pkg/go/client/shared"
)

// Client -
type Client struct {
	Impl shared.ClientInterface
	// Password string
}

// NewBasicClient -
func NewBasicClient(opts ...basic.Option) *Client {
	result := &Client{}
	result.Impl = basic.New(opts...)
	return result
}

// NewClientWithCircuitBreaker -
func NewClientWithCircuitBreaker(opts ...circuitbreaker.Option) *Client {
	result := &Client{}
	result.Impl = circuitbreaker.New(opts...)
	return result
}
