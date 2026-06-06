include .env

GOOSE_DBSTRING=root:Adam.123@tcp(127.0.0.1:3307)/shopdevgo
GOOSE_MIGRATION_DIR ?= sql/schema
GOOSE_DRIVER ?= mysql

# Tên của ứng dụng cửa bạn
APP_NAME := server

# Chạy ứng dụng
docker_build:
	docker-compose up -d --build
	docker-compose ps
docker_up:
	docker-compose -f environment/docker-compose-dev.yml up
docker_stop:
	docker-compose down

dev: 
	go run ./cmd/$(APP_NAME)
run:
	docker compose up -d && go run ./cmd/$(APP_NAME)
kill: 
	docker compose kill
up:
	docker compose up -d
down: 
	docker compose down
up_by_one:
	@set GOOSE_DRIVER=$(GOOSE_DRIVER)&& set GOOSE_DBSTRING=$(GOOSE_DBSTRING)&& goose -dir=$(GOOSE_MIGRATION_DIR) up-by-one
# create new a migration
create_migration:
	@goose -dir=$(GOOSE_MIGRATION_DIR) create $(name) sql
upse:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) up
downse:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) down
resetse:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) reset

sqlgen:
	@"$(USERPROFILE)\go\bin\sqlc.exe" generate

swag:
	swag init -g .\cmd\server\main.go -o .\cmd\swag\docs

.PHONY: run downse upse resetse

.PHONY: air