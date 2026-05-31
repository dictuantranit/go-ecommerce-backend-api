GOOSE_DBSTRING= "root:Adam.123@tcp(127.0.0.1:3307)/shopdevgo"
GOOSE_MIGRATION_DIR ?= sql/schema
GOOSE_DRIVER ?= mysql

# Tên của ứng dụng cửa bạn
APP_NAME := server

# Chạy ứng dụng
docker_build:
	docker-compose up -d --build
	docker-compose ps

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
upse:
	@goose -dir=$(GOOSE_MIGRATION_DIR) $(GOOSE_DRIVER) $(GOOSE_DBSTRING) up
downse:
	@goose -dir=$(GOOSE_MIGRATION_DIR) $(GOOSE_DRIVER) $(GOOSE_DBSTRING) down
resetse:
	@goose -dir=$(GOOSE_MIGRATION_DIR) $(GOOSE_DRIVER) $(GOOSE_DBSTRING) reset

sqlgen:
	@"$(USERPROFILE)\go\bin\sqlc.exe" generate

.PHONY: run downse upse resetse

.PHONY: air