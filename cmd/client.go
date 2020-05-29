package cmd

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"path"

	"github.com/pkg/errors"
)

// Clientの構造体定義
type Client struct {
	EndpointURL *url.URL
	HTTPClient  *http.Client
}

// コンストラクタの定義
func newClient(endpointURL string, httpClient *http.Client) (*Client, error) {

	parsedURL, err := url.ParseRequestURI(endpointURL)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse url: %s", endpointURL)
	}
	client := &Client{
		EndpointURL: parsedURL,
		HTTPClient:  httpClient,
	}
	return client, nil
}

// HTTPリクエスト生成
func (c *Client) newRequest(ctx context.Context, method, spath string, body io.Reader) (*http.Request, error) {

	u := *c.EndpointURL
	u.Path = path.Join(c.EndpointURL.Path, spath)

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	return req, nil
}
