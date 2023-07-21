package cmd

import (
	"context"
	"strings"

	apiManager "github.com/stackitcloud/stackit-api-manager-cli/pkg/stackit_api_manager/client"

	"github.com/spf13/cobra"
)

//nolint:gochecknoglobals // CLI command
var (
	version string
)

const messageRetireSuccess = "API retired successfully"

type retireResponse struct {
	Identifier string  `json:"identifier"`
	ProjectID  string  `json:"projectId"`
	Version    *string `json:"version,omitempty"`
}

func (r retireResponse) successMessage() string {
	return messageRetireSuccess
}

func init() {
	// nested target inside retire
	// used for retiring specific versions
	retireCmd.AddCommand(retireVersionCmd)

	// retireVersionCmd flags for the nested version target of the retire operation
	retireVersionCmd.Flags().StringVarP(&version, "version", "v", "", "API Version")
	retireVersionCmd.MarkFlagRequired("v") //nolint:errcheck // cobra flag
	retireVersionCmd.Flags().StringVarP(&identifier, "identifier", "i", "", "API Identifier")
	retireVersionCmd.MarkFlagRequired("identifier") //nolint:errcheck // cobra flag

}

var retireCmd = &cobra.Command{ //nolint:gochecknoglobals // CLI command
	Use:          "retire",
	Short:        "Retire a OpenAPI Spec from a Stackit API Gateway project",
	RunE:         retireCmdRunE,
	SilenceUsage: true,
}

var retireVersionCmd = &cobra.Command{ //nolint:gochecknoglobals // CLI command
	Use:          "version",
	Short:        "API Version to retire",
	RunE:         retireVersionCmdRunE,
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

	return printSuccessCLIResponse(cmd, httpResp, &retireResponse)
}

func retireVersionCmdRunE(cmd *cobra.Command, args []string) error {
	c := newAPIClient()

	req := apiManager.RetireVersionRequest{}

	if strings.HasPrefix(authToken, "Bearer ") {
		cmd.Print("Authorization token should have no Bearer prefix")
		return errBadToken
	}
	// add auth token
	ctx := context.WithValue(context.Background(), apiManager.ContextAccessToken, authToken)

	_, httpResp, err := c.APIManagerServiceApi.APIManagerServiceRetireVersion(
		ctx,
		projectID,
		identifier,
		version,
	).RetireVersionRequest(req).Execute()
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
		Version:    &version,
	}

	return printSuccessCLIResponse(cmd, httpResp, &retireResponse)
}
