package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/AndriyZaec/pokedexcli/internal/pokecache"
)

type Client struct {
	baseURL *url.URL
	http    *http.Client
	cache   *pokecache.Cache
}

func NewClient(base string) (*Client, error) {
	u, err := url.Parse(base)
	if err != nil {
		return nil, err
	}

	cache := pokecache.NewCache(2 * time.Minute)

	return &Client{
		baseURL: u,
		http:    &http.Client{},
		cache:   cache,
	}, nil
}

func (c *Client) getByURL(fullURL string, out any) error {
	if cache, ok := c.cache.Get(fullURL); ok {
		err := json.Unmarshal(cache, out)
		if err == nil {
			return nil
		}
	}

	req, err := http.NewRequest(http.MethodGet, fullURL, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		b, _ := io.ReadAll(io.LimitReader(resp.Body, 8<<10))
		return fmt.Errorf("api error: %s: %s", resp.Status, strings.TrimSpace(string(b)))
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	c.cache.Add(fullURL, b)

	if err := json.Unmarshal(b, out); err != nil {
		return err
	}

	return nil
}

func (c *Client) get(ctx context.Context, path string, query map[string]string, out any) error {
	rel, err := url.Parse(path)
	if err != nil {
		return err
	}

	u := c.baseURL.ResolveReference(rel)

	if len(query) > 0 {
		q := u.Query()
		for k, v := range query {
			q.Set(k, v)
		}
		u.RawQuery = q.Encode()
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		u.String(),
		nil,
	)
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		b, _ := io.ReadAll(io.LimitReader(resp.Body, 8<<10))
		return fmt.Errorf("api error: %s: %s", resp.Status, strings.TrimSpace(string(b)))
	}

	return json.NewDecoder(resp.Body).Decode(out)
}
