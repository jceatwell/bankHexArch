SHELL := /bin/bash

tidy:
	go mod tidy
	# go mod vendor

build:
	go mod build

run:
	go run main.go

uml:
	mkdir -p docs/uml
	goplantuml -recursive . > docs/uml/BankHexArh.puml
	plantuml -tpng  docs/uml/BankHexArh.puml
