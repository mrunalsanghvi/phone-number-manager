package client

import (
	"net/http"
	"time"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type defaultHTTPClient struct {
	client *http.Client
}

func NewHTTPClient() HTTPClient {
	return &defaultHTTPClient{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *defaultHTTPClient) Do(req *http.Request) (*http.Response, error) {
	return c.client.Do(req)
}