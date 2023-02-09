package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/spf13/cobra"
)

func Test_retrieveGatewayErrorMessage(t *testing.T) {
	tests := []struct {
		name            string
		gatewayResponse string
		httpStatusCode  int
		want            string
		wantErr         bool
	}{
		{
			name:            "retrieve message successful without statusCode in HTTP response",
			gatewayResponse: `{"Status": 123, "Message": "error message"}`,
			want:            "error message",
			wantErr:         false,
		},
		{
			name:            "retrieve message successful with statusCode in HTTP response",
			gatewayResponse: `{"Status": 123, "Message": "error message"}`,
			httpStatusCode:  321,
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
			name: "nil cmd response on successful publish returns error",
			args: args{
				cmd:         publishCmd,
				statusCode:  200,
				cmdResponse: nil,
			},
			wantErr: errNilCmdResponse,
		},
		{
			name: "printing CLI response on successful publish returns no error and prints as expected",
			args: args{
				cmd:        publishCmd,
				statusCode: 123,
				cmdResponse: publishResponse{
					Identifier: "api-id",
					ProjectID:  "project-id",
					Stage:      "stage",
					APIURL:     "test.url/fun",
				},
			},
			wantPrint: `{"success":true,"statusCode":123,"message":"API published successfully","response":{"identifier":"api-id","projectId":"project-id","stage":"stage","apiUrl":"test.url/fun"}}`,
			wantErr:   nil,
		},
		{
			name: "nil cmd response on successful retire returns error",
			args: args{
				cmd:         retireCmd,
				statusCode:  123,
				cmdResponse: nil,
			},
			wantErr: errNilCmdResponse,
		},
		{
			name: "printing CLI response on successful retire returns no error and prints as expected",
			args: args{
				cmd:        retireCmd,
				statusCode: 123,
				cmdResponse: retireResponse{
					Identifier: "api-id",
					ProjectID:  "project-id",
				},
			},
			wantPrint: `{"success":true,"statusCode":123,"message":"API retired successfully","response":{"identifier":"api-id","projectId":"project-id"}}`,
			wantErr:   nil,
		},
		{
			name: "nil cmd response on successful validate returns error",
			args: args{
				cmd:         validateCmd,
				statusCode:  123,
				cmdResponse: nil,
			},
			wantErr: errNilCmdResponse,
		},
		{
			name: "printing CLI response on successful validate returns no error and prints as expected",
			args: args{
				cmd:        validateCmd,
				statusCode: 123,
				cmdResponse: validateResponse{
					Identifier: "api-id",
					ProjectID:  "project-id",
					Stage:      "stage",
				},
			},
			wantPrint: `{"success":true,"statusCode":123,"message":"OpenAPI Specification validated successfully","response":{"identifier":"api-id","projectId":"project-id","stage":"stage"}}`,
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

			var wantPrintBytes []byte
			if gotErr == nil {
				wantPrintBytes = []byte(fmt.Sprintln(tt.wantPrint))
			}
			gotPrintBytes, err := io.ReadAll(outBuff)
			if err != nil {
				t.Error(err)
			}

			if !bytes.Equal(gotPrintBytes, wantPrintBytes) {
				t.Errorf("printSuccessCLIResponseJSON() expected message to be\n%v\nbut got\n%v", tt.wantPrint, string(gotPrintBytes))
			}
		})
	}
}

func Test_printErrorCLIResponseJSON(t *testing.T) {
	type args struct {
		cmd             *cobra.Command
		statusCode      int
		gatewayResponse string
	}
	tests := []struct {
		name      string
		args      args
		wantPrint string
		wantErr   bool
	}{
		{
			name: "printing CLI response on failed publish returns no error and prints as expected",
			args: args{
				cmd:             publishCmd,
				statusCode:      123,
				gatewayResponse: `{"Status": 123, "Message": "error message"}`,
			},
			wantPrint: `{"success":false,"statusCode":123,"message":"error message"}`,
			wantErr:   false,
		},
		{
			name: "printing CLI response on failed publish with invalid gateway response (string statuscode) returns error",
			args: args{
				cmd:             publishCmd,
				statusCode:      123,
				gatewayResponse: `{"Status": "123", "Message": "error message"}`,
			},
			wantErr: true,
		},
		{
			name: "printing CLI response on failed retire returns no error and prints as expected",
			args: args{
				cmd:             retireCmd,
				statusCode:      123,
				gatewayResponse: `{"Status": 123, "Message": "error message"}`,
			},
			wantPrint: `{"success":false,"statusCode":123,"message":"error message"}`,
			wantErr:   false,
		},
		{
			name: "printing CLI response on failed retire with invalid gateway response (string statuscode) returns error",
			args: args{
				cmd:             retireCmd,
				statusCode:      123,
				gatewayResponse: `{"Status": "123", "Message": "error message"}`,
			},
			wantErr: true,
		},
		{
			name: "printing CLI response on failed validate returns no error and prints as expected",
			args: args{
				cmd:             validateCmd,
				statusCode:      123,
				gatewayResponse: `{"Status": 123, "Message": "error message"}`,
			},
			wantPrint: `{"success":false,"statusCode":123,"message":"error message"}`,
			wantErr:   false,
		},
		{
			name: "printing CLI response on failed validate with invalid gateway response (string statuscode) returns error",
			args: args{
				cmd:             validateCmd,
				statusCode:      123,
				gatewayResponse: `{"Status": "123", "Message": "error message"}`,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			httpResp := &http.Response{
				Body:       io.NopCloser(bytes.NewBufferString(tt.args.gatewayResponse)),
				StatusCode: tt.args.statusCode,
			}

			outBuff := bytes.NewBuffer(nil)
			tt.args.cmd.SetOut(outBuff)

			gotErr := printErrorCLIResponseJSON(tt.args.cmd, httpResp)
			if (gotErr != nil) != tt.wantErr {
				t.Errorf("printErrorCLIResponseJSON() got error = %v, want %v", gotErr, tt.wantErr)
			}

			var wantPrintBytes []byte
			if gotErr == nil {
				wantPrintBytes = []byte(fmt.Sprintln(tt.wantPrint))
			}
			gotPrintBytes, err := io.ReadAll(outBuff)
			if err != nil {
				t.Error(err)
			}

			if !bytes.Equal(gotPrintBytes, wantPrintBytes) {
				t.Errorf("printErrorCLIResponseJSON() expected message to be\n%v\nbut got\n%v", tt.wantPrint, string(gotPrintBytes))
			}
		})
	}
}
