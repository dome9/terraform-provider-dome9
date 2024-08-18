package client

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/google/go-querystring/query"

	"github.com/dome9/dome9-sdk-go/dome9"
)

type Client struct {
	Config *dome9.Config
}

// NewClient returns a new client for the specified apiKey.
func NewClient(config *dome9.Config) (c *Client) {
	if config == nil {
		config, _ = dome9.NewConfig("", "", "")
	}
	c = &Client{Config: config}
	return
}

func (client *Client) NewRequestDo(method, url string, options, body, v interface{}) (*http.Response, error) {
	req, err := client.newRequest(method, url, options, body)
	if err != nil {
		return nil, err
	}
	client.logRequest(req)
	return client.do(req, v)
}

func (client *Client) NewRequestDoRetryWithOptions(method, url string, options, body, v interface{}, maxRetries int, retrySleepBetweenSecs int, shouldRetry func(*http.Response) bool) (*http.Response, error) {
	// Initialize the response and error variables outside the loop
	var resp *http.Response
	var err error

	// Attempt the request up to maxRetries times
	for i := 0; i < maxRetries; i++ {
		// Make the request
		resp, err = client.NewRequestDo(method, url, options, body, v)
		if err == nil {
			// If the request was successful, return the response
			return resp, nil
		}
		// If status code is 429 (API throttling), set the retrySleepBetweenSecs to 10 and maxRetries to fit 5 minutes of retries
		if resp.StatusCode == http.StatusTooManyRequests {
			retrySleepBetweenSecs = 10
			maxRetries = (5 * 60) / retrySleepBetweenSecs
		}
		if shouldRetry(resp) {
			time.Sleep(time.Second * time.Duration(retrySleepBetweenSecs))
		} else {
			return resp, err
		}
	}

	// If the function hasn't returned after maxRetries, return an error
	return nil, err
}

func (client *Client) NewRequestDoRetry(method, url string, options, body, v interface{}, shouldRetry func(*http.Response) bool) (*http.Response, error) {
	if shouldRetry == nil {
		// Default retry only on 4xx and 5xx status codes
		shouldRetry = func(resp *http.Response) bool {
			return resp != nil && resp.StatusCode >= 400 && resp.StatusCode < 600
		}
	}
	return client.NewRequestDoRetryWithOptions(method, url, options, body, v, 3, 5, shouldRetry)
}

// Generating the Http request
func (client *Client) newRequest(method, urlPath string, options, body interface{}) (*http.Request, error) {
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	// Join the path to the base-url
	u := *client.Config.BaseURL
	unescaped, err := url.PathUnescape(urlPath)
	if err != nil {
		return nil, err
	}

	// Set the encoded path data
	u.RawPath = client.Config.BaseURL.Path + urlPath
	u.Path = client.Config.BaseURL.Path + unescaped

	// Set the query parameters
	if options != nil {
		q, err := query.Values(options)
		if err != nil {
			return nil, err
		}
		u.RawQuery = q.Encode()
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(client.Config.AccessID, client.Config.SecretKey)
	req.Header.Add("Accept", "application/json")
	if body != nil {
		req.Header.Add("Content-Type", "application/json")
	}

	return req, nil
}

func (client *Client) do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := client.Config.HTTPClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer func() {
		if rerr := resp.Body.Close(); err == nil {
			err = rerr
		}
	}()

	if err := checkErrorInResponse(resp); err != nil {
		return resp, err
	}

	if v != nil {
		if err := decodeJSON(resp, v); err != nil {
			return resp, err
		}
	}
	client.logResponse(resp)

	return resp, nil
}

func decodeJSON(res *http.Response, v interface{}) error {
	return json.NewDecoder(res.Body).Decode(v)
}
