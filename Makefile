## teapotbot.dev build tooling

MAKEFLAGS += --silent

BUILD_NAME = teapotbot
GOLANGCI_LINT_VERSION = v1.50.1


all: help

## help: Prints a list of available build targets.
help:
	echo "Usage: make <OPTIONS> ... <TARGETS>"
	echo ""
	echo "Available targets are:"
	echo ''
	sed -n 's/^##//p' ${PWD}/Makefile | column -t -s ':' | sed -e 's/^/ /'
	echo
	echo "Targets run by default are: `sed -n 's/^all: //p' ./Makefile | sed -e 's/ /, /g' | sed -e 's/\(.*\), /\1, and /'`"

## docker: spin up local dev stack via docker-compose
docker:
	docker-compose -f dockerfiles/docker-compose.yml up -d --remove-orphans

## lint: Lint with golangci-lint
lint:
	docker run --rm -v $$(pwd):/repo -w /repo golangci/golangci-lint:${GOLANGCI_LINT_VERSION} golangci-lint run --verbose --color always ./...

pre-commit:
	pre-commit run --all-files --show-diff-on-failure

## fmt: Format with gofmt
fmt:
	go fmt ./...

# tidy: Tidy with go mod tidy
tidy:
	go mod tidy -compat=1.17

## pre-commit: Chain lint + test
pre-commit: test lint

## proto: Protobuf buildchain
proto:
	rm -f **/*.pb.go
	protoc --proto_path=proto proto/*.proto --go_out=. --fatal_warnings

## test: Test with go test
test: proto
	go test -test.v -race -covermode=atomic -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out
	rm coverage.out

## test-perf: Benchmark tests with go test -bench
test-perf:
	go test -test.v -benchmem -bench=. -coverprofile=coverage-bench.out ./... && go tool cover -html=coverage-bench.out && rm coverage-bench.out

## build: build teapotbot
build:
	go build -o ${BUILD_NAME} cmd/main.go

## run: go run main.go
run:
	go run cmd/main.go

## clean: cleanup build artifacts
clean:
	rm -f ${BUILD_NAME}

.PHONY: docker lint fmt tidy pre-commit proto test test-perf build run clean
