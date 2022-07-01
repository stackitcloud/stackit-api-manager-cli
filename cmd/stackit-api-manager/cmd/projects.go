package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-api-manager-cli/pkg/stackit_api_manager/client"
)

//nolint:gochecknoglobals // CLI command
var (
	serverURL           string
	authToken           string
	projectID           string
	identifier          string
	stage               string
	openAPISpecFilePath string
)

func newAPIClient() *client.Client {
	return client.NewClient(serverURL, authToken)
}

func newMetadata() *client.Metadata {
	return &client.Metadata{
		Identifier: &identifier,
		Stage:      &stage,
	}
}

var projectsCmd = &cobra.Command{ //nolint:gochecknoglobals // CLI command
	Use:   "projects",
	Short: "Manage your Stackit API Gateway projects",
}

var publishCmd = &cobra.Command{ //nolint:gochecknoglobals // CLI command
	Use:   "publish",
	Short: "Publish a OpenAPI Spec to a Stackit API Gateway project",
	RunE: func(cmd *cobra.Command, args []string) error {
		c := newAPIClient()

		base64Encoded, err := client.EncodeBase64File(openAPISpecFilePath)
		if err != nil {
			return err
		}

		resp, _, err := c.ProjectPublish(projectID, &client.ProjectPublish{
			Metadata: newMetadata(),
			Spec: &client.Spec{
				OpenAPI: &client.OpenAPI{
					Base64Encoded: &base64Encoded,
				},
			},
		})
		if err != nil {
			return err
		}
		cmd.Println(resp)

		return nil
	},
}

var retireCmd = &cobra.Command{ //nolint:gochecknoglobals // CLI command
	Use:   "retire",
	Short: "Retire a OpenAPI Spec from a Stackit API Gateway project",
	RunE: func(cmd *cobra.Command, args []string) error {
		c := newAPIClient()
		resp, _, err := c.ProjectRetire(projectID, &client.ProjectRetire{
			Metadata: newMetadata(),
		})
		if err != nil {
			return err
		}
		cmd.Println(resp)
		return nil
	},
}

func init() { //nolint:gochecknoinits // cobra CLI
	rootCmd.AddCommand(projectsCmd)
	projectsCmd.AddCommand(publishCmd)
	projectsCmd.AddCommand(retireCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// projectsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// projectsCmd flags
	projectsCmd.PersistentFlags().StringVarP(&serverURL, "url", "u", "", "Server base URL like https://example.com")
	projectsCmd.MarkPersistentFlagRequired("url")
	projectsCmd.PersistentFlags().StringVarP(&authToken, "token", "t", "", "Auth token for the API Manager")
	projectsCmd.MarkPersistentFlagRequired("token")
	projectsCmd.PersistentFlags().StringVarP(&projectID, "project", "p", "", "Project ID")
	projectsCmd.MarkPersistentFlagRequired("project")
	projectsCmd.PersistentFlags().StringVarP(&identifier, "identifier", "i", "", "Project Identifier")
	projectsCmd.MarkPersistentFlagRequired("identifier")
	projectsCmd.PersistentFlags().StringVarP(&stage, "stage", "s", "", "Project Stage")
	projectsCmd.MarkPersistentFlagRequired("stage")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:

	// publishCmd flags
	publishCmd.Flags().StringVarP(&openAPISpecFilePath, "oas", "o", "", "OpenAPI Spec file path")
	publishCmd.MarkFlagRequired("oas")
}
