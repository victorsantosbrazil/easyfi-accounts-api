install-dependencies:
	@echo "Installing dependencies..."
	@go get .
wire:
	@go generate ./src/app/wire.go > /dev/null
build:
	@echo "Building..."
	@go generate -skip=wire ./...
	@make wire > /dev/null
vulncheck: 
	@echo "Checking for vulnerabilities..."
	@go install golang.org/x/vuln/cmd/govulncheck@latest
	@govulncheck ./...
mock:
	@go generate -run=test ./...
test:
	@echo "Executing tests..."
	@go test ./src/...
install: install-dependencies build vulncheck test
run:
	@go run main.go