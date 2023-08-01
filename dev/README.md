# Dev Manual

<p>The aim of the dev manual is to help developers through the development process, providing guidelines, IDE setup, tools, local environment setup, helper scripts, and other useful information. </p>

##  📔 Table of Contents
<!--ts-->
   * [Pre-requirements](#✂️-pre-requirements)
   * [DevTools](#🔨-devtools)
      * [Instalation](#instalation)
<!--te-->

## ✂️ Pre-requirements
* [Golang 1.20](https://go.dev/doc/install)

## 🔨 DevTools
Tools used by the developers during the process of development.

* [golangci-lint](https://golangci-lint.run/)
* [govulncheck](https://pkg.go.dev/golang.org/x/vuln/cmd/govulncheck)

### Instalation

To install all tools run the following command based on your operation system.

* Linux

```
./dev/scripts/linux/install-dev-tools.sh
```

* Windows
```
./dev/scripts/windows/install-dev-tools.bat
```