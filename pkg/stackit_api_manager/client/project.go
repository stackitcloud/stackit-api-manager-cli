package client

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/hashicorp/go-multierror"
)

type Metadata struct {
	Identifier *string `json:"identifier,omitempty"`
	Stage      *string `json:"stage,omitempty"`
}

type OpenAPI struct {
	Base64Encoded *string `json:"base64Encoded,omitempty"`
}

type Spec struct {
	OpenAPI *OpenAPI `json:"openApi,omitempty"`
}

type ProjectPublish struct {
	Metadata *Metadata `json:"metadata,omitempty"`
	Spec     *Spec     `json:"spec,omitempty"`
}

type ProjectRetire struct {
	Metadata *Metadata `json:"metadata,omitempty"`
}

type ProjectPublishResponse struct {
	Code    *int    `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Details []*struct {
		Type *string `json:"@type,omitempty"`
	} `json:"details,omitempty"`
}

type ProjectRetireResponse struct {
	Code    *int    `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Details []*struct {
		Type *string `json:"@type,omitempty"`
	} `json:"details,omitempty"`
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
	url := fmt.Sprintf("%s/v1/projects/%s/publish", c.baseURL, projectID)
	j, err := json.Marshal(projectPublish)
	if err != nil {
		return nil, nil, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(j))
	if err != nil {
		return nil, nil, err
	}

	var response ProjectPublishResponse
	body, resp, err := c.doRequest(req)

	errResponse := json.Unmarshal(body, &response)
	if errResponse != nil {
		err = multierror.Append(err, errResponse)
	}
	return &response, resp, err
}

// ProjectRetire a OpenAPI Spec file for a project
func (c *Client) ProjectRetire( //nolint:dupl // API request
	projectID string,
	projectRetire *ProjectRetire,
) (*ProjectRetireResponse, *http.Response, error) {
	url := fmt.Sprintf("%s/v1/projects/%s/retire", c.baseURL, projectID)
	j, err := json.Marshal(projectRetire)
	if err != nil {
		return nil, nil, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(j))
	if err != nil {
		return nil, nil, err
	}

	var response ProjectRetireResponse
	body, resp, err := c.doRequest(req)

	errResponse := json.Unmarshal(body, &response)
	if errResponse != nil {
		err = multierror.Append(err, errResponse)
	}
	return &response, resp, err
}
