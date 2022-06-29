package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	URL   *string
	Token *string
}

func NewClient(url, token string) *Client {
	return &Client{
		URL:   &url,
		Token: &token,
	}
}

func (c *Client) PublishAPIEndpoint(projectID string, publishAPIEndpoint *PublishAPIEndpoint) (*Response, *http.Response, error) {
	url := fmt.Sprintf("%s/v1/projects/%s/publish", *c.URL, projectID)
	fmt.Println(url)
	j, err := json.Marshal(publishAPIEndpoint)
	if err != nil {
		return nil, nil, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(j))
	if err != nil {
		return nil, nil, err
	}

	body, resp, err := c.doRequest(req)
	if err != nil {
		return nil, resp, err
	}

	var response Response
	err = json.Unmarshal(body, &response)
	return &response, resp, err
}

func (c *Client) RetireAPIEndpoint(projectID string, retireAPIEndpoint *RetireAPIEndpoint) (*Response, *http.Response, error) {
	url := fmt.Sprintf("%s/v1/projects/%s/retire", *c.URL, projectID)
	fmt.Println(url)
	j, err := json.Marshal(retireAPIEndpoint)
	if err != nil {
		return nil, nil, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(j))
	if err != nil {
		return nil, nil, err
	}

	body, resp, err := c.doRequest(req)
	if err != nil {
		return nil, resp, err
	}

	var response Response
	err = json.Unmarshal(body, &response)
	return &response, resp, err
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
		return nil, resp, fmt.Errorf("%s", body) // todo add real error
	}

	return body, resp, nil
}
