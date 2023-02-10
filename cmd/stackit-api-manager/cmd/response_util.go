package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

var (
	errNilCmdResponse         = fmt.Errorf("invalid nil cmdResponse")
	errUnknownCmdResponseType = fmt.Errorf("unknown cmdResponse type")
)

// gatewayResponse contains the HTTP status code and error message
// which are returned by the gateway in case of an error
type gatewayResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// CLIResponse is what users of the CLI receive
// for both successful and erroneous requests
type CLIResponse struct {
	Success    bool        `json:"success"`
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Response   interface{} `json:"response,omitempty"`
}

// a response struct (such as publishResponse, retireResponse,..) has
// to implement this interface to be considered a command response
type cmdResponseInterface interface {
	successMessage() string
}

// prints the CLI response for successful requests
func printSuccessCLIResponse(cmd *cobra.Command, statusCode int, cmdResponse cmdResponseInterface) error {
	if cmdResponse == nil {
		return errNilCmdResponse
	}
	if printJSON {
		return printSuccessCLIResponseJSON(cmd, statusCode, cmdResponse)
	}

	return printSuccessCLIResponseHumanReadable(cmd, cmdResponse)
}

// prints the CLI response for successful requests in JSON format
func printSuccessCLIResponseJSON(cmd *cobra.Command, statusCode int, cmdResponse cmdResponseInterface) error {
	CLIResponse := CLIResponse{
		Success:    true,
		StatusCode: statusCode,
		Message:    cmdResponse.successMessage(),
		Response:   cmdResponse,
	}
	jsonCLIResponse, err := json.Marshal(CLIResponse)
	if err != nil {
		return fmt.Errorf("failed to encode CLI response: %w", err)
	}

	cmd.Println(string(jsonCLIResponse))
	return nil
}

// prints the CLI response for successful requests in human-readable format
func printSuccessCLIResponseHumanReadable(cmd *cobra.Command, cmdResponse cmdResponseInterface) error {
	switch r := cmdResponse.(type) {
	case *publishResponse:
		cmd.Printf("API with identifier \"%s\" published successfully for project \"%s\" and stage \"%s\" (API-URL: \"%s\")\n", r.Identifier, r.ProjectID, r.Stage, r.APIURL)
	case *retireResponse:
		cmd.Printf("API with identifier \"%s\" retired successfully for project \"%s\"\n", r.Identifier, r.ProjectID)
	case *validateResponse:
		cmd.Printf("OpenAPI specification for API with identifier \"%s\", project \"%s\" and stage \"%s\" validated successfully\n", r.Identifier, r.ProjectID, r.Stage)
	default:
		return fmt.Errorf("failed to print human readable success response: %w %T", errUnknownCmdResponseType, r)
	}

	return nil
}

// prints the CLI response for unsuccessful requests
func printErrorCLIResponse(cmd *cobra.Command, httpResp *http.Response) error {
	errorMessage, err := retrieveGatewayErrorMessage(httpResp)
	if err != nil {
		return fmt.Errorf("failed to decode gateway response: %w", err)
	}

	if printJSON {
		CLIResponse := CLIResponse{
			Success:    false,
			StatusCode: httpResp.StatusCode,
			Message:    errorMessage,
		}
		jsonCLIResponse, err := json.Marshal(CLIResponse)
		if err != nil {
			return fmt.Errorf("failed to encode CLI response: %w", err)
		}

		cmd.Println(string(jsonCLIResponse))
		return nil
	}

	cmd.Printf("Failed to %s! An error occured with statuscode %d: %s\n", cmd.Use, httpResp.StatusCode, errorMessage)
	return nil
}

// unmarshals the error message from the gateway HTTP response body
func retrieveGatewayErrorMessage(httpResp *http.Response) (string, error) {
	var gatewayResp gatewayResponse
	err := json.NewDecoder(httpResp.Body).Decode(&gatewayResp)
	if err != nil {
		return "", err
	}

	return gatewayResp.Message, nil
}
