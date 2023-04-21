package cmd

import (
	"context"
	"strings"

	apiManager "github.com/stackitcloud/stackit-api-manager-cli/pkg/stackit_api_manager/client"

	"github.com/spf13/cobra"
)

const messageRetireSuccess = "API retired successfully"

type retireResponse struct {
	Identifier string `json:"identifier"`
	ProjectID  string `json:"projectId"`
}

func (r retireResponse) successMessage() string {
	return messageRetireSuccess
}

var retireCmd = &cobra.Command{ //nolint:gochecknoglobals // CLI command
	Use:          "retire",
	Short:        "Retire a OpenAPI Spec from a Stackit API Gateway project",
	RunE:         retireCmdRunE,
	SilenceUsage: true,
}

func retireCmdRunE(cmd *cobra.Command, args []string) error {
	c := newAPIClient()

	req := apiManager.RetireRequest{}

	if strings.HasPrefix(authToken, "Bearer ") {
		cmd.Print("Authorization token should have no Bearer prefix")
		return errBadToken
	}
	// add auth token
	ctx := context.WithValue(context.Background(), apiManager.ContextAccessToken, authToken)

	_, httpResp, err := c.APIManagerServiceApi.APIManagerServiceRetire(
		ctx,
		projectID,
		identifier,
	).RetireRequest(req).Execute()
	if err != nil && httpResp == nil {
		return err
	}
	defer httpResp.Body.Close()

	if err != nil {
		return printErrorCLIResponse(cmd, httpResp)
	}

	retireResponse := retireResponse{
		Identifier: identifier,
		ProjectID:  projectID,
	}

	return printSuccessCLIResponse(cmd, httpResp.StatusCode, &retireResponse)
}
