SHELL := /bin/bash

.EXPORT_ALL_VARIABLES:
SERVER_ADDRESS = localhost
SERVER_PORT = 8000
DB_USER = root
DB_PASSWD = codepass
DB_ADDR = localhost
DB_PORT = 3306
DB_NAME = banking


tidy:
	go mod tidy
	go mod vendor

build:
	go mod build

run:
	go run main.go

test_env:
	@echo $$SERVER_ADDRESS $$SERVER_PORT

uml_overview:
	plantuml -tpng resources/uml/Overview_Hex_architecture.puml

uml:
	mkdir -p resources/uml
	goplantuml -aggregate-private-members -show-aliases -show-compositions -recursive . > resources/uml/BankHexArh.puml
	plantuml -tpng resources/uml/BankHexArh.puml
	mv resources/uml/BankHexArh.png resources/images/BankHexArh.latest.png

startdb:
	docker-compose -f resources/docker/docker-compose.yml up -d

stopdb:
	docker-compose -f resources/docker/docker-compose.yml down