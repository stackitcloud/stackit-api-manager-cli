package client

import (
	"fmt"
	"reflect"
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

func mockClient(t *testing.T) *Client {
	// start http mock server & reset when test is finished
	httpmock.Activate()
	t.Cleanup(func() { httpmock.DeactivateAndReset() })

	// return new client
	return NewClient(mockServerURL, "some-token")
}

func (m *mockResponses) mockJSONHTTPResponse(t *testing.T, method string) {
	jsonResponse, err := httpmock.NewJsonResponder(200, m.body)
	if err != nil {
		t.Error(err)
	}

	httpmock.RegisterResponder(method, fmt.Sprintf("%s%s", mockServerURL, m.path), jsonResponse)
}

func TestNewClient(t *testing.T) {
	type args struct {
		baseURL string
		token   string
	}
	tests := []struct {
		name string
		args args
		want *Client
	}{
		{
			name: "success",
			args: args{
				baseURL: mockServerURL,
				token:   "some-token",
			},
			want: NewClient(mockServerURL, "some-token"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewClient(tt.args.baseURL, tt.args.token)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClient() = %v, want %v", got, tt.want)
			}
			token, _ := got.ctx.Value(ContextAccessToken).(string)
			if !reflect.DeepEqual(token, tt.args.token) {
				t.Errorf("NewClient() token = %v, want %v", token, tt.args.token)
			}
		})
	}
}
