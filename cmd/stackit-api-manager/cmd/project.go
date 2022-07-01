package cmd

import (
	"encoding/json"

	"github.com/hashicorp/go-multierror"
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

func newAPIClient() *client.Client {
	return client.NewClient(serverBaseURL, authToken)
}

func newMetadata() *client.Metadata {
	return &client.Metadata{
		Identifier: &identifier,
		Stage:      &stage,
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

	resp, _, err := c.ProjectPublish(projectID, &client.ProjectPublish{
		Metadata: newMetadata(),
		Spec: &client.Spec{
			OpenAPI: &client.OpenAPI{
				Base64Encoded: &base64Encoded,
			},
		},
	})

	j, errJSON := json.Marshal(*resp)
	if errJSON != nil {
		err = multierror.Append(err, errJSON)
	} else {
		cmd.Println(string(j))
	}

	if err != nil {
		return err
	}

	return nil
}

var retireCmd = &cobra.Command{ //nolint:gochecknoglobals // CLI command
	Use:   "retire",
	Short: "Retire a OpenAPI Spec from a Stackit API Gateway project",
	RunE:  retireCmdRunE,
}

func retireCmdRunE(cmd *cobra.Command, args []string) error {
	c := newAPIClient()
	resp, _, err := c.ProjectRetire(projectID, &client.ProjectRetire{
		Metadata: newMetadata(),
	})

	j, errJSON := json.Marshal(*resp)
	if errJSON != nil {
		err = multierror.Append(err, errJSON)
	} else {
		cmd.Println(string(j))
	}

	if err != nil {
		return err
	}
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
	projectCmd.PersistentFlags().StringVarP(&serverBaseURL, "baseURL", "u", "", "Server base URL like https://example.com")
	projectCmd.MarkPersistentFlagRequired("url")
	projectCmd.PersistentFlags().StringVarP(&authToken, "token", "t", "", "Auth token for the API Manager")
	projectCmd.MarkPersistentFlagRequired("token")
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
