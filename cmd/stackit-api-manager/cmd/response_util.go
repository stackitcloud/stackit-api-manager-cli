package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
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

// prints the CLI response in JSON format for successful requests
func printSuccessCLIResponseJSON(cmd *cobra.Command, statusCode int, cmdResponse cmdResponseInterface) error {
	if cmdResponse == nil {
		return fmt.Errorf("invalid nil cmdResponse")
	}
	CLIResponse := CLIResponse{
		Success:    true,
		StatusCode: statusCode,
		Message:    cmdResponse.successMessage(),
		Response:   cmdResponse,
	}
	jsonCLIResponse, err := json.Marshal(CLIResponse)
	if err != nil {
		return fmt.Errorf("failed to encode publish CLI response: %w", err)
	}

	cmd.Println(string(jsonCLIResponse))
	return nil
}

// prints the CLI response in JSON format for unsuccessful requests
func printErrorCLIResponseJSON(cmd *cobra.Command, httpResp *http.Response) error {
	errorMessage, err := retrieveGatewayErrorMessage(httpResp)
	if err != nil {
		return fmt.Errorf("failed to decode gateway response: %w", err)
	}

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

// unmarshals the error message from the gateway HTTP response body
func retrieveGatewayErrorMessage(httpResp *http.Response) (string, error) {
	var gatewayResp gatewayResponse
	err := json.NewDecoder(httpResp.Body).Decode(&gatewayResp)
	if err != nil {
		return "", err
	}

	return gatewayResp.Message, nil
}
