SHELL := /bin/bash

tidy:
	go mod tidy
	# go mod vendor

build:
	go mod build

run:
	go run main.go

uml:
	mkdir -p resources/uml
	goplantuml -recursive . > resources/uml/BankHexArh.puml
	plantuml -tpng resources/uml/BankHexArh.puml
	mv resources/uml/BankHexArh.png resources/images/BankHexArh.latest.png

startdb:
	docker-compose -f resources/docker/docker-compose.yml up -d

stopdb:
	docker-compose -f resources/docker/docker-compose.yml down