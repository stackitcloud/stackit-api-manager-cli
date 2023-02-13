package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

const (
	validIdentifier      = "api-id"
	validIdentifierOther = "api-id-other"
	validProjectID       = "project-id"
	validStage           = "stage"
	validAPIURL          = "test.url/fun"
	validUpstreamURL     = "upstream.url/fun"
	validBase64Spec      = "e2FiY30="
	testErrorMessage     = "error message"
	testStatusCode       = 322
)

var (
	validIdentifiersList     = []string{validIdentifier, validIdentifierOther}
	validIdentifiersListJSON = `["` + strings.Join(validIdentifiersList, `","`) + `"]`

	validGatewayResponseBody   = fmt.Sprintf(`{"Status": %d, "Message": "%s"}`, testStatusCode, testErrorMessage)
	invalidGatewayResponseBody = fmt.Sprintf(`{"Status": "%d", "Message": "%s"}`, testStatusCode, testErrorMessage)

	publishSuccessMessageHumanReadable = fmt.Sprintf(`API with identifier "%s" published successfully for project "%s" and stage "%s" (API-URL: "%s")`, validIdentifier, validProjectID, validStage, validAPIURL)
	publishSuccessMessageJSON          = fmt.Sprintf(`{"success":true,"statusCode":%d,"message":"API published successfully","response":{"identifier":"%s","projectId":"%s","stage":"%s","apiUrl":"%s"}}`, testStatusCode, validIdentifier, validProjectID, validStage, validAPIURL)

	retireSuccessMessageHumanReadable = fmt.Sprintf(`API with identifier "%s" retired successfully for project "%s"`, validIdentifier, validProjectID)
	retireSuccessMessageJSON          = fmt.Sprintf(`{"success":true,"statusCode":%d,"message":"API retired successfully","response":{"identifier":"%s","projectId":"%s"}}`, testStatusCode, validIdentifier, validProjectID)

	validateSuccessMessageHumanReadable = fmt.Sprintf(`OpenAPI specification for API with identifier "%s", project "%s" and stage "%s" validated successfully`, validIdentifier, validProjectID, validStage)
	validateSuccessMessageJSON          = fmt.Sprintf(`{"success":true,"statusCode":%d,"message":"OpenAPI specification validated successfully","response":{"identifier":"%s","projectId":"%s","stage":"%s"}}`, testStatusCode, validIdentifier, validProjectID, validStage)

	listSuccessMessageHumanReadable = fmt.Sprintf(`Project "%s" has the following identifiers: %+v`, validProjectID, validIdentifiersList)
	listSuccessMessageJSON          = fmt.Sprintf(`{"success":true,"statusCode":%d,"message":"Listed API identifiers for project successfully","response":{"identifiers":%s,"projectId":"%s"}}`, testStatusCode, validIdentifiersListJSON, validProjectID)

	fetchSuccessMessageHumanReadable = fmt.Sprintf(`Base64 encoded OpenAPI specification for API with identifier "%s" for project "%s" and stage "%s" (API-URL: "%s", Upstream-URL: "%s") is: %v`,
		validIdentifier, validProjectID, validStage, validAPIURL, validUpstreamURL, validBase64Spec)
	fetchSuccessMessageJSON = fmt.Sprintf(`{"success":true,"statusCode":%d,"message":"Fetched API successfully","response":{"identifier":"%s","projectId":"%s","stage":"%s","apiUrl":"%s","upstreamUrl":"%s","base64EncodedSpec":"%s"}}`,
		testStatusCode, validIdentifier, validProjectID, validStage, validAPIURL, validUpstreamURL, validBase64Spec)
)

type unknownCmdResponse struct{}

func (r unknownCmdResponse) successMessage() string {
	return "success"
}

func Test_retrieveGatewayErrorMessage(t *testing.T) {
	tests := []struct {
		name            string
		gatewayResponse string
		httpStatusCode  int
		want            string
		wantErr         bool
	}{
		{
			name:            "retrieve message from HTTP response successful",
			gatewayResponse: `{"Status": 123, "Message": "error message"}`,
			want:            "error message",
			wantErr:         false,
		},
		{
			name:            "message contains quotation marks - fails",
			gatewayResponse: `{"Status": 123, "Message": "error "message""}`,
			wantErr:         true,
		},
		{
			name:            "status encoded as string returns error",
			gatewayResponse: `{"Status": "123", "Message": "error message"}`,
			wantErr:         true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			httpResp := &http.Response{
				Body:       io.NopCloser(bytes.NewBufferString(tt.gatewayResponse)),
				StatusCode: tt.httpStatusCode,
			}

			got, err := retrieveGatewayErrorMessage(httpResp)

			if (err != nil) != tt.wantErr {
				t.Errorf("retrieveGatewayErrorMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("retrieveGatewayErrorMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_printSuccessCLIResponse(t *testing.T) {
	type args struct {
		cmd         *cobra.Command
		statusCode  int
		printJSON   bool
		cmdResponse cmdResponseInterface
	}
	tests := []struct {
		name      string
		args      args
		wantPrint string
		wantErr   error
	}{
		{
			name: "nil cmd response returns error",
			args: args{
				cmd:         publishCmd,
				statusCode:  testStatusCode,
				cmdResponse: nil,
			},
			wantErr: errNilCmdResponse,
		},
		{
			name: "successful request returns no error and prints JSON when printJSON flag is true",
			args: args{
				cmd:        retireCmd,
				statusCode: testStatusCode,
				printJSON:  true,
				cmdResponse: &retireResponse{
					Identifier: validIdentifier,
					ProjectID:  validProjectID,
				},
			},
			wantPrint: retireSuccessMessageJSON,
			wantErr:   nil,
		},
		{
			name: "successful request returns no error and prints human-readable response when printJSON flag is false",
			args: args{
				cmd:        retireCmd,
				statusCode: testStatusCode,
				printJSON:  false,
				cmdResponse: &retireResponse{
					Identifier: validIdentifier,
					ProjectID:  validProjectID,
				},
			},
			wantPrint: retireSuccessMessageHumanReadable,
			wantErr:   nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			outBuff := bytes.NewBuffer(nil)
			tt.args.cmd.SetOut(outBuff)

			printJSON = tt.args.printJSON

			gotErr := printSuccessCLIResponse(tt.args.cmd, tt.args.statusCode, tt.args.cmdResponse)
			if !errors.Is(gotErr, tt.wantErr) {
				t.Errorf("printSuccessCLIResponse() got error = %v, want %v", gotErr, tt.wantErr)
			}

			gotPrintBytes, err := io.ReadAll(outBuff)
			if err != nil {
				t.Error(err)
			}
			gotPrint := strings.TrimRight(string(gotPrintBytes), "\n")

			if gotErr == nil && gotPrint != tt.wantPrint {
				t.Errorf("printSuccessCLIResponse() expected message to be\n%v\nbut got\n%v", tt.wantPrint, gotPrint)
			}
		})
	}
}

func Test_printSuccessCLIResponseJSON(t *testing.T) {
	type args struct {
		cmd         *cobra.Command
		statusCode  int
		cmdResponse cmdResponseInterface
	}
	tests := []struct {
		name      string
		args      args
		wantPrint string
		wantErr   error
	}{
		{
			name: "successful publish returns no error and prints JSON as expected",
			args: args{
				cmd:        publishCmd,
				statusCode: testStatusCode,
				cmdResponse: &publishResponse{
					Identifier: validIdentifier,
					ProjectID:  validProjectID,
					Stage:      validStage,
					APIURL:     validAPIURL,
				},
			},
			wantPrint: publishSuccessMessageJSON,
			wantErr:   nil,
		},
		{
			name: "successful retire returns no error and prints JSON as expected",
			args: args{
				cmd:        retireCmd,
				statusCode: testStatusCode,
				cmdResponse: &retireResponse{
					Identifier: validIdentifier,
					ProjectID:  validProjectID,
				},
			},
			wantPrint: retireSuccessMessageJSON,
			wantErr:   nil,
		},
		{
			name: "successful validate returns no error and prints JSON as expected",
			args: args{
				cmd:        validateCmd,
				statusCode: testStatusCode,
				cmdResponse: &validateResponse{
					Identifier: validIdentifier,
					ProjectID:  validProjectID,
					Stage:      validStage,
				},
			},
			wantPrint: validateSuccessMessageJSON,
			wantErr:   nil,
		},
		{
			name: "successful list returns no error and prints JSON as expected",
			args: args{
				cmd:        listCmd,
				statusCode: testStatusCode,
				cmdResponse: &listResponse{
					Identifiers: validIdentifiersList,
					ProjectID:   validProjectID,
				},
			},
			wantPrint: listSuccessMessageJSON,
			wantErr:   nil,
		},
		{
			name: "successful fetch returns no error and prints JSON as expected",
			args: args{
				cmd:        fetchCmd,
				statusCode: testStatusCode,
				cmdResponse: &fetchResponse{
					Identifier:        validIdentifier,
					ProjectID:         validProjectID,
					Stage:             validStage,
					APIURL:            validAPIURL,
					UpstreamURL:       validUpstreamURL,
					Base64EncodedSpec: validBase64Spec,
				},
			},
			wantPrint: fetchSuccessMessageJSON,
			wantErr:   nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			outBuff := bytes.NewBuffer(nil)
			tt.args.cmd.SetOut(outBuff)

			gotErr := printSuccessCLIResponseJSON(tt.args.cmd, tt.args.statusCode, tt.args.cmdResponse)
			if !errors.Is(gotErr, tt.wantErr) {
				t.Errorf("printSuccessCLIResponseJSON() got error = %v, want %v", gotErr, tt.wantErr)
			}

			gotPrintBytes, err := io.ReadAll(outBuff)
			if err != nil {
				t.Error(err)
			}
			gotPrint := strings.TrimRight(string(gotPrintBytes), "\n")

			if gotErr == nil && gotPrint != tt.wantPrint {
				t.Errorf("printSuccessCLIResponseJSON() expected message to be\n%v\nbut got\n%v", tt.wantPrint, gotPrint)
			}
		})
	}
}

func Test_printSuccessCLIResponseHumanReadable(t *testing.T) {
	type args struct {
		cmd         *cobra.Command
		statusCode  int
		cmdResponse cmdResponseInterface
	}

	tests := []struct {
		name      string
		args      args
		wantPrint string
		wantErr   bool
	}{
		{
			name: "unknown cmd response type returns error",
			args: args{
				cmd:         publishCmd,
				statusCode:  testStatusCode,
				cmdResponse: &unknownCmdResponse{},
			},
			wantErr: true,
		},
		{
			name: "successful publish returns no error and prints expected human-readable message",
			args: args{
				cmd:        publishCmd,
				statusCode: testStatusCode,
				cmdResponse: &publishResponse{
					Identifier: validIdentifier,
					ProjectID:  validProjectID,
					Stage:      validStage,
					APIURL:     validAPIURL,
				},
			},
			wantPrint: publishSuccessMessageHumanReadable,
			wantErr:   false,
		},
		{
			name: "successful retire returns no error and prints expected human-readable message",
			args: args{
				cmd:        retireCmd,
				statusCode: testStatusCode,
				cmdResponse: &retireResponse{
					Identifier: validIdentifier,
					ProjectID:  validProjectID,
				},
			},
			wantPrint: retireSuccessMessageHumanReadable,
			wantErr:   false,
		},
		{
			name: "successful validate returns no error and prints expected human-readable message",
			args: args{
				cmd:        validateCmd,
				statusCode: testStatusCode,
				cmdResponse: &validateResponse{
					Identifier: validIdentifier,
					ProjectID:  validProjectID,
					Stage:      validStage,
				},
			},
			wantPrint: validateSuccessMessageHumanReadable,
			wantErr:   false,
		},
		{
			name: "successful list returns no error and prints expected human-readable message",
			args: args{
				cmd:        listCmd,
				statusCode: testStatusCode,
				cmdResponse: &listResponse{
					Identifiers: validIdentifiersList,
					ProjectID:   validProjectID,
				},
			},
			wantPrint: listSuccessMessageHumanReadable,
			wantErr:   false,
		},
		{
			name: "successful fetch returns no error and prints expected human-readable message",
			args: args{
				cmd:        fetchCmd,
				statusCode: testStatusCode,
				cmdResponse: &fetchResponse{
					Identifier:        validIdentifier,
					ProjectID:         validProjectID,
					Stage:             validStage,
					APIURL:            validAPIURL,
					UpstreamURL:       validUpstreamURL,
					Base64EncodedSpec: validBase64Spec,
				},
			},
			wantPrint: fetchSuccessMessageHumanReadable,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			outBuff := bytes.NewBuffer(nil)
			tt.args.cmd.SetOut(outBuff)

			gotErr := printSuccessCLIResponseHumanReadable(tt.args.cmd, tt.args.cmdResponse)
			if (gotErr != nil) != tt.wantErr {
				t.Errorf("printSuccessCLIResponseHumanReadable() got error = %v, want %v", gotErr, tt.wantErr)
			}

			gotPrintBytes, err := io.ReadAll(outBuff)
			if err != nil {
				t.Error(err)
			}
			gotPrint := strings.TrimRight(string(gotPrintBytes), "\n")

			if gotErr == nil && gotPrint != tt.wantPrint {
				t.Errorf("printSuccessCLIResponseHumanReadable() expected message to be\n%v\nbut got\n%v", tt.wantPrint, gotPrint)
			}
		})
	}
}

func Test_printErrorCLIResponse(t *testing.T) {
	type args struct {
		cmd                 *cobra.Command
		statusCode          int
		printJSON           bool
		gatewayResponseBody string
	}

	tests := []struct {
		name      string
		args      args
		wantPrint string
		wantErr   bool
	}{
		{
			name: "failed publish returns no error and prints JSON as expected",
			args: args{
				cmd:                 publishCmd,
				statusCode:          testStatusCode,
				printJSON:           true,
				gatewayResponseBody: validGatewayResponseBody,
			},
			wantErr: false,
		},
		{
			name: "failed publish returns no error and prints human-readable as expected",
			args: args{
				cmd:                 publishCmd,
				statusCode:          testStatusCode,
				printJSON:           false,
				gatewayResponseBody: validGatewayResponseBody,
			},
			wantErr: false,
		},
		{
			name: "failed publish with invalid gateway response (string statuscode) returns error",
			args: args{
				cmd:                 publishCmd,
				statusCode:          testStatusCode,
				gatewayResponseBody: invalidGatewayResponseBody,
			},
			wantErr: true,
		},
		{
			name: "failed retire returns no error and prints JSON as expected",
			args: args{
				cmd:                 retireCmd,
				statusCode:          testStatusCode,
				printJSON:           true,
				gatewayResponseBody: validGatewayResponseBody,
			},
			wantErr: false,
		},
		{
			name: "failed retire returns no error and prints human-readable as expected",
			args: args{
				cmd:                 retireCmd,
				statusCode:          testStatusCode,
				printJSON:           false,
				gatewayResponseBody: validGatewayResponseBody,
			},
			wantErr: false,
		},
		{
			name: "failed retire with invalid gateway response (string statuscode) returns error",
			args: args{
				cmd:                 retireCmd,
				statusCode:          testStatusCode,
				gatewayResponseBody: invalidGatewayResponseBody,
			},
			wantErr: true,
		},
		{
			name: "failed validate returns no error and prints JSON as expected",
			args: args{
				cmd:                 validateCmd,
				statusCode:          testStatusCode,
				printJSON:           true,
				gatewayResponseBody: validGatewayResponseBody,
			},
			wantErr: false,
		},
		{
			name: "failed validate returns no error and prints human-readable as expected",
			args: args{
				cmd:                 validateCmd,
				statusCode:          testStatusCode,
				printJSON:           false,
				gatewayResponseBody: validGatewayResponseBody,
			},
			wantErr: false,
		},
		{
			name: "failed validate with invalid gateway response (string statuscode) returns error",
			args: args{
				cmd:                 validateCmd,
				statusCode:          testStatusCode,
				gatewayResponseBody: invalidGatewayResponseBody,
			},
			wantErr: true,
		},
		{
			name: "failed list returns no error and prints JSON as expected",
			args: args{
				cmd:                 listCmd,
				statusCode:          testStatusCode,
				printJSON:           true,
				gatewayResponseBody: validGatewayResponseBody,
			},
			wantErr: false,
		},
		{
			name: "failed list returns no error and prints human-readable as expected",
			args: args{
				cmd:                 listCmd,
				statusCode:          testStatusCode,
				printJSON:           false,
				gatewayResponseBody: validGatewayResponseBody,
			},
			wantErr: false,
		},
		{
			name: "failed list with invalid gateway response (string statuscode) returns error",
			args: args{
				cmd:                 listCmd,
				statusCode:          testStatusCode,
				gatewayResponseBody: invalidGatewayResponseBody,
			},
			wantErr: true,
		},
		{
			name: "failed fetch returns no error and prints JSON as expected",
			args: args{
				cmd:                 fetchCmd,
				statusCode:          testStatusCode,
				printJSON:           true,
				gatewayResponseBody: validGatewayResponseBody,
			},
			wantErr: false,
		},
		{
			name: "failed fetch returns no error and prints human-readable as expected",
			args: args{
				cmd:                 fetchCmd,
				statusCode:          testStatusCode,
				printJSON:           false,
				gatewayResponseBody: validGatewayResponseBody,
			},
			wantErr: false,
		},
		{
			name: "failed fetch with invalid gateway response (string statuscode) returns error",
			args: args{
				cmd:                 fetchCmd,
				statusCode:          testStatusCode,
				gatewayResponseBody: invalidGatewayResponseBody,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			httpResp := &http.Response{
				Body:       io.NopCloser(bytes.NewBufferString(tt.args.gatewayResponseBody)),
				StatusCode: tt.args.statusCode,
			}

			outBuff := bytes.NewBuffer(nil)
			tt.args.cmd.SetOut(outBuff)

			printJSON = tt.args.printJSON

			gotErr := printErrorCLIResponse(tt.args.cmd, httpResp)
			if (gotErr != nil) != tt.wantErr {
				t.Errorf("printErrorCLIResponse() got error = %v, want %v", gotErr, tt.wantErr)
			}

			wantPrint := fmt.Sprintf("Failed to %s! An error occurred with statuscode %d: %s", tt.args.cmd.Use, testStatusCode, testErrorMessage)
			if tt.args.printJSON {
				wantPrint = fmt.Sprintf(`{"success":false,"statusCode":%d,"message":"%s"}`, testStatusCode, testErrorMessage)
			}

			gotPrintBytes, err := io.ReadAll(outBuff)
			if err != nil {
				t.Error(err)
			}
			gotPrint := strings.TrimRight(string(gotPrintBytes), "\n")

			if gotErr == nil && gotPrint != wantPrint {
				t.Errorf("printErrorCLIResponse() expected message to be\n%v\nbut got\n%v", wantPrint, gotPrint)
			}
		})
	}
}
