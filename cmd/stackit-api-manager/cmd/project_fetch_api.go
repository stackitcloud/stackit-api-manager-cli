package cmd

import (
	"context"
	"strings"

	apiManager "github.com/stackitcloud/stackit-api-manager-cli/pkg/stackit_api_manager/client"

	"github.com/spf13/cobra"
)

const messageFetchSuccess = "Fetched API successfully"

type fetchResponse struct {
	Identifier        string `json:"identifier"`
	ProjectID         string `json:"projectId"`
	Stage             string `json:"stage"`
	APIURL            string `json:"apiUrl"`
	UpstreamURL       string `json:"upstreamUrl"`
	Base64EncodedSpec string `json:"base64EncodedSpec"`
}

func (r fetchResponse) successMessage() string {
	return messageFetchSuccess
}

var fetchCmd = &cobra.Command{ //nolint:gochecknoglobals // CLI command
	Use:   "fetch",
	Short: "Fetches the OpenAPI Spec and metadata for an existing Stackit API Gateway project API",
	RunE:  fetchCmdRunE,
}

func fetchCmdRunE(cmd *cobra.Command, args []string) error {
	c := newAPIClient()

	if strings.HasPrefix(authToken, "Bearer ") {
		cmd.Printf("Authorization token should have no Bearer prefix")
		return errBadToken
	}
	// add auth token
	ctx := context.WithValue(context.Background(), apiManager.ContextAccessToken, authToken)

	grpcResp, httpResp, err := c.APIManagerServiceApi.APIManagerServiceFetchAPI(
		ctx,
		projectID,
		identifier,
	).Execute()
	if err != nil && httpResp == nil {
		return err
	}
	defer httpResp.Body.Close()
	if err != nil {
		return printErrorCLIResponse(cmd, httpResp)
	}

	fetchResponse := fetchResponse{
		Identifier:        identifier,
		ProjectID:         projectID,
		Stage:             grpcResp.GetStage(),
		APIURL:            grpcResp.GetApiUrl(),
		UpstreamURL:       grpcResp.GetUpstreamUrl(),
		Base64EncodedSpec: grpcResp.Spec.OpenApi.GetBase64Encoded(),
	}

	return printSuccessCLIResponse(cmd, httpResp.StatusCode, &fetchResponse)
}
