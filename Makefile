.PHONY: build test lint vulncheck proto-gen

build:
	go build ./...

test:
	go test -race ./...

lint:
	golangci-lint run ./...

vulncheck:
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...

proto-gen:
	./scripts/protocgen.sh
