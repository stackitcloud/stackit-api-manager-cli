# STACKIT API Manager CLI

[![GoTemplate](https://img.shields.io/badge/go/template-black?logo=go)](https://github.com/SchwarzIT/go-template)
[![CI](https://github.com/stackitcloud/stackit-api-manager-cli/actions/workflows/main.yml/badge.svg)](https://github.com/stackitcloud/stackit-api-manager-cli/actions/workflows/main.yml)
[![Semgrep](https://github.com/stackitcloud/stackit-api-manager-cli/actions/workflows/semgrep.yml/badge.svg)](https://github.com/stackitcloud/stackit-api-manager-cli/actions/workflows/semgrep.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/stackitcloud/stackit-api-manager-cli)](https://goreportcard.com/report/github.com/stackitcloud/stackit-api-manager-cli)

CLI for interacting with STACKIT API Manager Service.

The project uses `make` to make your life easier. If you're not familiar with Makefiles you can take a look at [this quickstart guide](https://makefiletutorial.com).

Whenever you need help regarding the available actions, just use the following command.

```bash
make help
```

## Usage

### Installation

#### From source

If you have Go 1.16+, you can directly install by running:

```bash
go install github.com/stackitcloud/stackit-api-manager-cli/cmd/stackit-api-manager@latest
```

> Based on your go configuration the `stackit-api-manager` binary can be found in `$GOPATH/bin` or `$HOME/go/bin` in case `$GOPATH` is not set.
> Make sure to add the respective directory to your `$PATH`.
> [For more information see go docs for further information](https://golang.org/ref/mod#go-install). Run `go env` to view your current configuration.

#### From the released binaries

Download the desired version for your operating system and processor architecture from the [STACKIT API Manager CLI](https://github.com/stackitcloud/stackit-api-manager-cli/releases).
Make the file executable and place it in a directory available in your `$PATH`.

### Interact with STACKIT API Manager service

Use the CLI to publish your OpenAPI Spec:

```bash
stackit-api-manager project publish \
  --identifier <YourIdentifier> \
  --baseURL <API-Manager-BaseURL> \
  --project <YourProject> \
  --stage <YourStage> \
  --token <YourAuthToken> \
  --oas <PathToOpenAPISpec>
```

Use the CLI to retire your OpenAPI Spec:

```bash
stackit-api-manager project retire \
  --identifier <YourIdentifier> \
  --baseURL <API-Manager-BaseURL> \
  --project <YourProject> \
  --token <YourAuthToken>
```

Use the CLI to validate your OpenAPI Spec:

```bash
stackit-api-manager project validate \
  --identifier <YourIdentifier> \
  --baseURL <API-Manager-BaseURL> \
  --project <YourProject> \
  --stage <YourStage> \
  --token <YourAuthToken> \
  --oas <PathToOpenAPISpec>
```

For each request, you can add the `--json` flag to print the CLI response in JSON format instead of receiving a human-readable message.

#### API-Manager-BaseURL possible values

- `https://api-manager.api.eu01.dev.stackit.cloud` for `dev-eu01`
- `https://api-manager.api.dev.stackit.cloud` for `dev-global`
- `https://api-manager.api.eu01.stg.stackit.cloud` for `stg-eu01`  
- `https://api-manager.api.stg.stackit.cloud` for `stg-global`
- `https://api-manager.api.eu01.stackit.cloud` for `prd-eu01`
- `https://api-manager.api.stackit.cloud` for `prd-global`

#### Authenticated token

The given authentication must be of Bearer type: `Bearer <token>`

The Bearer prefix is already included in the authorization header, so the user only needs to provide the OAuth 2.0 `<token>` for the `--token` flag argument.

For more information regarding the Bearer Authentication, please [click here](https://swagger.io/docs/specification/authentication/bearer-authentication/) (last accessed on 30/09/2022).

## Setup

To get your setup up and running the only thing you have to do is

```bash
make all
```

This will initialize a git repo, download the dependencies in the latest versions and install all needed tools.
If needed code generation will be triggered in this target as well.

## Test & lint

Run linting

```bash
make lint
```

Run tests

```bash
make test
```

## Update API client

To update the API client:

- get the OpenAPI Spec from the `https://internal-docs.api.stackit.cloud/oas/api-manager`.
- store OpenAPI Spec in `api/api_manager.openapi.json`.
- then run:

```bash
make generate-client
```

## Maintainers

| Name                                               | Email                        |
| :------------------------------------------------- | :--------------------------- |
| [@fabiolamicela](https://github.com/fabiolamicela) | fabio.la_micela@mail.schwarz |
