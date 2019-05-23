package client

import (
	"io"
	"net/http"

	"github.com/palantir/stacktrace"
)

// Get -
func (c *Client) Get(url string, headers http.Header) (*http.Response, error) {
	var response *http.Response
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return response, stacktrace.Propagate(err, "[GET] request creation failed")
	}

	request.Header = headers

	return c.Impl.Do(request)
}

// Post -
func (c *Client) Post(url string, body io.Reader, headers http.Header) (*http.Response, error) {
	var response *http.Response
	request, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return response, stacktrace.Propagate(err, "[POST] request creation failed")
	}

	request.Header = headers

	return c.Impl.Do(request)
}

// Put -
func (c *Client) Put(url string, body io.Reader, headers http.Header) (*http.Response, error) {
	var response *http.Response
	request, err := http.NewRequest(http.MethodPut, url, body)
	if err != nil {
		return response, stacktrace.Propagate(err, "[PUT] request creation failed")
	}

	request.Header = headers

	return c.Impl.Do(request)
}

// Delete -
func (c *Client) Delete(url string, headers http.Header) (*http.Response, error) {
	var response *http.Response
	request, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return response, stacktrace.Propagate(err, "[DELETE] request creation failed")
	}
	request.Header = headers

	return c.Impl.Do(request)
}
