.PHONY: defaultbuild test docs clean dev
# Variables
APP_NAME=gocrud

# Tasks
default: run-with-docs

run-with-docs:
	@export PATH=$$PATH:$(shell go env GOPATH)/bin && swag init && go run main.go
build:
	@go build -o $(APP_NAME) main.go
test:
	@go test ./ ...
docs:
	@swag init
clean:
	@rm -f $(APP_NAME)
	@rm -rf ./docs
dev:
	@go run main.go