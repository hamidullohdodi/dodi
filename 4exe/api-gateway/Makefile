CURRENT_DIR=$(shell pwd)

run:
	@go run cmd/main.go

tidy:
	@go mod tidy

mig-create:
	@if [ -z "$(name)" ]; then \
	read -p "Enter migration name: " name; \
	fi; \
	migrate create -ext sql -dir migrations -seq $$name

mig-up:
	@migrate -database "$(DATABASE_URL)" -path db/migrations up

mig-down:
	@migrate -database "$(DATABASE_URL)" -path db/migrations down


SWAGGER := ~/go/bin/swag
SWAGGER_DOCS := api/docs
SWAGGER_INIT := $(SWAGGER) init -g api/api.go -o $(SWAGGER_DOCS)

# Target to generate swagger documentation
swag-gen:
	$(SWAGGER_INIT)
