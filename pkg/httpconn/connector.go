package httpconn

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

var defaultTimeout = time.Second * 15

type Connector struct {
	cli  *http.Client
	base *url.URL
}

func NewConnector(base *url.URL, timeout *time.Duration) *Connector {
	if timeout == nil {
		timeout = &defaultTimeout
	}

	return &Connector{
		base: base,
		cli: &http.Client{
			Timeout: *timeout,
		},
	}
}

func (cn *Connector) Post(ctx context.Context, path string, body []byte, headers map[string]string) (*int, []byte, error) {
	u, err := cn.base.Parse(path)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to parse url: %w", err)
	}

	r, err := http.NewRequest(http.MethodPost, u.String(), bytes.NewReader(body))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create request: %w", err)
	}

	setHeaders(r, headers)

	return cn.Do(ctx, r)
}

func (cn *Connector) Get(ctx context.Context, path string, headers map[string]string) (*int, []byte, error) {
	u, err := cn.base.Parse(path)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to parse url: %w", err)
	}

	r, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create request: %w", err)
	}

	setHeaders(r, headers)

	return cn.Do(ctx, r)
}

func (cn *Connector) Do(ctx context.Context, r *http.Request) (*int, []byte, error) {
	r = r.WithContext(ctx)
	// put here auth header adding
	resp, err := cn.cli.Do(r)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to do request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read body: %w", err)
	}

	return &resp.StatusCode, body, nil
}

func setHeaders(r *http.Request, headers map[string]string) {
	if headers == nil && len(headers) == 0 {
		return
	}

	for key, val := range headers {
		r.Header.Set(key, val)
	}
}
