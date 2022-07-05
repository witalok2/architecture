# Export enviroment variables to commands
export

# Variables
go_cover_file=coverage.out
VERSION = $(shell git branch --show-current)

help:: ## Show this help
	@fgrep -h "##" $(MAKEFILE_LIST) | sort | fgrep -v fgrep | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

compose-up:: 
	cd ./infrastructure && docker-compose --env-file ./.env up

run:: 
	VERSION=$(VERSION) go run ./cmd/main.go -port=9000 -debug=true

test:: ## Do the tests in go
	@ go test -race -coverprofile $(go_cover_file) ./...

lint-install:: ## Install golangci-lint
	@ curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.43.0

lint:: ## Run golang linter
	@ golangci-lint run --verbose --max-issues-per-linter 1000 --max-same-issues 1000

cover-html:: test ## See of the coverage of the code on your default navigator
	@ go tool cover -html=$(go_cover_file)
