.DEFAULT_GOAL := build

fmt: 
  go fmt ./...
.PHONY:fmt

lint: fmt
  golint ./...
.PHONY:lint

vet: fmt
  go vet ./...
.PHONY:vet

build: vet
  go build hello.go
.PHONY: build

# The .DEFAULT_GOAL defines which target is run when no target are specified, in this case the build target
# The word before the : is the name of the target, any words after the target (like vet in the line build:vet)
# are the other tagets that must run before the specified target runs
# Tasks that are performed by the target are on the indented lines after the target
# The .PHONY line keeps make from getting confused if you ever create a directory in your project with the same a
# name as the target

# type "make" to run
