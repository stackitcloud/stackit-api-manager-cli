package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-api-manager-cli/pkg/stackit_api_manager/client"
)

//nolint:gochecknoglobals // CLI command
var (
	serverBaseURL       string
	authToken           string
	projectID           string
	identifier          string
	stage               string
	openAPISpecFilePath string
)

const (
	defaultBaseURL = "https://api-manager.api.stackit.cloud"
)

func newAPIClient() *client.Client {
	return client.NewClient(serverBaseURL, authToken)
}

func newMetadata() client.Metadata {
	return client.Metadata{
		Stage: stage,
	}
}

var projectCmd = &cobra.Command{ //nolint:gochecknoglobals // CLI command
	Use:   "project",
	Short: "Manage your Stackit API Gateway project",
}

var publishCmd = &cobra.Command{ //nolint:gochecknoglobals // CLI command
	Use:   "publish",
	Short: "Publish a OpenAPI Spec to a Stackit API Gateway project",
	RunE:  publishCmdRunE,
}

func publishCmdRunE(cmd *cobra.Command, args []string) error {
	c := newAPIClient()

	base64Encoded, err := client.EncodeBase64File(openAPISpecFilePath)
	if err != nil {
		return err
	}

	err = c.ProjectPublish(projectID, identifier, &client.ProjectPublish{
		Metadata: newMetadata(),
		Spec: client.Spec{
			OpenAPI: &client.OpenAPI{
				Base64Encoded: base64Encoded,
			},
		},
	})
	if err != nil {
		return err
	}
	cmd.Println(fmt.Sprintf("API Gateway project %s published successfully with identifier %s", projectID, identifier))
	return nil
}

var retireCmd = &cobra.Command{ //nolint:gochecknoglobals // CLI command
	Use:   "retire",
	Short: "Retire a OpenAPI Spec from a Stackit API Gateway project",
	RunE:  retireCmdRunE,
}

func retireCmdRunE(cmd *cobra.Command, args []string) error {
	c := newAPIClient()
	err := c.ProjectRetire(projectID, identifier, &client.ProjectRetire{
		Metadata: newMetadata(),
	})
	if err != nil {
		return err
	}
	cmd.Println(fmt.Sprintf("API Gateway project %s retired successfully with identifier %s", projectID, identifier))

	return nil
}

func init() { //nolint:gochecknoinits // cobra CLI
	rootCmd.AddCommand(projectCmd)
	projectCmd.AddCommand(publishCmd)
	projectCmd.AddCommand(retireCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// projectsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// projectsCmd flags
	projectCmd.PersistentFlags().StringVarP(&serverBaseURL, "baseURL", "u", defaultBaseURL, "Server base URL")
	projectCmd.MarkPersistentFlagRequired("url")
	projectCmd.PersistentFlags().StringVarP(&authToken, "token", "t", "", "Auth token for the API Manager")
	projectCmd.PersistentFlags().StringVarP(&projectID, "project", "p", "", "Project ID")
	projectCmd.MarkPersistentFlagRequired("project")
	projectCmd.PersistentFlags().StringVarP(&identifier, "identifier", "i", "", "Project Identifier")
	projectCmd.MarkPersistentFlagRequired("identifier")
	projectCmd.PersistentFlags().StringVarP(&stage, "stage", "s", "", "Project Stage")
	projectCmd.MarkPersistentFlagRequired("stage")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:

	// publishCmd flags
	publishCmd.Flags().StringVarP(&openAPISpecFilePath, "oas", "o", "", "OpenAPI Spec file path")
	publishCmd.MarkFlagRequired("oas")
}
