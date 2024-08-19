CURRENT_DIR=$(shell pwd)

DB_URL := postgres://postgres:1@localhost:5432/auth?sslmode=disable

APP=template

run:
	go run cmd/main.go
init:
	go mod init
	go mod tidy
	go mod vendor

proto-gen:
	./scripts/gen-proto.sh ${CURRENT_DIR}

swag-init:
	swag init -g api/docs/docs.go --output api/handler/docs

mig-create-users:
	migrate create -ext sql -dir migrations -seq auth_service_table

mig-up:
	migrate -path migrations -database '${DB_URL}' -verbose up

mig-down:
	migrate -path migrations -database '${DB_URL}' -verbose down

mig-force:
	migrate -path migrations -database '${DB_URL}' -verbose force 1

mig-insert:
	migrate create -ext sql -dir migrations -seq auth_service_table

swag-gen:
	~/go/bin/swag init -g api/api.go -o api/docs

swag-gens:
	~/go/bin/swag init -g ./api/api.go -o api/docs force 1