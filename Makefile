POSTGRESQL_URL = postgres://postgres:0000@localhost:5432/conduit?sslmode=disable

migrate-create:
	migrate create -ext sql -dir migrations/ -seq init

migrate-up:
	migrate -database postgres://postgres:0000@localhost:5432/conduit?sslmode=disable -path migrations up

migrate-down:
	migrate -database postgres://postgres:0000@localhost:5432/conduit?sslmode=disable -path migrations down

run:
	go run ./cmd/app/main.go

.PHONY: migrate-create migrate-up migrate-down run