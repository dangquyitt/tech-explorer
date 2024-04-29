MAKE_FILE_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

.PHONY: server
server:
	go run cmd/api/main.go

.PHONY: mock
mock:
	docker run --rm -v ${MAKE_FILE_DIR}:/src -w /src vektra/mockery:v2 --all