package cmd

import (
	"context"
	"fmt"
	"strings"

	apiManager "github.com/stackitcloud/stackit-api-manager-cli/pkg/stackit_api_manager/client"

	"github.com/spf13/cobra"
)

var (
	version string
)

const messageRetireSuccess = "API with identifier %s retired successfully"
const messageRetireVersionSuccess = "API with identifier version %s of %s retired successfully"

const messageHumanReadableRetireSuccess = "API with identifier: %q retired successfully for project: %q\n"
const messageHumanReadableRetireVersionSuccess = "API with identifier %q version: %q retired successfully for project: %q\n"

type retireResponse struct {
	Identifier string  `json:"identifier"`
	ProjectID  string  `json:"projectId"`
	Version    *string `json:"version,omitempty"`
}

func (r retireResponse) HumanReadableMessage() string {
	if r.Version != nil {
		return fmt.Sprintf(messageHumanReadableRetireVersionSuccess, r.Identifier, *r.Version, r.ProjectID)
	}
	return fmt.Sprintf(messageHumanReadableRetireSuccess, r.Identifier, r.ProjectID)
}

func (r retireResponse) successMessage() string {
	if r.Version != nil {
		return fmt.Sprintf(messageRetireVersionSuccess, *r.Version, r.Identifier)
	}
	return fmt.Sprintf(messageRetireSuccess, r.Identifier)
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
