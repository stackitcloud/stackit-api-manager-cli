package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-api-manager-cli/pkg/stackit_api_manager/client"
)

var (
	serverURL           string
	authToken           string
	projectID           string
	identifier          string
	stage               string
	OpenAPISpecFilePath string
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

// projectsCmd represents the projects command
var projectsCmd = &cobra.Command{
	Use:   "projects",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("projects called")
	},
}

// publishCmd represents the publish command
var publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		c := newAPIClient()

		base64Encoded, err := client.EncodeOpenAPISpecFile(OpenAPISpecFilePath)
		if err != nil {
			log.Fatal(err)
		}

		resp, _, err := c.ProjectPublish(projectID, &client.ProjectPublish{
			Metadata: newMetadata(),
			Spec: &client.Spec{
				OpenAPI: &client.OpenAPI{
					Base64Encoded: &base64Encoded,
				},
			},
		})
		log.Println("Response:", resp)
		if err != nil {
			log.Fatal(err)
		}
	},
}

// retireCmd represents the retire command
var retireCmd = &cobra.Command{
	Use:   "retire",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		c := newAPIClient()
		resp, _, err := c.ProjectRetire(projectID, &client.ProjectRetire{
			Metadata: newMetadata(),
		})
		log.Println("Response:", resp)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
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
	publishCmd.Flags().StringVarP(&OpenAPISpecFilePath, "oas", "o", "", "OpenAPI Spec file path")
	publishCmd.MarkFlagRequired("oas")
}
