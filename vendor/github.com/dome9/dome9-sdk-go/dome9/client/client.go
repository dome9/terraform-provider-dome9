package client

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"

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
