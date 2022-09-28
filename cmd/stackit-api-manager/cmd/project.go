package cmd

import (
	"context"

	"github.com/spf13/cobra"
	apiManager "github.com/stackitcloud/stackit-api-manager-cli/pkg/stackit_api_manager/client"
	"github.com/stackitcloud/stackit-api-manager-cli/pkg/stackit_api_manager/util"
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

type supportedToken string

const (
	defaultBaseURL                = "https://api-manager.api.stackit.cloud"
	Bearer         supportedToken = "Bearer"
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

var publishCmd = &cobra.Command{ //nolint:gochecknoglobals // CLI command
	Use:   "publish",
	Short: "Publish a OpenAPI Spec to a Stackit API Gateway project",
	RunE:  publishCmdRunE,
}

func publishCmdRunE(cmd *cobra.Command, args []string) error {
	c := newAPIClient()

	base64Encoded, err := util.EncodeBase64File(openAPISpecFilePath)
	if err != nil {
		return err
	}

	body := *apiManager.NewAPIManagerServicePublishRequest()
	body.Metadata = &apiManager.V1Metadata{Stage: &stage}
	body.Spec = &apiManager.PublishRequestSpec{
		OpenApi: &apiManager.PublishRequestOpenApi{
			Base64Encoded: &base64Encoded,
		},
	}

	// add auth token
	ctx := context.WithValue(context.Background(), apiManager.ContextAccessToken, authToken)

	_, r, err := c.APIManagerServiceApi.APIManagerServicePublish(
		ctx,
		projectID,
		identifier,
	).Body(body).Execute()
	if err != nil {
		cmd.Printf("Error when calling `APIManagerServiceApi.APIManagerServicePublish``: %v\n", err)
		cmd.Printf("Full HTTP response: %v\n", r)
		return err
	}
	defer r.Body.Close()
	cmd.Printf("API Gateway project %s published successfully with identifier %s", projectID, identifier)
	return nil
}

var retireCmd = &cobra.Command{ //nolint:gochecknoglobals // CLI command
	Use:   "retire",
	Short: "Retire a OpenAPI Spec from a Stackit API Gateway project",
	RunE:  retireCmdRunE,
}

func retireCmdRunE(cmd *cobra.Command, args []string) error {
	c := newAPIClient()

	body := *apiManager.NewAPIManagerServiceRetireRequest()
	body.Metadata = &apiManager.V1Metadata{Stage: &stage}

	// add auth token
	ctx := context.WithValue(context.Background(), apiManager.ContextAccessToken, authToken)

	resp, r, err := c.APIManagerServiceApi.APIManagerServiceRetire(ctx, projectID, identifier).Body(body).Execute()
	if err != nil {
		cmd.Printf("Error when calling `APIManagerServiceApi.APIManagerServiceRetire``: %v\n", err)
		cmd.Printf("Full HTTP response: %v\n", r)
		return err
	}
	defer r.Body.Close()
	cmd.Printf("Response from `APIManagerServiceApi.APIManagerServiceRetire`: %v\n", resp)

	return nil
}

func init() {
	rootCmd.AddCommand(projectCmd)
	projectCmd.AddCommand(publishCmd)
	projectCmd.AddCommand(retireCmd)

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
	projectCmd.PersistentFlags().StringVarP(&identifier, "identifier", "i", "", "Project Identifier")
	projectCmd.MarkPersistentFlagRequired("identifier") //nolint:errcheck // cobra flag
	projectCmd.PersistentFlags().StringVarP(&stage, "stage", "s", "", "Project Stage")
	projectCmd.MarkPersistentFlagRequired("stage") //nolint:errcheck // cobra flag

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:

	// publishCmd flags
	publishCmd.Flags().StringVarP(&openAPISpecFilePath, "oas", "o", "", "OpenAPI Spec file path")
	publishCmd.MarkFlagRequired("oas") //nolint:errcheck // cobra flag
}
