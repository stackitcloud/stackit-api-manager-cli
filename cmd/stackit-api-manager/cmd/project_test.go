package cmd

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/spf13/cobra"
)

const (
	mockServerURL = "http://localhost"
)

type projectCmdArgs struct {
	serverBaseURL       string
	authToken           string
	projectID           string
	identifier          string
	stage               string
	openAPISpecFilePath string
}

type mockResponses struct {
	path       string
	statusCode int
	body       interface{}
}

func (m *mockResponses) mockJSONHTTPResponse(t *testing.T, method string) {
	jsonResponse, err := httpmock.NewJsonResponder(m.statusCode, m.body)
	if err != nil {
		t.Error(err)
	}
	httpmock.RegisterResponder(method, fmt.Sprintf("%s%s", mockServerURL, m.path), jsonResponse)
}

// setArgs for project CMD CLI flags
func (args *projectCmdArgs) setArgs() {
	serverBaseURL = args.serverBaseURL
	authToken = args.authToken
	projectID = args.projectID
	identifier = args.identifier
	stage = args.stage
	openAPISpecFilePath = args.openAPISpecFilePath
}

func Test_publishCmdRunE(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	tests := []struct {
		name            string
		args            projectCmdArgs
		mockResponses   []mockResponses
		mockNilResponse bool
		wantErr         bool
	}{
		{
			name: "success",
			args: projectCmdArgs{
				serverBaseURL:       mockServerURL,
				authToken:           "some-auth-token",
				projectID:           "some-project-id",
				identifier:          "some-identifier",
				stage:               "some-stage",
				openAPISpecFilePath: "../../../pkg/stackit_api_manager/util/test_data/test.json",
			},
			mockResponses: []mockResponses{
				{
					path:       "/v1/projects/some-project-id/api/some-identifier",
					statusCode: 200,
				},
			},
			wantErr: false,
		},
		{
			name: "missing file - error",
			args: projectCmdArgs{
				openAPISpecFilePath: "./no-test.json",
			},
			wantErr: true,
		},
		{
			name: "status code 400 - error",
			args: projectCmdArgs{
				serverBaseURL:       mockServerURL,
				authToken:           "some-auth-token",
				projectID:           "some-project-id",
				identifier:          "some-identifier",
				stage:               "some-stage",
				openAPISpecFilePath: "../../../pkg/stackit_api_manager/util/test_data/test.json",
			},
			mockResponses: []mockResponses{
				{
					path:       "/v1/projects/some-project-id/api/some-identifier",
					statusCode: 400,
				},
			},
			wantErr: true,
		},
		{
			name: "nil http response - error",
			args: projectCmdArgs{
				serverBaseURL:       mockServerURL,
				authToken:           "some-auth-token",
				projectID:           "some-project-id",
				identifier:          "some-identifier",
				stage:               "some-stage",
				openAPISpecFilePath: "../../../pkg/stackit_api_manager/util/test_data/test.json",
			},
			mockNilResponse: true,
			wantErr:         true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.setArgs()
			for _, mockResponse := range tt.mockResponses {
				mockResponse.mockJSONHTTPResponse(t, http.MethodPost)
			}
			if tt.mockNilResponse {
				httpmock.Reset()
			}
			if err := publishCmdRunE(&cobra.Command{Use: "publish"}, []string{}); (err != nil) != tt.wantErr {
				t.Errorf("publishCmdRunE() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_retireCmdRunE(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	tests := []struct {
		name            string
		args            projectCmdArgs
		mockResponses   []mockResponses
		mockNilResponse bool
		wantErr         bool
	}{
		{
			name: "success",
			args: projectCmdArgs{
				serverBaseURL: mockServerURL,
				authToken:     "some-auth-token",
				projectID:     "some-project-id",
				identifier:    "some-identifier",
			},
			mockResponses: []mockResponses{
				{
					path:       "/v1/projects/some-project-id/api/some-identifier",
					statusCode: 200,
				},
			},
			wantErr: false,
		},
		{
			name: "status code 400 - error",
			args: projectCmdArgs{
				serverBaseURL: mockServerURL,
				authToken:     "some-auth-token",
				projectID:     "some-project-id",
				identifier:    "some-identifier",
			},
			mockResponses: []mockResponses{
				{
					path:       "/v1/projects/some-project-id/api/some-identifier",
					statusCode: 400,
				},
			},
			wantErr: true,
		},
		{
			name: "nil http response - error",
			args: projectCmdArgs{
				serverBaseURL: mockServerURL,
				authToken:     "some-auth-token",
				projectID:     "some-project-id",
				identifier:    "some-identifier",
			},
			mockNilResponse: true,
			wantErr:         true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.setArgs()
			for _, mockResponse := range tt.mockResponses {
				mockResponse.mockJSONHTTPResponse(t, http.MethodDelete)
			}
			if tt.mockNilResponse {
				httpmock.Reset()
			}
			if err := retireCmdRunE(&cobra.Command{Use: "retire"}, []string{}); (err != nil) != tt.wantErr {
				t.Errorf("retireCmdRunE() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_validateCmdRunE(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	tests := []struct {
		name            string
		args            projectCmdArgs
		mockResponses   []mockResponses
		mockNilResponse bool
		wantErr         bool
	}{
		{
			name: "success",
			args: projectCmdArgs{
				serverBaseURL:       mockServerURL,
				authToken:           "some-auth-token",
				projectID:           "some-project-id",
				identifier:          "some-identifier",
				stage:               "some-stage",
				openAPISpecFilePath: "../../../pkg/stackit_api_manager/util/test_data/test.json",
			},
			mockResponses: []mockResponses{
				{
					path:       "/v1/projects/some-project-id/api/some-identifier/validate",
					statusCode: 200,
				},
			},
			wantErr: false,
		},
		{
			name: "error: missing file",
			args: projectCmdArgs{
				openAPISpecFilePath: "./no-test.json",
			},
			wantErr: true,
		},
		{
			name: "status code 400 - error",
			args: projectCmdArgs{
				serverBaseURL:       mockServerURL,
				authToken:           "some-auth-token",
				projectID:           "some-project-id",
				identifier:          "some-identifier",
				stage:               "some-stage",
				openAPISpecFilePath: "../../../pkg/stackit_api_manager/util/test_data/test.json",
			},
			mockResponses: []mockResponses{
				{
					path:       "/v1/projects/some-project-id/api/some-identifier/validate",
					statusCode: 400,
				},
			},
			wantErr: true,
		},
		{
			name: "nil http response - error",
			args: projectCmdArgs{
				serverBaseURL:       mockServerURL,
				authToken:           "some-auth-token",
				projectID:           "some-project-id",
				identifier:          "some-identifier",
				stage:               "some-stage",
				openAPISpecFilePath: "../../../pkg/stackit_api_manager/util/test_data/test.json",
			},
			mockNilResponse: true,
			wantErr:         true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.setArgs()
			for _, mockResponse := range tt.mockResponses {
				mockResponse.mockJSONHTTPResponse(t, http.MethodPost)
			}
			if tt.mockNilResponse {
				httpmock.Reset()
			}
			if err := validateCmdRunE(&cobra.Command{Use: "validate"}, []string{}); (err != nil) != tt.wantErr {
				t.Errorf("validateCmdRunE() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_listCmdRunE(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	tests := []struct {
		name            string
		args            projectCmdArgs
		mockResponses   []mockResponses
		mockNilResponse bool
		wantErr         bool
	}{
		{
			name: "success",
			args: projectCmdArgs{
				serverBaseURL: mockServerURL,
				authToken:     "some-auth-token",
				projectID:     "some-project-id",
			},
			mockResponses: []mockResponses{
				{
					path:       "/v1/projects/some-project-id",
					statusCode: 200,
				},
			},
			wantErr: false,
		},
		{
			name: "missing project id - error",
			args: projectCmdArgs{
				serverBaseURL: mockServerURL,
				authToken:     "some-auth-token",
			},
			wantErr: true,
		},
		{
			name: "status code 400 - error",
			args: projectCmdArgs{
				serverBaseURL: mockServerURL,
				authToken:     "some-auth-token",
				projectID:     "some-project-id",
			},
			mockResponses: []mockResponses{
				{
					path:       "/v1/projects/some-project-id",
					statusCode: 400,
				},
			},
			wantErr: true,
		},
		{
			name: "nil http response - error",
			args: projectCmdArgs{
				serverBaseURL:       mockServerURL,
				authToken:           "some-auth-token",
				projectID:           "some-project-id",
				identifier:          "some-identifier",
				stage:               "some-stage",
				openAPISpecFilePath: "../../../pkg/stackit_api_manager/util/test_data/test.json",
			},
			mockNilResponse: true,
			wantErr:         true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.setArgs()
			for _, mockResponse := range tt.mockResponses {
				mockResponse.mockJSONHTTPResponse(t, http.MethodGet)
			}
			if tt.mockNilResponse {
				httpmock.Reset()
			}
			if err := listCmdRunE(&cobra.Command{Use: "list"}, []string{}); (err != nil) != tt.wantErr {
				t.Errorf("listCmdRunE() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_fetchAPICmsdRunE(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	type OpenAPI struct {
		Base64Encoded string `json:"base64_encoded"`
	}
	type Spec struct {
		OpenAPI OpenAPI `json:"open_api"`
	}
	type fetchResponseBody struct {
		Stage       string `json:"stage"`
		APIURL      string `json:"api_url"`
		UpstreamURL string `json:"upstream_url"`
		Spec        Spec   `json:"spec"`
	}

	tests := []struct {
		name            string
		args            projectCmdArgs
		mockResponses   []mockResponses
		mockNilResponse bool
		wantErr         bool
	}{
		{
			name: "success",
			args: projectCmdArgs{
				serverBaseURL: mockServerURL,
				authToken:     "some-auth-token",
				projectID:     "some-project-id",
				identifier:    "some-identifier",
			},
			mockResponses: []mockResponses{
				{
					path:       "/v1/projects/some-project-id/api/some-identifier",
					statusCode: 200,
					body:       fetchResponseBody{},
				},
			},
			wantErr: false,
		},
		{
			name: "missing project id - error",
			args: projectCmdArgs{
				serverBaseURL: mockServerURL,
				authToken:     "some-auth-token",
				identifier:    "some-identifier",
			},
			wantErr: true,
		},
		{
			name: "missing identifier - error",
			args: projectCmdArgs{
				serverBaseURL: mockServerURL,
				authToken:     "some-auth-token",
				projectID:     "some-project-id",
			},
			wantErr: true,
		},

		{
			name: "status code 400 - error",
			args: projectCmdArgs{
				serverBaseURL: mockServerURL,
				authToken:     "some-auth-token",
				projectID:     "some-project-id",
				identifier:    "some-identifier",
			},
			mockResponses: []mockResponses{
				{
					path:       "/v1/projects/some-project-id/api/some-identifier",
					statusCode: 400,
				},
			},
			wantErr: true,
		},
		{
			name: "nil http response - error",
			args: projectCmdArgs{
				serverBaseURL:       mockServerURL,
				authToken:           "some-auth-token",
				projectID:           "some-project-id",
				identifier:          "some-identifier",
				stage:               "some-stage",
				openAPISpecFilePath: "../../../pkg/stackit_api_manager/util/test_data/test.json",
			},
			mockNilResponse: true,
			wantErr:         true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.setArgs()
			for _, mockResponse := range tt.mockResponses {
				mockResponse.mockJSONHTTPResponse(t, http.MethodGet)
			}
			if tt.mockNilResponse {
				httpmock.Reset()
			}
			err := fetchCmdRunE(&cobra.Command{Use: "fetch"}, []string{})
			fmt.Printf("err: %v", err)
			if (err != nil) != tt.wantErr {
				t.Errorf("fetchAPICmdRunE() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
