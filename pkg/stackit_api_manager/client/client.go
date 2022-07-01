package client

import (
	"context"
	"errors"
	"io"
	"net/http"
)

type contextKey string

func (c contextKey) String() string {
	return "auth " + string(c)
}

const (
	ContextAccessToken = contextKey("accesstoken")
)

var (
	errInvalidRequest = errors.New("invalid API request")
)

type Client struct {
	url *string
	ctx context.Context
}

// NewClient for API Manager interaction
func NewClient(url, token string) *Client {
	ctx := context.WithValue(context.Background(), ContextAccessToken, token)
	return &Client{
		url: &url,
		ctx: ctx,
	}
}

func (c *Client) doRequest(req *http.Request) ([]byte, *http.Response, error) {
	// AccessToken Authentication
	if auth, ok := c.ctx.Value(ContextAccessToken).(string); ok {
		req.Header.Add("Authorization", "Bearer "+auth)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, resp, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp, err
	}

	if resp.StatusCode != http.StatusOK {
		return body, resp, errInvalidRequest
	}

	return body, resp, nil
}
