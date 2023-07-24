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
	validVersion             = "v1"
	validIdentifiersList     = []string{validIdentifier, validIdentifierOther}
	validIdentifiersListJSON = `["` + strings.Join(validIdentifiersList, `","`) + `"]`

	validGatewayResponseBody   = fmt.Sprintf(`{"Status": %d, "Message": "%s"}`, testStatusCode, testErrorMessage)
	invalidGatewayResponseBody = fmt.Sprintf(`{"Status": "%d", "Message": "%s"}`, testStatusCode, testErrorMessage)

	publishSuccessMessageHumanReadable = fmt.Sprintf(`API with identifier "%s" published successfully for project "%s" and stage "%s" (API-URL: "%s")`, validIdentifier, validProjectID, validStage, validAPIURL)
	publishSuccessMessageJSON          = fmt.Sprintf(`{"success":true,"statusCode":%d,"message":"API published successfully","response":{"identifier":"%s","projectId":"%s","stage":"%s","apiUrl":"%s"}}`, testStatusCode, validIdentifier, validProjectID, validStage, validAPIURL)

	retireSuccessMessageHumanReadable = fmt.Sprintf(`API with identifier: "%s" retired successfully for project: "%s"`, validIdentifier, validProjectID)
	retireSuccessMessageJSON          = fmt.Sprintf(`{"success":true,"statusCode":%d,"message":"API %s retired successfully","response":{"identifier":"%s","projectId":"%s"}}`, testStatusCode, validIdentifier, validIdentifier, validProjectID)

	retireVersionSuccessMessageHumanReadable = fmt.Sprintf(`API with identifier "%s" version: "%s" retired successfully for project: "%s"`, validIdentifier, validVersion, validProjectID)
	retireVersionSuccessMessageJSON          = fmt.Sprintf(`{"success":true,"statusCode":%d,"message":"API version %s of %s retired successfully","response":{"identifier":"%s","projectId":"%s","version":"%s"}}`, testStatusCode, validVersion, validIdentifier, validIdentifier, validProjectID, "v1")

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
		resp        *http.Response
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
				cmd: publishCmd,
				resp: &http.Response{
					StatusCode: int(testStatusCode),
				},
				cmdResponse: nil,
			},
			wantErr: errNilCmdResponse,
		},
		{
			name: "successful request returns no error and prints JSON when printJSON flag is true",
			args: args{
				cmd: retireCmd,
				resp: &http.Response{
					StatusCode: int(testStatusCode),
				},
				printJSON: true,
				cmdResponse: &retireResponse{
					Identifier: validIdentifier,
					ProjectID:  validProjectID,
				},
			},
			wantPrint: retireSuccessMessageJSON,
			wantErr:   nil,
		},
		{
			name: "retire version - successful request returns no error and prints JSON when printJSON flag is true",
			args: args{
				cmd: retireVersionCmd,
				resp: &http.Response{
					StatusCode: int(testStatusCode),
				},
				printJSON: true,
				cmdResponse: &retireResponse{
					Identifier: validIdentifier,
					ProjectID:  validProjectID,
					Version:    &validVersion,
				},
			},
			wantPrint: retireVersionSuccessMessageJSON,
			wantErr:   nil,
		},
		{
			name: "successful request returns no error and prints human-readable response when printJSON flag is false",
			args: args{
				cmd: retireCmd,
				resp: &http.Response{
					StatusCode: int(testStatusCode),
				},
				printJSON: false,
				cmdResponse: &retireResponse{
					Identifier: validIdentifier,
					ProjectID:  validProjectID,
				},
			},
			wantPrint: retireSuccessMessageHumanReadable,
			wantErr:   nil,
		},
		{
			name: "retire version - successful request returns no error and prints human-readable response when printJSON flag is false",
			args: args{
				cmd: retireVersionCmd,
				resp: &http.Response{
					StatusCode: int(testStatusCode),
				},
				printJSON: false,
				cmdResponse: &retireResponse{
					Identifier: validIdentifier,
					ProjectID:  validProjectID,
					Version:    &validVersion,
				},
			},
			wantPrint: retireVersionSuccessMessageHumanReadable,
			wantErr:   nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			outBuff := bytes.NewBuffer(nil)
			tt.args.cmd.SetOut(outBuff)

			printJSON = tt.args.printJSON

			gotErr := printSuccessCLIResponse(tt.args.cmd, tt.args.resp, tt.args.cmdResponse)
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
		resp        *http.Response
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
				cmd: publishCmd,
				resp: &http.Response{
					StatusCode: int(testStatusCode),
				},
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
				cmd: retireCmd,
				resp: &http.Response{
					StatusCode: int(testStatusCode),
				},
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
				cmd: validateCmd,
				resp: &http.Response{
					StatusCode: int(testStatusCode),
				},
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
				cmd: listCmd,
				resp: &http.Response{
					StatusCode: int(testStatusCode),
				},
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
				cmd: fetchCmd,
				resp: &http.Response{
					StatusCode: int(testStatusCode),
				},
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

			gotErr := printSuccessCLIResponseJSON(tt.args.cmd, tt.args.resp, tt.args.cmdResponse)
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
		resp        *http.Response
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
				cmd: publishCmd,
				resp: &http.Response{
					StatusCode: int(testStatusCode),
				},
				cmdResponse: &unknownCmdResponse{},
			},
			wantErr: true,
		},
		{
			name: "successful publish returns no error and prints expected human-readable message",
			args: args{
				cmd: publishCmd,
				resp: &http.Response{
					StatusCode: int(testStatusCode),
				},
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
				cmd: retireCmd,
				resp: &http.Response{
					StatusCode: int(testStatusCode),
				},
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
				cmd: validateCmd,
				resp: &http.Response{
					StatusCode: int(testStatusCode),
				},
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
				cmd: listCmd,
				resp: &http.Response{
					StatusCode: int(testStatusCode),
				},
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
				cmd: fetchCmd,
				resp: &http.Response{
					StatusCode: int(testStatusCode),
				},
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

			gotErr := printSuccessCLIResponseHumanReadable(tt.args.cmd, tt.args.resp, tt.args.cmdResponse)
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
	}{
		{
			name: "failed publish returns no error and prints JSON as expected",
			args: args{
				cmd:                 publishCmd,
				statusCode:          testStatusCode,
				printJSON:           true,
				gatewayResponseBody: validGatewayResponseBody,
			},
		},
		{
			name: "failed publish returns no error and prints human-readable as expected",
			args: args{
				cmd:                 publishCmd,
				statusCode:          testStatusCode,
				printJSON:           false,
				gatewayResponseBody: validGatewayResponseBody,
			},
		},
		{
			name: "failed publish with invalid gateway response (string statuscode) returns error",
			args: args{
				cmd:                 publishCmd,
				statusCode:          testStatusCode,
				gatewayResponseBody: invalidGatewayResponseBody,
			},
		},
		{
			name: "failed retire returns no error and prints JSON as expected",
			args: args{
				cmd:                 retireCmd,
				statusCode:          testStatusCode,
				printJSON:           true,
				gatewayResponseBody: validGatewayResponseBody,
			},
		},
		{
			name: "failed retire returns no error and prints human-readable as expected",
			args: args{
				cmd:                 retireCmd,
				statusCode:          testStatusCode,
				printJSON:           false,
				gatewayResponseBody: validGatewayResponseBody,
			},
		},
		{
			name: "failed retire with invalid gateway response (string statuscode) returns error",
			args: args{
				cmd:                 retireCmd,
				statusCode:          testStatusCode,
				gatewayResponseBody: invalidGatewayResponseBody,
			},
		},
		{
			name: "failed validate returns no error and prints JSON as expected",
			args: args{
				cmd:                 validateCmd,
				statusCode:          testStatusCode,
				printJSON:           true,
				gatewayResponseBody: validGatewayResponseBody,
			},
		},
		{
			name: "failed validate returns no error and prints human-readable as expected",
			args: args{
				cmd:                 validateCmd,
				statusCode:          testStatusCode,
				printJSON:           false,
				gatewayResponseBody: validGatewayResponseBody,
			},
		},
		{
			name: "failed validate with invalid gateway response (string statuscode) returns error",
			args: args{
				cmd:                 validateCmd,
				statusCode:          testStatusCode,
				gatewayResponseBody: invalidGatewayResponseBody,
			},
		},
		{
			name: "failed list returns no error and prints JSON as expected",
			args: args{
				cmd:                 listCmd,
				statusCode:          testStatusCode,
				printJSON:           true,
				gatewayResponseBody: validGatewayResponseBody,
			},
		},
		{
			name: "failed list returns no error and prints human-readable as expected",
			args: args{
				cmd:                 listCmd,
				statusCode:          testStatusCode,
				printJSON:           false,
				gatewayResponseBody: validGatewayResponseBody,
			},
		},
		{
			name: "failed list with invalid gateway response (string statuscode) returns error",
			args: args{
				cmd:                 listCmd,
				statusCode:          testStatusCode,
				gatewayResponseBody: invalidGatewayResponseBody,
			},
		},
		{
			name: "failed fetch returns no error and prints JSON as expected",
			args: args{
				cmd:                 fetchCmd,
				statusCode:          testStatusCode,
				printJSON:           true,
				gatewayResponseBody: validGatewayResponseBody,
			},
		},
		{
			name: "failed fetch returns no error and prints human-readable as expected",
			args: args{
				cmd:                 fetchCmd,
				statusCode:          testStatusCode,
				printJSON:           false,
				gatewayResponseBody: validGatewayResponseBody,
			},
		},
		{
			name: "failed fetch with invalid gateway response (string statuscode) returns error",
			args: args{
				cmd:                 fetchCmd,
				statusCode:          testStatusCode,
				gatewayResponseBody: invalidGatewayResponseBody,
			},
		},
		{
			name: "failed fetch with invalid gateway response (string statuscode) returns error",
			args: args{
				cmd:                 fetchCmd,
				statusCode:          testStatusCode,
				gatewayResponseBody: invalidGatewayResponseBody,
			},
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

			wantErr := true
			gotErr := printErrorCLIResponse(tt.args.cmd, httpResp)
			if (gotErr != nil) != wantErr {
				t.Errorf("printErrorCLIResponse() got error = %v, want %v", gotErr, wantErr)
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

func Test_traceID(t *testing.T) {
	type args struct {
		cmd          *cobra.Command
		printFuncion func(cmd *cobra.Command, resp *http.Response) error
		printJSON    bool
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "printErrorCLIResponse traceid",
			args: args{
				cmd:          publishCmd,
				printJSON:    false,
				printFuncion: printErrorCLIResponse,
			},
		},
		{
			name: "printErrorCLIResponse with json traceid",
			args: args{
				cmd:          publishCmd,
				printJSON:    true,
				printFuncion: printErrorCLIResponse,
			},
		},
		{
			name: "printSuccessCLIResponseHumanReadable traceid",
			args: args{
				cmd:       publishCmd,
				printJSON: false,
				printFuncion: func(cmd *cobra.Command, resp *http.Response) error {
					return printSuccessCLIResponseHumanReadable(cmd, resp, &publishResponse{})
				},
			},
		},
		{
			name: "printSuccessCLIResponseHumanReadable with json traceid",
			args: args{
				cmd:       publishCmd,
				printJSON: true,
				printFuncion: func(cmd *cobra.Command, resp *http.Response) error {
					return printSuccessCLIResponseHumanReadable(cmd, resp, &publishResponse{})
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			printJSON = tt.args.printJSON
			printTraceID = true

			httpResp := &http.Response{
				Body:       io.NopCloser(bytes.NewBufferString(validGatewayResponseBody)),
				StatusCode: testStatusCode,
				Header: http.Header{
					traceParentHeader: []string{"00-de0ae651ce4c183a3e5d3eb4827c4fc8-43f4db3d9431bc34-01"},
				},
			}

			outBuff := bytes.NewBuffer(nil)
			tt.args.cmd.SetOut(outBuff)

			_ = tt.args.printFuncion(tt.args.cmd, httpResp)
			gotPrintBytes, err := io.ReadAll(outBuff)
			if err != nil {
				t.Error(err)
			}
			gotPrint := strings.TrimRight(string(gotPrintBytes), "\n")

			if !strings.Contains(gotPrint, "de0ae651ce4c183a3e5d3eb4827c4fc8") {
				t.Errorf(`expected message %q to have trace id "de0ae651ce4c183a3e5d3eb4827c4fc8"`, gotPrint)
			}
		})
	}
}
