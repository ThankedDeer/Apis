package httpclient

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	httpClient *http.Client
}

func Newclient() *Client {
	return &Client{
		httpClient: &http.Client{},
	}
}

func (cl *Client) CreateRequest(
	c context.Context,
	url string,
	method string,
	body io.Reader,
	headers map[string]string,
) (*http.Response, error) {

	req, err := http.NewRequestWithContext(c, method, url, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Default headers
	req.Header.Set("Accept", "application/json")

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := cl.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}

	// Accept any 2xx as success
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		defer resp.Body.Close()
		return nil, fmt.Errorf(
			"external API returned non-success status %d: %s",
			resp.StatusCode,
			http.StatusText(resp.StatusCode),
		)
	}

	return resp, nil
}
