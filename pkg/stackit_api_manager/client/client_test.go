package client

import (
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
)

const (
	mockServerURL = "http://localhost"
)

type mockResponses struct {
	path string
	body interface{}
}

func stringPtn(value string) *string { return &value }
func intPtn(value int) *int          { return &value }

func mockClient(t *testing.T) *Client {
	// start http mock server & reset when test is finished
	httpmock.Activate()
	t.Cleanup(func() { httpmock.DeactivateAndReset() })

	// return new client
	return NewClient(mockServerURL, "some-token")
}

func mockJSONHTTPResponse(t *testing.T, path string, body interface{}) {
	jsonResponse, err := httpmock.NewJsonResponder(200, body)
	if err != nil {
		t.Error(err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s%s", mockServerURL, path), jsonResponse)
}
