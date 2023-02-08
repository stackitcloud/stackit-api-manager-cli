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
	Use:   "retire",
	Short: "Retire a OpenAPI Spec from a Stackit API Gateway project",
	RunE:  retireCmdRunE,
}

func retireCmdRunE(cmd *cobra.Command, args []string) error {
	c := newAPIClient()

	req := apiManager.RetireRequest{}

	if strings.HasPrefix(authToken, "Bearer ") {
		cmd.Printf("Authorization token should have no Bearer prefix")
		return errBadToken
	}
	// add auth token
	ctx := context.WithValue(context.Background(), apiManager.ContextAccessToken, authToken)

	_, httpResp, err := c.APIManagerServiceApi.APIManagerServiceRetire(
		ctx,
		projectID,
		identifier,
	).RetireRequest(req).Execute()
	defer httpResp.Body.Close()
	if err != nil {
		err := printErrorCLIResponseJSON(cmd, httpResp)
		if err != nil {
			return err
		}

		return nil
	}

	retireResponse := retireResponse{
		Identifier: identifier,
		ProjectID:  projectID,
	}
	err = printSuccessCLIResponseJSON(cmd, httpResp.StatusCode, &retireResponse)
	if err != nil {
		return err
	}

	return nil
}
