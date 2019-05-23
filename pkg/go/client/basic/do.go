package basic

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"

	multierror "github.com/bifrostcloud/protoc-gen-httpclient/pkg/go/multi-error"
)

// Do - http.Client implements do method
func (c *Client) Do(request *http.Request) (*http.Response, error) {
	request.Close = true

	var bodyReader *bytes.Reader

	if request.Body != nil {
		reqData, err := ioutil.ReadAll(request.Body)
		if err != nil {
			return nil, err
		}
		bodyReader = bytes.NewReader(reqData)
		request.Body = ioutil.NopCloser(bodyReader)
	}
	multiErr := &multierror.MultiError{}
	var response *http.Response

	for i := 0; i <= c.RetryCount; i++ {
		if response != nil {
			response.Body.Close()
		}

		var err error
		response, err = c.Client.Do(request)
		if bodyReader != nil {
			_, _ = bodyReader.Seek(0, 0)
		}

		if err != nil {
			multiErr.Add(err.Error())

			backoffTime := c.Retrier.NextInterval(i)
			time.Sleep(backoffTime)
			continue
		}

		if response.StatusCode >= http.StatusInternalServerError {
			backoffTime := c.Retrier.NextInterval(i)
			time.Sleep(backoffTime)
			continue
		}

		multiErr = &multierror.MultiError{}
		break
	}

	return response, multiErr.HasError()
}
