include .env
LAST_MIGRATION = $(shell ls -tr migrations/sql/ | tail -n 1 | cut -d'_' -f1)

ifeq ($(LAST_MIGRATION),)
	LAST_MIGRATION := 0
endif


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


build: build-service

run: build 
	bin/eligibility

build-service:
	@echo "Building service"
	mkdir -p ./bin
	go build -ldflags "-w -X main.GitCommit=${GIT_COMMIT} -X main.Version=${VERSION} -X main.BuildDate=${BUILD_DATE}" -o ./bin/eligibility ./cmd/eligibility/

up:
	docker-compose up --build
	
update-deps:
	dep ensure -update

package: build-migrations
	@echo "Building image ${BIN_NAME} ${VERSION} $(GIT_COMMIT)"
	docker build --build-arg VERSION=${VERSION} --build-arg GIT_COMMIT=$(GIT_COMMIT) -t $(IMAGE_NAME):${VERSION} .

clean:
	@test ! -e bin/${BIN_NAME} || rm bin/${BIN_NAME}

deploy-dev: get-credentials-dev docker-build-dev push-dev
	kubectl set image deployment/eligibility eligibility=gcr.io/gems-org/eligibility-dev:$(VERSION)

deploy-prod: get-credentials-prod docker-build-prod push-prod
	kubectl set image deployment/eligibility eligibility=gcr.io/gems-org/eligibility:$(VERSION)

run-tests:
	go test ./... -v -count=1

down:
	docker-compose down

add-migration:
	touch migrations/sql/$(shell expr $(LAST_MIGRATION) + 1 )_$(name).up.sql
	touch migrations/sql/$(shell expr $(LAST_MIGRATION) + 1 )_$(name).down.sql

build-migrations:
	docker build -t eligibility-migration migrations

run-migrations:
	docker run --network host eligibility-migration \
	$(action) $(version) \
	"mysql://$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)"

db-seed:
	@echo "Seeding db"
	mkdir -p ./bin
	go build -o ./bin/dbseed ./pkg/database/dbseed
	./bin/dbseed