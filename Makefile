POSTGRESQL_URL = postgres://root:secret@localhost:5432/conduit?sslmode=disable

migrate-create:
	migrate create -ext sql -dir migrations/ -seq init

migrate-up:
	migrate -database ${POSTGRESQL_URL} -path migrations up

migrate-down:
	migrate -database ${POSTGRESQL_URL} -path migrations down

run:
	go run ./cmd/app/main.go

.PHONY: migrate-create migrate-up migrate-down run