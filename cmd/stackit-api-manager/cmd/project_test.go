package cmd

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-api-manager-cli/pkg/stackit_api_manager/client"
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

func (m *mockResponses) mockJSONHTTPResponse(t *testing.T) {
	jsonResponse, err := httpmock.NewJsonResponder(m.statusCode, m.body)
	if err != nil {
		t.Error(err)
	}
	httpmock.RegisterResponder("POST", fmt.Sprintf("%s%s", mockServerURL, m.path), jsonResponse)
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

func stringPtn(value string) *string { return &value }
func intPtn(value int) *int          { return &value }

func Test_newAPIClient(t *testing.T) {
	tests := []struct {
		name string
		want *client.Client
	}{
		{
			name: "success",
			want: client.NewClient(defaultBaseUrl, ""),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newAPIClient(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newAPIClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newMetadata(t *testing.T) {
	testIdentifier := "identifier-test"
	testStage := "stage-test"
	emptyString := ""

	tests := []struct {
		name       string
		want       *client.Metadata
		identifier string
		stage      string
	}{
		{
			name: "success with empty values",
			want: &client.Metadata{
				Identifier: &emptyString,
				Stage:      &emptyString,
			},
		},
		{
			name: "success with test values",
			want: &client.Metadata{
				Identifier: &testIdentifier,
				Stage:      &testStage,
			},
			identifier: "identifier-test",
			stage:      "stage-test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			identifier = tt.identifier
			stage = tt.stage

			if got := newMetadata(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newMetadata() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
				serverBaseURL:       "http://localhost",
				authToken:           "some-auth-token",
				projectID:           "some-project-id",
				identifier:          "some-identifier",
				stage:               "some-stage",
				openAPISpecFilePath: "../../../pkg/stackit_api_manager/client/testdata/test.json",
			},
			mockResponses: []mockResponses{
				{
					path:       "/v1/projects/some-project-id/publish",
					statusCode: 200,
					body: client.ProjectPublishResponse{
						Code:    intPtn(200),
						Message: stringPtn("Success"),
					},
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
			name: "error: status code 400",
			args: projectCmdArgs{
				serverBaseURL:       "http://localhost",
				authToken:           "some-auth-token",
				projectID:           "some-project-id",
				identifier:          "some-identifier",
				stage:               "some-stage",
				openAPISpecFilePath: "../../../pkg/stackit_api_manager/client/testdata/test.json",
			},
			mockResponses: []mockResponses{
				{
					path:       "/v1/projects/some-project-id/publish",
					statusCode: 400,
					body: client.ProjectPublishResponse{
						Code:    intPtn(400),
						Message: stringPtn("Not Found"),
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.setArgs()
			for _, mockResponse := range tt.mockResponses {
				mockResponse.mockJSONHTTPResponse(t)
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
				serverBaseURL: "http://localhost",
				authToken:     "some-auth-token",
				projectID:     "some-project-id",
				identifier:    "some-identifier",
				stage:         "some-stage",
			},
			mockResponses: []mockResponses{
				{
					path:       "/v1/projects/some-project-id/retire",
					statusCode: 200,
					body: client.ProjectRetireResponse{
						Code:    intPtn(200),
						Message: stringPtn("Success"),
					},
				},
			},
			wantErr: false,
		},
		{
			name: "error: status code 400",
			args: projectCmdArgs{
				serverBaseURL: "http://localhost",
				authToken:     "some-auth-token",
				projectID:     "some-project-id",
				identifier:    "some-identifier",
				stage:         "some-stage",
			},
			mockResponses: []mockResponses{
				{
					path:       "/v1/projects/some-project-id/retire",
					statusCode: 400,
					body: client.ProjectRetireResponse{
						Code:    intPtn(400),
						Message: stringPtn("Not Found"),
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.setArgs()
			for _, mockResponse := range tt.mockResponses {
				mockResponse.mockJSONHTTPResponse(t)
			}
			if err := retireCmdRunE(&cobra.Command{}, []string{}); (err != nil) != tt.wantErr {
				t.Errorf("retireCmdRunE() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
