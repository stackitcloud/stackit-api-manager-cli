package client

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Metadata struct {
	Identifier *string `json:"identifier"`
	Stage      *string `json:"stage"`
}

type OpenAPI struct {
	Base64Encoded *string `json:"base64Encoded"`
}

type Spec struct {
	OpenAPI *OpenAPI `json:"openApi"`
}

type ProjectPublish struct {
	Metadata *Metadata `json:"metadata"`
	Spec     *Spec     `json:"spec"`
}

type ProjectRetire struct {
	Metadata *Metadata `json:"metadata"`
}

type ProjectPublishResponse struct {
	Code    *int    `json:"code"`
	Message *string `json:"message"`
	Details *[]struct {
		Type *string `json:"@type"`
	} `json:"details"`
}

type ProjectRetireResponse struct {
	Code    *int    `json:"code"`
	Message *string `json:"message"`
	Details *[]struct {
		Type *string `json:"@type"`
	} `json:"details"`
}

// EncodeBase64File to encode a file into base64 for uploading it via an API request
func EncodeBase64File(file string) (string, error) {
	oas, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(oas), nil
}

// ProjectPublish a OpenAPI Spec file for a project
func (c *Client) ProjectPublish( //nolint:dupl // API request
	projectID string,
	projectPublish *ProjectPublish,
) (*ProjectPublishResponse, *http.Response, error) {
	url := fmt.Sprintf("%s/v1/projects/%s/publish", *c.url, projectID)
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

	var response ProjectPublishResponse
	err = json.Unmarshal(body, &response)
	return &response, resp, err
}

// ProjectRetire a OpenAPI Spec file for a project
func (c *Client) ProjectRetire( //nolint:dupl // API request
	projectID string,
	projectRetire *ProjectRetire,
) (*ProjectRetireResponse, *http.Response, error) {
	url := fmt.Sprintf("%s/v1/projects/%s/retire", *c.url, projectID)
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

	var response ProjectRetireResponse
	err = json.Unmarshal(body, &response)
	return &response, resp, err
}
