MAKE_FILE_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

.PHONY: server
server:
	go run cmd/api/main.go

.PHONY: sqlc
sqlc:
	docker run --rm -v ${MAKE_FILE_DIR}:/src -w /src sqlc/sqlc generate