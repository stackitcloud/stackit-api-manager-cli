package util

import (
	"encoding/base64"
	"os"
)

// EncodeBase64File to encode a file into base64 for uploading it via an API request
func EncodeBase64File(file string) (string, error) {
	oas, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(oas), nil
}
