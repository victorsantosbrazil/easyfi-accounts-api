# Dev Manual

<p>The aim of the dev manual is to help developers, providing guidelines, IDE setup, tools, local environment boostrap, helper scripts, and other useful stuff. </p>

##  📔 Table of Contents
<!--ts-->
   * [Pre-requirements](#✂️-pre-requirements)
   * [DevTools](#🔨-devtools)
      * [Instalation](#instalation)
   * [Local environment](#🎨-local-environment)
      * [Start](#start)
      * [Stop](#stop)
<!--te-->

## ✂️ Pre-requirements
* [Golang 1.20](https://go.dev/doc/install)
* [Docker](https://docs.docker.com/get-docker/)

## 🔨 DevTools

* [golangci-lint](https://golangci-lint.run/)
* [govulncheck](https://pkg.go.dev/golang.org/x/vuln/cmd/govulncheck)
* [gomock](https://github.com/golang/mock)

### Instalation

To install all tools, run the following command:

```
./dev/scripts/install-dev-tools.sh
```

## 🎨 Local Environment

### Start

To start local environment, run the following command:

```
./dev/scripts/start-local-env.sh
```

### Stop

To stop local environment, run the following command:

```
./dev/scripts/stop-local-env.sh
```

## Tests

### Run
```
go test .
```