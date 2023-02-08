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

//nolint:dupl // ignore dupl linter error for testing
func Test_publishCmdRunE(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	tests := []struct {
		name          string
		args          projectCmdArgs
		mockResponses []mockResponses
		wantErr       bool
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
			name: "error: missing file",
			args: projectCmdArgs{
				openAPISpecFilePath: "./no-test.json",
			},
			wantErr: true,
		},
		{
			name: "status code 400 - no error",
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
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.setArgs()
			for _, mockResponse := range tt.mockResponses {
				mockResponse.mockJSONHTTPResponse(t, http.MethodPost)
			}
			if err := publishCmdRunE(&cobra.Command{}, []string{}); (err != nil) != tt.wantErr {
				t.Errorf("publishCmdRunE() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_retireCmdRunE(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	tests := []struct {
		name          string
		args          projectCmdArgs
		mockResponses []mockResponses
		wantErr       bool
	}{
		{
			name: "success",
			args: projectCmdArgs{
				serverBaseURL: mockServerURL,
				authToken:     "some-auth-token",
				projectID:     "some-project-id",
				identifier:    "some-identifier",
				stage:         "some-stage",
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
			name: "status code 400 - no error",
			args: projectCmdArgs{
				serverBaseURL: mockServerURL,
				authToken:     "some-auth-token",
				projectID:     "some-project-id",
				identifier:    "some-identifier",
				stage:         "some-stage",
			},
			mockResponses: []mockResponses{
				{
					path:       "/v1/projects/some-project-id/api/some-identifier",
					statusCode: 400,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.setArgs()
			for _, mockResponse := range tt.mockResponses {
				mockResponse.mockJSONHTTPResponse(t, http.MethodDelete)
			}
			if err := retireCmdRunE(&cobra.Command{}, []string{}); (err != nil) != tt.wantErr {
				t.Errorf("retireCmdRunE() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

//nolint:dupl // ignore dupl linter error for testing
func Test_validateCmdRunE(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	tests := []struct {
		name          string
		args          projectCmdArgs
		mockResponses []mockResponses
		wantErr       bool
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
			name: "status code 400 - no error",
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
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.setArgs()
			for _, mockResponse := range tt.mockResponses {
				mockResponse.mockJSONHTTPResponse(t, http.MethodPost)
			}
			if err := validateCmdRunE(&cobra.Command{}, []string{}); (err != nil) != tt.wantErr {
				t.Errorf("validateCmdRunE() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
