package cmd

import (
	"context"
	"strings"

	apiManager "github.com/stackitcloud/stackit-api-manager-cli/pkg/stackit_api_manager/client"

	"github.com/spf13/cobra"
)

const messageListSuccess = "Listed API identifiers for project successfully"

type listResponse struct {
	Identifiers []string `json:"identifiers"`
	ProjectID   string   `json:"projectId"`
}

func (r listResponse) successMessage() string {
	return messageListSuccess
}

var listCmd = &cobra.Command{ //nolint:gochecknoglobals // CLI command
	Use:          "list",
	Short:        "List all API identifiers for a Stackit API Gateway project",
	RunE:         listCmdRunE,
	SilenceUsage: true,
}

func listCmdRunE(cmd *cobra.Command, args []string) error {
	c := newAPIClient()

	if strings.HasPrefix(authToken, "Bearer ") {
		cmd.Print("Authorization token should have no Bearer prefix")
		return errBadToken
	}
	// add auth token
	ctx := context.WithValue(context.Background(), apiManager.ContextAccessToken, authToken)

	grpcResp, httpResp, err := c.APIManagerServiceApi.APIManagerServiceFetchProjectAPIIdentifiers(
		ctx,
		projectID,
	).Execute()
	if err != nil && httpResp == nil {
		return err
	}
	defer httpResp.Body.Close()

	if err != nil {
		return printErrorCLIResponse(cmd, httpResp)
	}

	listResponse := listResponse{
		Identifiers: grpcResp.GetIdentifiers(),
		ProjectID:   projectID,
	}

	return printSuccessCLIResponse(cmd, httpResp, &listResponse)
}
