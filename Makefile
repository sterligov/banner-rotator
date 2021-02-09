include ./deployments/.env

BIN := "./bin/rotator"
DOCKER_IMG="rotator:develop"

GIT_HASH := $(shell git log --format="%h" -n 1)
LDFLAGS := -X main.release="develop" -X main.buildDate=$(shell date -u +%Y-%m-%dT%H:%M:%S) -X main.gitHash=$(GIT_HASH)

build:
	go build -v -o $(BIN) -ldflags "$(LDFLAGS)" ./cmd

run:
	docker-compose --env-file deployments/.env -f deployments/docker-compose.yml up -d --build --remove-orphans

down:
	docker-compose --env-file deployments/.env -f deployments/docker-compose.yml down

build-img:
	docker build \
		--build-arg=LDFLAGS="$(LDFLAGS)" \
		-t $(DOCKER_IMG) \
		-f build/Dockerfile .

run-img: build-img
	docker run $(DOCKER_IMG)

version: build
	$(BIN) version

test:
	go test -race ./internal/...

long-test:
	go test -race -count=100 ./internal/...

integration-test:
	  chmod +x ./scripts/integration-test.sh && ./scripts/integration-test.sh

install-lint-deps:
	(which golangci-lint > /dev/null) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.30.0

lint: install-lint-deps
	golangci-lint run ./...

wire:
	wire cmd/wire.go

mocks:
	mockery --all --dir internal --output ./internal/mocks --case underscore

migrations:
	goose -dir migrations mysql "${DB_USER}:${DB_PASSWORD}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}?parseTime=true" up

generate:
	go generate ./internal/...

.PHONY: migrations build run build-img run-img version test lint
