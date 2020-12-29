SHELL := /bin/bash

tidy:
	go mod tidy
	go mod vendor

build:
	go mod build

run:
	go run main.go