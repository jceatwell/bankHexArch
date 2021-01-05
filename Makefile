SHELL := /bin/bash

.EXPORT_ALL_VARIABLES:
MY_VAR_1 = foo
MY_VAR_2 = bar
MY_VAR_3 = baz

tidy:
	go mod tidy
	go mod vendor

build:
	go mod build

run:
	go run main.go

test:
	@echo $$MY_VAR_1 $$MY_VAR_2 $$MY_VAR_3

uml:
	mkdir -p resources/uml
	goplantuml -recursive . > resources/uml/BankHexArh.puml
	plantuml -tpng resources/uml/BankHexArh.puml
	mv resources/uml/BankHexArh.png resources/images/BankHexArh.latest.png

startdb:
	docker-compose -f resources/docker/docker-compose.yml up -d

stopdb:
	docker-compose -f resources/docker/docker-compose.yml down