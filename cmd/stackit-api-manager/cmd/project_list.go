package cmd

import (
	"context"
	"strings"

	apiManager "github.com/stackitcloud/stackit-api-manager-cli/pkg/stackit_api_manager/client"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{ //nolint:gochecknoglobals // CLI command
	Use:   "list",
	Short: "List all API identifiers for a Stackit API Gateway project",
	RunE:  listCmdRunE,
}

func listCmdRunE(cmd *cobra.Command, args []string) error {
	c := newAPIClient()

	if strings.HasPrefix(authToken, "Bearer ") {
		cmd.Printf("Authorization token should have no Bearer prefix")
		return errBadToken
	}
	// add auth token
	ctx := context.WithValue(context.Background(), apiManager.ContextAccessToken, authToken)

	grpcResponse, httpResponse, err := c.APIManagerServiceApi.APIManagerServiceFetchProjectAPIIdentifiers(
		ctx,
		projectID,
	).Execute()
	if err != nil {
		cmd.Printf("Error when calling `APIManagerServiceApi.APIManagerFetchProjectAPIIdentifiers``: %v\n", err)
		cmd.Printf("Full HTTP response: %v\n", httpResponse)
		return err
	}
	defer httpResponse.Body.Close()

	identifiers := grpcResponse.GetIdentifiers()
	if identifiers == nil {
		cmd.Printf("API Gateway project %s has no identifiers", projectID)
		return nil
	}
	cmd.Printf("API Gateway project %s has following identifiers:\n%+s\n", projectID, identifiers)
	return nil
}
