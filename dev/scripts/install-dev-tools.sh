#!/bin/bash

echo "Installing golangci-lint..."

go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest || { echo "Error: golangci-lint installation failed"; exit 1; }

echo "Dev tools installed successfully."