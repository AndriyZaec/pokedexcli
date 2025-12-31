package api

import (
	"context"
	"net/http"
	"net/url"
)

type Client struct {
	baseUrl *url.URL
	http    *http.Client
}

func NewClient(base string) (*Client, error) {
	u, err := url.Parse(base)
	if err != nil {
		return nil, err
	}

	return &Client{
		baseUrl: u,
		http:    &http.Client{},
	}, nil
}

func (c *Client) get(ctx context.Context, path string, query map[string]string, out any) error {
}
