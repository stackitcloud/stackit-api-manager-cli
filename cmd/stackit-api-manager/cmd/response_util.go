package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

const (
	traceParentHeader        = "Traceparent"
	numberOfRootCMDsToIgnore = 2
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

	if printTraceID {
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
//
//nolint:cyclop // barely above max. complexity, will be refactored soon
func printSuccessCLIResponseHumanReadable(cmd *cobra.Command, resp *http.Response, cmdResponse cmdResponseInterface) error {
	switch r := cmdResponse.(type) {
	case *publishResponse:
		if r.LinterWarningsCount != "0" && r.LinterWarningsCount != "" {
			cmd.Printf("OpenAPI specification for API with identifier \"%s\", project \"%s\" and stage \"%s\" published successfully\nOAS linting resulted in %s warnings:\n  %+s\n", r.Identifier, r.ProjectID, r.Stage, r.LinterWarningsCount, strings.Join(r.LinterWarnings, "\n  "))
			break
		}
		cmd.Printf("API with identifier \"%s\" published successfully for project \"%s\" and stage \"%s\" (API-URL: \"%s\")\n", r.Identifier, r.ProjectID, r.Stage, r.APIURL)
	case *retireResponse:
		cmd.Printf(r.HumanReadableMessage())
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

	printHumanReadableTraceID(cmd, resp)

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

		if printTraceID {
			cliResponse.TraceID = getTraceID(resp)
		}

		jsonCLIResponse, err := json.Marshal(cliResponse)
		if err != nil {
			return fmt.Errorf("%s: %w", errEncodingCLIResponseMessage, err)
		}

		cmd.Println(string(jsonCLIResponse))
		return errRequestFailed
	}

	currentTarget := cmd
	var targets []string = []string{currentTarget.Use}
	for {
		if currentTarget.Parent() == nil {
			break
		}
		currentTarget = currentTarget.Parent()
		targets = append([]string{currentTarget.Use}, targets...)
	}

	if len(targets) > numberOfRootCMDsToIgnore {
		cmd.Printf("Failed to %s! An error occurred with statuscode %d: %s\n", strings.Join(targets[2:], " "), resp.StatusCode, errorMessage)
	} else {
		cmd.Printf("An error occurred with statuscode %d: %s\n", resp.StatusCode, errorMessage)
	}

	printHumanReadableTraceID(cmd, resp)

	return errRequestFailed
}

func printHumanReadableTraceID(cmd *cobra.Command, resp *http.Response) {
	traceParentValue := getTraceID(resp)

	if printTraceID && traceParentValue != "" {
		cmd.Printf(traceIDMessage{traceID: traceParentValue}.Message())
	}
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

	pattern := regexp.MustCompile(`^00-([a-f0-9]{32})-`)

	// Extract the trace ID using the regular expression
	match := pattern.FindStringSubmatch(traceParentValue)
	if len(match) <= 1 {
		return ""
	}

	return match[1]
}
