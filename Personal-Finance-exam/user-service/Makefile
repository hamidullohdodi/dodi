CURRENT_DIR := $(shell pwd)

DATABASE_URL ?= postgres://postgres:1@localhost:5432/auth

run:
	@go run cmd/main.go

tidy:
	@go mod tidy

mig-create:
	@if [ -z "$(name)" ]; then \
	read -p "Enter migration name: " name; \
	fi; \
	migrate create -ext sql -dir db/migrations -seq $$name

mig-up:
	@migrate -database "$(DATABASE_URL)" -path db/migrations up

mig-down:
	@migrate -database "$(DATABASE_URL)" -path db/migrations down

swag-gen:
	~/go/bin/swag init -g api/api.go -o api/docs
swag-init:
	swag init -g api/docs/docs.go --output api