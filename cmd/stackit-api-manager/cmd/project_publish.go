package cmd

import (
	"context"
	"strings"

	apiManager "github.com/stackitcloud/stackit-api-manager-cli/pkg/stackit_api_manager/client"

	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-api-manager-cli/pkg/stackit_api_manager/util"
)

const messagePublishSuccess = "API published successfully"

type publishResponse struct {
	Identifier string `json:"identifier"`
	ProjectID  string `json:"projectId"`
	Stage      string `json:"stage"`
	APIURL     string `json:"apiUrl"`
}

func (r publishResponse) successMessage() string {
	return messagePublishSuccess
}

var publishCmd = &cobra.Command{ //nolint:gochecknoglobals // CLI command
	Use:          "publish",
	Short:        "Publish a OpenAPI Spec to a Stackit API Gateway project",
	RunE:         publishCmdRunE,
	SilenceUsage: true,
}

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
		IgnoreLintingErrors:   &ignoreLintingErrors,
		IgnoreBreakingChanges: &ignoreBreakingChanges,
	}

	if strings.HasPrefix(authToken, "Bearer ") {
		cmd.Print("Authorization token should have no Bearer prefix")
		return errBadToken
	}
	// add auth token
	ctx := context.WithValue(context.Background(), apiManager.ContextAccessToken, authToken)

	grpcResp, httpResp, err := c.APIManagerServiceApi.APIManagerServicePublish(
		ctx,
		projectID,
		identifier,
	).PublishRequest(req).Execute()
	if err != nil && httpResp == nil {
		return err
	}
	defer httpResp.Body.Close()
	if err != nil {
		return printErrorCLIResponse(cmd, httpResp)
	}

	publishResponse := publishResponse{
		Identifier: identifier,
		ProjectID:  projectID,
		Stage:      stage,
		APIURL:     grpcResp.GetApiUrl(),
	}

	return printSuccessCLIResponse(cmd, httpResp.StatusCode, &publishResponse)
}
