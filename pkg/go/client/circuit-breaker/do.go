package circuitbreaker

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"

	client "github.com/bifrostcloud/protoc-gen-httpclient/pkg/go/client/shared"

	"github.com/afex/hystrix-go/hystrix"
)

// Do - http.Client implements do method
func (c *Client) Do(request *http.Request) (*http.Response, error) {
	var response *http.Response
	var err error

	var bodyReader *bytes.Reader

	if request.Body != nil {
		reqData, err := ioutil.ReadAll(request.Body)
		if err != nil {
			return nil, err
		}
		bodyReader = bytes.NewReader(reqData)
		request.Body = ioutil.NopCloser(bodyReader) // prevents closing the body between retries
	}

	for i := 0; i <= c.RetryCount; i++ {
		if response != nil {
			response.Body.Close()
		}

		err = hystrix.Do(c.HystrixCommandName, func() error {
			response, err = c.Client.Do(request)
			if bodyReader != nil {
				_, _ = bodyReader.Seek(0, 0)
			}

			if err != nil {
				return err
			}

			if response.StatusCode >= http.StatusInternalServerError {
				return client.Err5xx
			}
			return nil
		}, c.FallbackFunc)

		if err != nil {
			backoffTime := c.Retrier.NextInterval(i)
			time.Sleep(backoffTime)
			continue
		}

		break
	}

	if err == client.Err5xx {
		return response, nil
	}

	return response, err
}
