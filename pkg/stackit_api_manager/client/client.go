package client

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
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

func EncodeOpenAPISpecFile(file string) (string, error) {
	oas, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(oas), nil
}

func (c *Client) ProjectPublish(projectID string, projectPublish *ProjectPublish) (*ProjectResponse, *http.Response, error) {
	url := fmt.Sprintf("%s/v1/projects/%s/publish", *c.URL, projectID)
	log.Println(url)
	j, err := json.Marshal(projectPublish)
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

	var response ProjectResponse
	err = json.Unmarshal(body, &response)
	return &response, resp, err
}

func (c *Client) ProjectRetire(projectID string, projectRetire *ProjectRetire) (*ProjectResponse, *http.Response, error) {
	url := fmt.Sprintf("%s/v1/projects/%s/retire", *c.URL, projectID)
	fmt.Println(url)
	j, err := json.Marshal(projectRetire)
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

	var response ProjectResponse
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
