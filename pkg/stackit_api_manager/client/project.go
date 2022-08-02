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
	Stage string `json:"stage"`
}

type OpenAPI struct {
	Base64Encoded string `json:"base64Encoded"`
}

type Spec struct {
	OpenAPI *OpenAPI `json:"openApi,omitempty"`
}

type ProjectPublish struct {
	Metadata Metadata `json:"metadata"`
	Spec     Spec     `json:"spec"`
}

type ProjectRetire struct {
	Metadata Metadata `json:"metadata"`
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
func (c *Client) ProjectPublish(
	projectID,
	identifier string,
	projectPublish *ProjectPublish,
) error {
	url := fmt.Sprintf("%s/v1/projects/%s/api/%s", c.baseURL, projectID, identifier)
	j, err := json.Marshal(projectPublish)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(j))
	if err != nil {
		return err
	}
	return c.doRequest(req)
}

// ProjectRetire a OpenAPI Spec file for a project
func (c *Client) ProjectRetire(
	projectID,
	identifier string,
	projectRetire *ProjectRetire,
) error {
	url := fmt.Sprintf("%s/v1/projects/%s/api/%s", c.baseURL, projectID, identifier)
	j, err := json.Marshal(projectRetire)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodDelete, url, bytes.NewBuffer(j))
	if err != nil {
		return err
	}
	return c.doRequest(req)
}
