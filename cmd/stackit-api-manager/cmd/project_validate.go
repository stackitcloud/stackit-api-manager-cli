package cmd

import (
	"context"
	"strings"

	apiManager "github.com/stackitcloud/stackit-api-manager-cli/pkg/stackit_api_manager/client"

	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-api-manager-cli/pkg/stackit_api_manager/util"
)

const messageValidateSuccess = "OpenAPI specification validated successfully"

type validateResponse struct {
	Identifier string `json:"identifier"`
	ProjectID  string `json:"projectId"`
	Stage      string `json:"stage"`
}

func (r validateResponse) successMessage() string {
	return messageValidateSuccess
}

var validateCmd = &cobra.Command{ //nolint:gochecknoglobals // CLI command
	Use:           "validate",
	Short:         "Validate an OpenAPI Spec for a Stackit API Gateway project",
	RunE:          validateCmdRunE,
	SilenceErrors: true,
	SilenceUsage:  true,
}

func validateCmdRunE(cmd *cobra.Command, args []string) error {
	c := newAPIClient()

	base64Encoded, err := util.EncodeBase64File(openAPISpecFilePath)
	if err != nil {
		cmd.Print(err)
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

	_, httpResp, err := c.APIManagerServiceApi.APIManagerServicePublishValidate(
		ctx,
		projectID,
		identifier,
	).PublishValidateRequest(req).Execute()
	if err != nil && httpResp == nil {
		cmd.Print(err)
		return err
	}
	defer httpResp.Body.Close()

	if err != nil {
		return printErrorCLIResponse(cmd, httpResp)
	}

	validateResponse := &validateResponse{
		Identifier: identifier,
		ProjectID:  projectID,
		Stage:      stage,
	}

	return printSuccessCLIResponse(cmd, httpResp.StatusCode, validateResponse)
}
