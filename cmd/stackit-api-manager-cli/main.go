package main

import (
	"fmt"
	"os"

	"github.com/stackitcloud/stackit-api-manager-cli/internal/log"
	"github.com/stackitcloud/stackit-api-manager-cli/pkg/stackit_api_manager/client"
	"go.uber.org/zap"

	// This controls the maxprocs environment variable in container runtimes.
	// see https://martin.baillie.id/wrote/gotchas-in-the-go-network-packages-defaults/#bonus-gomaxprocs-containers-and-the-cfs
	_ "go.uber.org/automaxprocs"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "an error occurred: %s\n", err)
		os.Exit(1)
	}
}

func run() error {
	logger, err := log.NewAtLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		return err
	}

	defer logger.Sync()

	logger.Info("Hello world!", zap.String("location", "world"))

	identifier := "dev"
	stage := "dev"
	c := client.NewClient("http://127.0.0.1:9002", "test-token")
	projectID := "1"

	// base64Encoded := "b3BlbmFwaTogMy4wLjEKaW5mbzoKICB2ZXJzaW9uOiAiMS4wIgogIHRpdGxlOiBTVEFDS0lUIENvc3RzIEFQSQogIHRlcm1zT2ZTZXJ2aWNlOiAiaHR0cHM6Ly9zdGFja2l0LmRlL2VuL2ltcHJpbnQiCiAgZGVzY3JpcHRpb246IHwKICAgIFRoZSBjb3N0cyBBUEkgcHJvdmlkZXMgZGV0YWlsZWQgcmVwb3J0cyBvbiB0aGUgY29zdHMgZm9yIGEgY3VzdG9tZXIgb3IgcHJvamVjdCBvdmVyIGEgY2VydGFpbiBhbW91bnQgb2YgdGltZQp0YWdzOgogIC0gbmFtZTogc2VydmljZS1jb3N0cwogICAgZGVzY3JpcHRpb246IENvc3RzIHNlcnZpY2UKc2VydmVyczoKICAtIHVybDogImh0dHBzOi8vYXBpLWRldi5zdGFja2l0LmNsb3VkIgogICAgZGVzY3JpcHRpb246IERldmVsb3BtZW50IGVuZHBvaW50CiAgICB4LXN0YWNraXQtdXBzdHJlYW0tc3RhZ2U6IGRldgoKcGF0aHM6CiAgIi9jb3N0cy1zZXJ2aWNlL3YxL2Nvc3RzL3tjdXN0b21lckFjY291bnRJZH0iOgogICAgcGFyYW1ldGVyczoKICAgICAgLSBzY2hlbWE6CiAgICAgICAgICB0eXBlOiBzdHJpbmcKICAgICAgICAgIGZvcm1hdDogdXVpZAogICAgICAgIG5hbWU6IGN1c3RvbWVyQWNjb3VudElkCiAgICAgICAgaW46IHBhdGgKICAgICAgICByZXF1aXJlZDogdHJ1ZQogICAgICAgIGRlc2NyaXB0aW9uOiBJRCBvZiBhIGN1c3RvbWVyIGFjY291bnQKICAgIGdldDoKICAgICAgc3VtbWFyeTogUmVwb3J0cyBmb3IgYWxsIHByb2plY3RzCiAgICAgIHRhZ3M6CiAgICAgICAgLSBzZXJ2aWNlLWNvc3RzCiAgICAgIHJlc3BvbnNlczoKICAgICAgICAiMjAwIjoKICAgICAgICAgIGRlc2NyaXB0aW9uOiBPSwogICAgICB4LXN0YWNraXQtYXV0aG9yaXphdGlvbjoKICAgICAgICByZXNvdXJjZUlEOiBjdXN0b21lckFjY291bnRJZAogICAgICAgIGFjdGlvbjogYWNjb3VudC5jb250cm9sbGluZy5yZWFkCiAgIi9jb3N0cy1zZXJ2aWNlL3YxL2Nvc3RzL3tjdXN0b21lckFjY291bnRJZH0vcHJvamVjdHMve3Byb2plY3RJZH0iOgogICAgcGFyYW1ldGVyczoKICAgICAgLSBzY2hlbWE6CiAgICAgICAgICB0eXBlOiBzdHJpbmcKICAgICAgICAgIGZvcm1hdDogdXVpZAogICAgICAgIG5hbWU6IGN1c3RvbWVyQWNjb3VudElkCiAgICAgICAgaW46IHBhdGgKICAgICAgICByZXF1aXJlZDogdHJ1ZQogICAgICAgIGRlc2NyaXB0aW9uOiBJRCBvZiBhIGN1c3RvbWVyIGFjY291bnQKICAgICAgLSBzY2hlbWE6CiAgICAgICAgICB0eXBlOiBzdHJpbmcKICAgICAgICAgIGZvcm1hdDogdXVpZAogICAgICAgIG5hbWU6IHByb2plY3RJZAogICAgICAgIGluOiBwYXRoCiAgICAgICAgcmVxdWlyZWQ6IHRydWUKICAgICAgICBkZXNjcmlwdGlvbjogSUQgb2YgYSBwcm9qZWN0CiAgICBwb3N0OgogICAgICByZXF1ZXN0Qm9keToKICAgICAgICBjb250ZW50OgogICAgICAgICAgYXBwbGljYXRpb24vanNvbjoKICAgICAgICAgICAgc2NoZW1hOgogICAgICAgICAgICAgIHR5cGU6IHN0cmluZwogICAgICAgICAgYXBwbGljYXRpb24veG1sOgogICAgICAgICAgICBzY2hlbWE6CiAgICAgICAgICAgICAgdHlwZTogc3RyaW5nCiAgICAgIHN1bW1hcnk6IFJlcG9ydHMgZm9yIGEgY2VydGFpbiBwcm9qZWN0CiAgICAgIHRhZ3M6CiAgICAgICAgLSBzZXJ2aWNlLWNvc3RzCiAgICAgIHJlc3BvbnNlczoKICAgICAgICAiMjAwIjoKICAgICAgICAgIGRlc2NyaXB0aW9uOiBPSwogICAgICB4LXN0YWNraXQtYXV0aG9yaXphdGlvbjoKICAgICAgICByZXNvdXJjZUlEOiBwcm9qZWN0SWQKICAgICAgICByZXNvdXJjZVR5cGU6IHByb2plY3QKICAgICAgICBhY3Rpb246IHByb2plY3QuY29udHJvbGxpbmcud3JpdGUK"
	// resp, _, err := c.PublishAPIEndpoint(projectID, &client.PublishAPIEndpoint{
	// 	Metadata: &client.Metadata{
	// 		Identifier: &identifier,
	// 		Stage:      &stage,
	// 	},
	// 	Spec: &client.Spec{
	// 		OpenAPI: &client.OpenAPI{
	// 			Base64Encoded: &base64Encoded,
	// 		},
	// 	},
	// })
	// if err != nil {
	// 	return err
	// }
	// fmt.Println(resp)

	resp, _, err := c.RetireAPIEndpoint(projectID, &client.RetireAPIEndpoint{
		Metadata: &client.Metadata{
			Identifier: &identifier,
			Stage:      &stage,
		},
	})
	if err != nil {
		return err
	}
	fmt.Println(resp)

	return nil
}
