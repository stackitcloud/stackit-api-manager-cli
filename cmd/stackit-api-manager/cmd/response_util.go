package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

const (
	traceParentHeader = "Traceparent"
)

var (
	errEncodingCLIResponseMessage     = "failed to encode CLI response"
	errDecodingGatewayResponseMessage = "failed to decode gateway response"

	errNilCmdResponse         = fmt.Errorf("invalid nil cmdResponse")
	errUnknownCmdResponseType = fmt.Errorf("unknown cmdResponse type")
	errRequestFailed          = fmt.Errorf("request failed")
)

type traceIDMessage struct {
	traceID string
}

func (msg traceIDMessage) Message() string {
	return fmt.Sprintf("TraceID: %s\n", msg.traceID)
}

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
	TraceID    string      `json:"trace_id,omitempty"`
}

// a response struct (such as publishResponse, retireResponse,..) has
// to implement this interface to be considered a command response
type cmdResponseInterface interface {
	successMessage() string
}

// prints the CLI response for successful requests
func printSuccessCLIResponse(cmd *cobra.Command, resp *http.Response, cmdResponse cmdResponseInterface) error {
	if cmdResponse == nil {
		return errNilCmdResponse
	}
	if printJSON {
		return printSuccessCLIResponseJSON(cmd, resp, cmdResponse)
	}

	return printSuccessCLIResponseHumanReadable(cmd, resp, cmdResponse)
}

// prints the CLI response for successful requests in JSON format
func printSuccessCLIResponseJSON(cmd *cobra.Command, resp *http.Response, cmdResponse cmdResponseInterface) error {
	cliResponse := CLIResponse{
		Success:    true,
		StatusCode: resp.StatusCode,
		Message:    cmdResponse.successMessage(),
		Response:   cmdResponse,
	}

	if traceIDEnabled {
		cliResponse.TraceID = getTraceID(resp)
	}

	jsonCLIResponse, err := json.Marshal(cliResponse)
	if err != nil {
		return fmt.Errorf("%s: %w", errEncodingCLIResponseMessage, err)
	}

	cmd.Println(string(jsonCLIResponse))
	return nil
}

// prints the CLI response for successful requests in human-readable format
func printSuccessCLIResponseHumanReadable(cmd *cobra.Command, resp *http.Response, cmdResponse cmdResponseInterface) error {
	switch r := cmdResponse.(type) {
	case *publishResponse:
		cmd.Printf("API with identifier \"%s\" published successfully for project \"%s\" and stage \"%s\" (API-URL: \"%s\")\n", r.Identifier, r.ProjectID, r.Stage, r.APIURL)
	case *retireResponse:
		cmd.Printf("API with identifier \"%s\" retired successfully for project \"%s\"\n", r.Identifier, r.ProjectID)
	case *validateResponse:
		if r.LinterWarningsCount != "0" && r.LinterWarningsCount != "" {
			cmd.Printf("OpenAPI specification for API with identifier \"%s\", project \"%s\" and stage \"%s\" validated successfully\nOAS linting resulted in %s warnings:\n  %+s\n", r.Identifier, r.ProjectID, r.Stage, r.LinterWarningsCount, strings.Join(r.LinterWarnings, "\n  "))
			break
		}
		cmd.Printf("OpenAPI specification for API with identifier \"%s\", project \"%s\" and stage \"%s\" validated successfully\n", r.Identifier, r.ProjectID, r.Stage)
	case *listResponse:
		cmd.Printf("Project \"%s\" has the following identifiers: %+v\n", r.ProjectID, r.Identifiers)
	case *fetchResponse:
		cmd.Printf("Base64 encoded OpenAPI specification for API with identifier \"%s\" for project \"%s\" and stage \"%s\" (API-URL: \"%s\", Upstream-URL: \"%s\") is: %v\n", r.Identifier, r.ProjectID, r.Stage, r.APIURL, r.UpstreamURL, r.Base64EncodedSpec)
	default:
		return fmt.Errorf("%w %T", errUnknownCmdResponseType, r)
	}

	printTraceID(cmd, resp)

	return nil
}

// prints the CLI response for unsuccessful requests
func printErrorCLIResponse(cmd *cobra.Command, resp *http.Response) error {
	errorMessage, err := retrieveGatewayErrorMessage(resp)
	if err != nil {
		return fmt.Errorf("%s: %w", errDecodingGatewayResponseMessage, err)
	}

	if printJSON {
		cliResponse := CLIResponse{
			Success:    false,
			StatusCode: resp.StatusCode,
			Message:    errorMessage,
		}

		if traceIDEnabled {
			cliResponse.TraceID = getTraceID(resp)
		}

		jsonCLIResponse, err := json.Marshal(cliResponse)
		if err != nil {
			return fmt.Errorf("%s: %w", errEncodingCLIResponseMessage, err)
		}

		cmd.Println(string(jsonCLIResponse))
		return errRequestFailed
	}

	cmd.Printf("Failed to %s! An error occurred with statuscode %d: %s\n", cmd.Use, resp.StatusCode, errorMessage)
	printTraceID(cmd, resp)

	return errRequestFailed
}

// unmarshals the error message from the gateway HTTP response body
func retrieveGatewayErrorMessage(resp *http.Response) (string, error) {
	var gatewayResp gatewayResponse
	err := json.NewDecoder(resp.Body).Decode(&gatewayResp)
	if err != nil {
		return "", err
	}

	return gatewayResp.Message, nil
}

func getTraceID(resp *http.Response) string {
	traceParentValue := resp.Header.Get(traceParentHeader)
	if traceParentValue == "" {
		return ""
	}

	// e.g., of a trace parent: "00-de0ae651ce4c183a3e5d3eb4827c4fc8-43f4db3d9431bc34-01"
	// de0ae651ce4c183a3e5d3eb4827c4fc8 is the trace ID and 43f4db3d9431bc34 is the parent span ID
	traceParentComponents := strings.Split(traceParentValue, "-")
	if len(traceParentComponents) != 4 {
		return ""
	}

	traceParentValue = traceParentComponents[1]

	return traceParentValue
}

func printTraceID(cmd *cobra.Command, resp *http.Response) {
	traceParentValue := getTraceID(resp)

	if traceIDEnabled && traceParentValue != "" {
		cmd.Printf(traceIDMessage{traceID: traceParentValue}.Message())
	}
}
