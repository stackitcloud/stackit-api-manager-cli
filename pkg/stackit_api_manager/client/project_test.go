package client

import (
	"net/http"
	"testing"
)

func TestEncodeBase64File(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				file: "./testdata/test.json",
			},
			want:    "eyAiaGVsbG8iOiAiZHVkZXR0ZSIgfQo=",
			wantErr: false,
		},
		{
			name: "file not found",
			args: args{
				file: "./testdata/test-not-found.json",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EncodeBase64File(tt.args.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("EncodeBase64File() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("EncodeBase64File() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_ProjectPublish(t *testing.T) {
	type args struct {
		projectID         string
		projectIdentifier string
		projectPublish    *ProjectPublish
	}
	tests := []struct {
		name          string
		args          args
		mockResponses []mockResponses
		wantErr       bool
	}{
		{
			name: "success",
			args: args{
				projectID:         "some-project-id",
				projectIdentifier: "some-identifier",
				projectPublish: &ProjectPublish{
					Metadata: Metadata{
						Identifier: "some-identifier",
						Stage:      "some-stage",
					},
					Spec: Spec{
						OpenAPI: &OpenAPI{
							Base64Encoded: "aGVsbG86IGR1ZGV0dGUK",
						},
					},
				},
			},
			mockResponses: []mockResponses{
				{
					path:       "/v1/projects/some-project-id/api/some-identifier",
					statusCode: 200,
				},
			},
		},
		{
			name: "invalid http response",
			args: args{
				projectID:      "some-project-id",
				projectPublish: &ProjectPublish{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := mockClient(t)
			for _, mockResponse := range tt.mockResponses {
				mockResponse.mockJSONHTTPResponse(t, http.MethodPost)
			}
			err := c.ProjectPublish(tt.args.projectID, tt.args.projectIdentifier, tt.args.projectPublish)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.ProjectPublish() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestClient_ProjectRetire(t *testing.T) {
	type args struct {
		projectID         string
		projectIdentifier string
		projectRetire     *ProjectRetire
	}
	tests := []struct {
		name          string
		args          args
		mockResponses []mockResponses
		wantErr       bool
	}{
		{
			name: "success",
			args: args{
				projectID:         "some-project-id",
				projectIdentifier: "some-identifier",
				projectRetire: &ProjectRetire{
					Metadata: Metadata{
						Identifier: "some-identifier",
						Stage:      "some-stage",
					},
				},
			},
			mockResponses: []mockResponses{
				{
					path:       "/v1/projects/some-project-id/api/some-identifier",
					statusCode: 200,
				},
			},
		},
		{
			name: "invalid http response",
			args: args{
				projectID:     "some-project-id",
				projectRetire: &ProjectRetire{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := mockClient(t)
			for _, mockResponse := range tt.mockResponses {
				mockResponse.mockJSONHTTPResponse(t, http.MethodDelete)
			}
			err := c.ProjectRetire(tt.args.projectID, tt.args.projectIdentifier, tt.args.projectRetire)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.ProjectRetire() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
