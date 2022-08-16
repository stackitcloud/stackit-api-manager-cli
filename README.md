# STACKIT API Manager CLI

[![GoTemplate](https://img.shields.io/badge/go/template-black?logo=go)](https://github.com/SchwarzIT/go-template)
[![SIT](https://img.shields.io/badge/SIT-awesome-blueviolet.svg)](https://jobs.schwarz)
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
go install github.com/stackitcloud/stackit-api-manager-cli/cmd/gt@latest
```

> Based on your go configuration the `stackit-api-manager-cli` binary can be found in `$GOPATH/bin` or `$HOME/go/bin` in case `$GOPATH` is not set.
> Make sure to add the respective directory to your `$PATH`.
> [For more information see go docs for further information](https://golang.org/ref/mod#go-install). Run `go env` to view your current configuration.

#### From the released binaries

Download the desired version for your operating system and processor architecture from the [STACKIT API Manager CLI](https://github.com/stackitcloud/stackit-api-manager-cli/releases).
Make the file executable and place it in a directory available in your `$PATH`.

### Interact with STACKIT API Manager service

Use the CLI to publish your OpenAPI Spec:

```bash
stackit-api-manager-cli project publish \
  ---identifier <YourIdentifier> \
  --project <YourProject> \
  --stage <YourStage> \
  --token <YourAuthToken> \
  --oas <PathToOpenAPISpec>
```

Use the CLI to retire your OpenAPI Spec:

```bash
stackit-api-manager-cli project retire \
  ---identifier <YourIdentifier> \
  --project <YourProject> \
  --stage <YourStage> \
  --token <YourAuthToken>
```

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

## Maintainers

| Name                                               | Email                        |
| :------------------------------------------------- | :--------------------------- |
| [@fabiolamicela](https://github.com/fabiolamicela) | fabio.la_micela@mail.schwarz |
