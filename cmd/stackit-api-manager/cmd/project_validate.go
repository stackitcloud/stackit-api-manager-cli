//nolint:dupl // more clear without reusing functionality
package cmd

import (
	"context"
	"strings"

	apiManager "github.com/stackitcloud/stackit-api-manager-cli/pkg/stackit_api_manager/client"

	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-api-manager-cli/pkg/stackit_api_manager/util"
)

var validateCmd = &cobra.Command{ //nolint:gochecknoglobals // CLI command
	Use:   "validate",
	Short: "Validate an OpenAPI Spec for a Stackit API Gateway project",
	RunE:  validateCmdRunE,
}

func validateCmdRunE(cmd *cobra.Command, args []string) error {
	c := newAPIClient()

	base64Encoded, err := util.EncodeBase64File(openAPISpecFilePath)
	if err != nil {
		return err
	}

	req := apiManager.PublishValidateRequest{
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

	_, r, err := c.APIManagerServiceApi.APIManagerServicePublishValidate(
		ctx,
		projectID,
		identifier,
	).PublishValidateRequest(req).Execute()
	if err != nil {
		cmd.Printf("Error when calling `APIManagerServiceApi.APIManagerServicePublishValidate``: %v\n", err)
		cmd.Printf("Full HTTP response: %v\n", r)
		return err
	}
	defer r.Body.Close()
	cmd.Printf("OpenAPI Spec for API Gateway project %s with identifier %s validated successfully\n", projectID, identifier)
	return nil
}
