#!/bin/bash

echo "Installing golangci-lint..."
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest || { echo "Error: golangci-lint installation failed"; exit 1; }

echo "Installing govulncheck..."
go install golang.org/x/vuln/cmd/govulncheck@latest || { echo "Error: govulncheck installation failed"; exit 1; }

echo "Installing gomock..."
go install github.com/golang/mock/mockgen@v1.6.0 || { echo "Error: gomock installation failed"; exit 1; }

echo "Installing golang-migrate..."
go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

echo "Dev tools installed successfully."