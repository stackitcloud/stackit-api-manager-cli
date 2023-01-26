package cmd

import (
	"context"
	"strings"

	apiManager "github.com/stackitcloud/stackit-api-manager-cli/pkg/stackit_api_manager/client"

	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-api-manager-cli/pkg/stackit_api_manager/util"
)

var publishCmd = &cobra.Command{ //nolint:gochecknoglobals // CLI command
	Use:   "publish",
	Short: "Publish a OpenAPI Spec to a Stackit API Gateway project",
	RunE:  publishCmdRunE,
}

//nolint:dupl // more clear without reusing publish functionality
func publishCmdRunE(cmd *cobra.Command, args []string) error {
	c := newAPIClient()

	base64Encoded, err := util.EncodeBase64File(openAPISpecFilePath)
	if err != nil {
		return err
	}

	req := apiManager.PublishRequest{
		Metadata: &apiManager.PublishMetadata{
			Stage: &stage,
		},
		Spec: &apiManager.Spec{
			OpenApi: &apiManager.SpecOpenApi{
				Base64Encoded: &base64Encoded,
			},
		},
	}

	if strings.HasPrefix(authToken, "Bearer ") {
		cmd.Printf("Authorization token should have no Bearer prefix")
		return errBadToken
	}
	// add auth token
	ctx := context.WithValue(context.Background(), apiManager.ContextAccessToken, authToken)

	_, r, err := c.APIManagerServiceApi.APIManagerServicePublish(
		ctx,
		projectID,
		identifier,
	).PublishRequest(req).Execute()
	if err != nil {
		cmd.Printf("Error when calling `APIManagerServiceApi.APIManagerServicePublish``: %v\n", err)
		cmd.Printf("Full HTTP response: %v\n", r)
		return err
	}
	defer r.Body.Close()
	cmd.Printf("API Gateway project %s published successfully with identifier %s\n", projectID, identifier)
	return nil
}
