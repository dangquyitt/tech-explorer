MAKE_FILE_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

.PHONY: server
server:
	go run cmd/api/main.go

.PHONY: generate
generate:
	go generate ./...
