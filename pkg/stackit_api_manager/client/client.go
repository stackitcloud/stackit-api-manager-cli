package client

import (
	"errors"
	"io"
	"net/http"
)

var (
	errInvalidRequest = errors.New("invalid API request")
)

type Client struct {
	URL   *string
	Token *string
}

// NewClient for API Manager interaction
func NewClient(url, token string) *Client {
	return &Client{
		URL:   &url,
		Token: &token,
	}
}

func (c *Client) doRequest(req *http.Request) ([]byte, *http.Response, error) {
	req.Header.Set("token", *c.Token) // todo add real auth
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
		return nil, resp, errInvalidRequest
	}

	return body, resp, nil
}
