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
Tools used for the developers during development process.

* [golangci-lint](https://golangci-lint.run/)
* [govulncheck](https://pkg.go.dev/golang.org/x/vuln/cmd/govulncheck)

### Instalation

To install all tools, run the following command:

```
./dev/scripts/linux/install-dev-tools.sh
```

## 🎨 Local Environment

### Start

To start local environment, run the following command:

```
./dev/scripts/windows/start-local-env.sh
```

### Stop

To stop local environment, run the following command:

```
./dev/scripts/linux/stop-local-env.sh
```