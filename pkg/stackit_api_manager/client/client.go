package client

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
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
	baseURL string
	ctx     context.Context
}

// NewClient for API Manager interaction
func NewClient(baseURL, token string) *Client {
	ctx := context.WithValue(context.Background(), ContextAccessToken, token)
	baseURL = strings.TrimSuffix(baseURL, "/")
	return &Client{
		baseURL: baseURL,
		ctx:     ctx,
	}
}

func (c *Client) doRequest(req *http.Request) error {
	// AccessToken Authentication
	if auth, ok := c.ctx.Value(ContextAccessToken).(string); ok {
		req.Header.Add("Authorization", auth)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%w: %s", errInvalidRequest, string(body))
	}

	return nil
}
