include .env

.PHONY: default

BIN_NAME=eligibility
VERSION ?= dev
GIT_COMMIT ?=$(shell git rev-parse HEAD)
BUILD_DATE ?= $(shell date +%FT%T%z)
TIMESTAMP ?= $(shell date +%Y%m%d%H%M)

default: help

help:
	@echo 'Management commands for eligibility:'
	@echo
	@echo 'Usage:'
	@echo '    make build        			Builds the binary locally.'
	@echo '    make update-deps      	Runs dep ensure.'
	@echo '    make package       		Build final docker image with just the go binary inside.'
	@echo '    make add-migration  		Create a new migration file.'
	@echo '    make build-migration  	Create a new migration file.'
	@echo '    make test          		Run tests on a compiled project.'
	@echo '    make run          			Build and run'
	@echo '    make up          			Start containers'
	@echo '    make down          		Stop and delete containers'
	@echo '    make deploy-dev    		Deploy tagged image to staging'
	@echo '    make deploy-prod   		Deploy tagged image to production'
	@echo '    make clean         		Clean the directory tree.'
	@echo


build: build-service build-migrations

run: build 
	bin/eligibility

build-service:
	@echo "Building service"
	mkdir -p ./bin
	go build -ldflags "-w -X main.GitCommit=${GIT_COMMIT} -X main.Version=${VERSION} -X main.BuildDate=${BUILD_DATE}" -o ./bin/eligibility ./cmd/eligibility/

build-migrations:
	docker build -t eligibility-migrations migrations

up:
	docker-compose up --build -d

run:
	@echo "Running service"
	go run ./cmd/eligibility/

update-deps:
	dep ensure -update

package: build-migrations
	@echo "Building image ${BIN_NAME} ${VERSION} $(GIT_COMMIT)"
	docker build --build-arg VERSION=${VERSION} --build-arg GIT_COMMIT=$(GIT_COMMIT) -t $(IMAGE_NAME):${VERSION} .

add-migration:
	touch ./migrations/sql/${TIMESTAMP}_$(name).sql

clean:
	@test ! -e bin/${BIN_NAME} || rm bin/${BIN_NAME}

deploy-dev: get-credentials-dev docker-build-dev push-dev
	kubectl set image deployment/eligibility eligibility=gcr.io/gems-org/eligibility-dev:$(VERSION)

deploy-prod: get-credentials-prod docker-build-prod push-prod
	kubectl set image deployment/eligibility eligibility=gcr.io/gems-org/eligibility:$(VERSION)

test:
	go test ./... -v -count=1

down:
	docker-compose down
