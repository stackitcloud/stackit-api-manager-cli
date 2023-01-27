//nolint:dupl // more clear without reusing functionality
package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	apiManager "github.com/stackitcloud/stackit-api-manager-cli/pkg/stackit_api_manager/client"

	"github.com/spf13/cobra"
)

type fetchAPIResponse struct {
	Stage             string `json:"stage"`
	ApiURL            string `json:"api_url"`
	UpstreamURL       string `json:"upstream_url"`
	Base64EncodedSpec string `json:"base64_encoded_spec"`
}

var fetchAPICmd = &cobra.Command{ //nolint:gochecknoglobals // CLI command
	Use:   "fetch",
	Short: "Fetches the OpenAPI Spec and metadata for an existing Stackit API Gateway project API",
	RunE:  fetchAPICmdRunE,
}

func fetchAPICmdRunE(cmd *cobra.Command, args []string) error {
	c := newAPIClient()

	if strings.HasPrefix(authToken, "Bearer ") {
		cmd.Printf("Authorization token should have no Bearer prefix")
		return errBadToken
	}
	// add auth token
	ctx := context.WithValue(context.Background(), apiManager.ContextAccessToken, authToken)

	grpcResponse, httpResponse, err := c.APIManagerServiceApi.APIManagerServiceFetchAPI(
		ctx,
		projectID,
		identifier,
	).Execute()
	if err != nil {
		cmd.Printf("Error when calling `APIManagerServiceApi.APIManagerServiceFetchAPI``: %v\n", err)
		cmd.Printf("Full HTTP response: %v\n", httpResponse)
		return err
	}
	defer httpResponse.Body.Close()

	if grpcResponse == nil {
		return fmt.Errorf("nil grpc response")
	}
	jsonResponse, err := json.Marshal(fetchAPIResponse{
		Stage:             grpcResponse.GetStage(),
		ApiURL:            grpcResponse.GetApiUrl(),
		UpstreamURL:       grpcResponse.GetUpstreamUrl(),
		Base64EncodedSpec: grpcResponse.Spec.OpenApi.GetBase64Encoded(),
	})
	if err != nil {
		return err
	}

	cmd.Printf("Successfully fetched API for API Gateway project %s with identifier %s\n%v\n",
		projectID,
		identifier,
		string(jsonResponse),
	)
	return nil
}
