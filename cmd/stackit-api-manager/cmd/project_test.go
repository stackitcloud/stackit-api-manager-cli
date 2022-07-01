package cmd

import (
	"reflect"
	"testing"

	"github.com/stackitcloud/stackit-api-manager-cli/pkg/stackit_api_manager/client"
)

func Test_newAPIClient(t *testing.T) {
	tests := []struct {
		name string
		want *client.Client
	}{
		{
			name: "success",
			want: client.NewClient("", ""),
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
