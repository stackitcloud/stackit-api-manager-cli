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
	Run: func(cmd *cobra.Command, args []string) {
		c := newAPIClient()

		base64Encoded, err := client.EncodeBase64File(openAPISpecFilePath)
		if err != nil {
			cmd.PrintErr(err)
		}

		resp, _, err := c.ProjectPublish(projectID, &client.ProjectPublish{
			Metadata: newMetadata(),
			Spec: &client.Spec{
				OpenAPI: &client.OpenAPI{
					Base64Encoded: &base64Encoded,
				},
			},
		})
		cmd.Println(resp)
		if err != nil {
			cmd.PrintErr(err)
		}
	},
}

var retireCmd = &cobra.Command{ //nolint:gochecknoglobals // CLI command
	Use:   "retire",
	Short: "Retire a OpenAPI Spec from a Stackit API Gateway project",
	Run: func(cmd *cobra.Command, args []string) {
		c := newAPIClient()
		resp, _, err := c.ProjectRetire(projectID, &client.ProjectRetire{
			Metadata: newMetadata(),
		})
		cmd.Println(resp)
		if err != nil {
			cmd.PrintErr(err)
		}
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
