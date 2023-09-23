install-dependencies:
	go get .
vulncheck: 
	go install golang.org/x/vuln/cmd/govulncheck@latest
	govulncheck ./...
build:
	go generate ./src/...
	go generate ./src/main/wire.go
install: install-dependencies vulncheck build
run:
	go run main.go

