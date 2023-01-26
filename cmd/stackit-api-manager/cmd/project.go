package cmd

import (
	"fmt"

	apiManager "github.com/stackitcloud/stackit-api-manager-cli/pkg/stackit_api_manager/client"

	"github.com/spf13/cobra"
)

//nolint:gochecknoglobals // CLI command
var (
	serverBaseURL       string
	authToken           string
	projectID           string
	identifier          string
	stage               string
	openAPISpecFilePath string
	printJSON           bool
)

var errBadToken = fmt.Errorf("bad token")

const (
	defaultBaseURL = "https://api-manager.api.stackit.cloud"
)

func newAPIClient() *apiManager.APIClient {
	cfg := apiManager.NewConfiguration()
	cfg.Servers[0].URL = serverBaseURL
	return apiManager.NewAPIClient(cfg)
}

var projectCmd = &cobra.Command{ //nolint:gochecknoglobals // CLI command
	Use:   "project",
	Short: "Manage your Stackit API Gateway project",
}

func init() {
	rootCmd.AddCommand(projectCmd)
	projectCmd.AddCommand(publishCmd)
	projectCmd.AddCommand(retireCmd)
	projectCmd.AddCommand(validateCmd)
	projectCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// projectsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// projectsCmd flags
	projectCmd.PersistentFlags().StringVarP(&serverBaseURL, "baseURL", "u", defaultBaseURL, "Server base URL")
	projectCmd.MarkPersistentFlagRequired("url") //nolint:errcheck // cobra flag
	projectCmd.PersistentFlags().StringVarP(&authToken, "token", "t", "", "Auth token for the API Manager")
	projectCmd.MarkPersistentFlagRequired("token") //nolint:errcheck // cobra flag
	projectCmd.PersistentFlags().StringVarP(&projectID, "project", "p", "", "Project ID")
	projectCmd.MarkPersistentFlagRequired("project") //nolint:errcheck // cobra flag
	projectCmd.PersistentFlags().BoolVar(&printJSON, "json", false, "Print JSON instead of human readable response")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:

	// retireCmd flags
	retireCmd.Flags().StringVarP(&identifier, "identifier", "i", "", "API Identifier")
	retireCmd.MarkFlagRequired("identifier") //nolint:errcheck // cobra flag

	// publishCmd flags
	publishCmd.Flags().StringVarP(&stage, "stage", "s", "", "Project Stage")
	publishCmd.MarkFlagRequired("stage") //nolint:errcheck // cobra flag
	publishCmd.Flags().StringVarP(&openAPISpecFilePath, "oas", "o", "", "OpenAPI Spec file path")
	publishCmd.MarkFlagRequired("oas") //nolint:errcheck // cobra flag
	publishCmd.Flags().StringVarP(&identifier, "identifier", "i", "", "API Identifier")
	publishCmd.MarkFlagRequired("identifier") //nolint:errcheck // cobra flag

	// validateCmd flags
	validateCmd.Flags().StringVarP(&stage, "stage", "s", "", "Project Stage")
	validateCmd.MarkFlagRequired("stage") //nolint:errcheck // cobra flag
	validateCmd.Flags().StringVarP(&openAPISpecFilePath, "oas", "o", "", "OpenAPI Spec file path")
	validateCmd.MarkFlagRequired("oas") //nolint:errcheck // cobra flag
	validateCmd.Flags().StringVarP(&identifier, "identifier", "i", "", "API Identifier")
	validateCmd.MarkFlagRequired("identifier") //nolint:errcheck // cobra flag
}
