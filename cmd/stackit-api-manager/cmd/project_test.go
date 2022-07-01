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
