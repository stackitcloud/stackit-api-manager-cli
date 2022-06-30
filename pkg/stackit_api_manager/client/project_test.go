package client

import (
	"reflect"
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
				file: "./testdata/test.yaml",
			},
			want:    "aGVsbG86IGR1ZGV0dGUK",
			wantErr: false,
		},
		{
			name: "file not found",
			args: args{
				file: "./testdata/test-not-found.yaml",
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
		projectID      string
		projectPublish *ProjectPublish
	}
	tests := []struct {
		name          string
		args          args
		mockResponses []mockResponses
		want          *ProjectPublishResponse
		wantErr       bool
	}{
		{
			name: "success",
			args: args{
				projectID: "some-project-id",
				projectPublish: &ProjectPublish{
					Metadata: &Metadata{
						Identifier: stringPtn("some-identifier"),
						Stage:      stringPtn("some-stage"),
					},
					Spec: &Spec{
						OpenAPI: &OpenAPI{
							Base64Encoded: stringPtn("aGVsbG86IGR1ZGV0dGUK"),
						},
					},
				},
			},
			mockResponses: []mockResponses{
				{
					path: "/v1/projects/some-project-id/publish",
					body: ProjectPublishResponse{
						Code:    intPtn(200),
						Message: stringPtn("Success"),
					},
				},
			},
			want: &ProjectPublishResponse{
				Code:    intPtn(200),
				Message: stringPtn("Success"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := mockClient(t)

			for _, mockResponse := range tt.mockResponses {
				mockJSONHTTPResponse(t, mockResponse.path, mockResponse.body)
			}

			got, _, err := c.ProjectPublish(tt.args.projectID, tt.args.projectPublish)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.ProjectPublish() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.ProjectPublish() got = %v, want %v", got, tt.want)
			}
		})
	}
}
