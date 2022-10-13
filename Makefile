.DEFAULT_GOAL := build

.PHONY:fmt
fmt:
	go fmt ./...

.PHONY:lint
lint: fmt
	golint ./...

.PHONY:vet
vet: fmt
	go vet ./...
