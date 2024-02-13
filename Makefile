install-dependencies:
	@echo "Installing dependencies..."
	@go get .
wire:
	@go install github.com/google/wire/cmd/wire@v0.5.0
	@go generate ./src/app/wire.go > /dev/null
build:
	@echo "Building..."
	@make wire > /dev/null
vulncheck: 
	@echo "Checking for vulnerabilities..."
	@go install golang.org/x/vuln/cmd/govulncheck@latest
	@govulncheck -show verbose ./...
mock:
	@echo "Creating mocks for tests..."
	@go install github.com/golang/mock/mockgen@v1.6.0
	@go generate -run=mockgen ./...
test:
	@echo "Executing tests..."
	@go test ./src/...
doc:
	@echo "Generating docs..."
	@go install github.com/swaggo/swag/cmd/swag@latest
	@swag init -o ./docs/swagger	
install: install-dependencies build vulncheck mock test doc
run:
	@go run main.go