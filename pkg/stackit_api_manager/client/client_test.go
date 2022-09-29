package client

import (
	"errors"
	"testing"
)

const (
	token = "some-random-token"
)

func Test_GetToken(t *testing.T) {
	tests := []struct {
		description      string
		auth             string
		wantedError      error
		expectedResponse string
	}{
		{
			description:      "auth is empty",
			auth:             "",
			wantedError:      errMissingAuthentication,
			expectedResponse: "",
		},
		{
			description:      "token is empty with Bearer in auth",
			auth:             bearer,
			wantedError:      errMissingToken,
			expectedResponse: "",
		},
		{
			description:      "auth is invalid",
			auth:             "invalid token",
			wantedError:      errInvalidAuthentication,
			expectedResponse: "",
		},
		{
			description:      "auth is valid",
			auth:             bearer + " " + token,
			wantedError:      nil,
			expectedResponse: bearer + " " + token,
		},
	}
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			token, err := getToken(tt.auth)
			if token != tt.expectedResponse {
				t.Fatalf("response does not match, expecting %s, got %s", tt.expectedResponse, token)
			}
			if !errors.Is(err, tt.wantedError) {
				t.Fatalf("error does not match, expecting %s, got %s", tt.wantedError, err)
			}
		})
	}
}
