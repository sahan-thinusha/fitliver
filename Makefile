# Helper variables and Make settings
.PHONY: help clean build proto-link proto-vendor run
.DEFAULT_GOAL := help
.ONESHELL :
.SHELLFLAGS := -ec
SHELL := /bin/bash

help:                                  ## Print list of tasks
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z0-9_%-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' Makefile

clean:
	rm -rf vendor
	rm -rf protobuf-import
	rm -rf pb/*
	rm -rf build/*

proto-link:                            ## Generate go protobuf files using symlinked modules
	./protobuf-import.sh
	protoc -I ./protobuf-import -I ./ ./pub/job.proto --gofast_out=plugins=grpc:.

proto-vendor:                          ## Generate go protobuf files using go mode vendor
	go mod vendor
	protoc -I ./vendor -I ./ ./pub/job.proto --gofast_out=plugins=grpc:.


build:
	go build

init:
	go mod init fitliver && go mod tidy &&  go get -u github.com/swaggo/swag/cmd/swag && swag init

run:
	go run main.go
