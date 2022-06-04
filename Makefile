compose-up: ### Run docker-compose
	docker-compose up --build -d postgres && docker-compose logs -f
.PHONY: compose-up

compose-up-integration-test: ### Run docker-compose with integration test
	docker-compose up --build --abort-on-container-exit --exit-code-from integration
.PHONY: compose-up-integration-test

compose-down: ### Down docker-compose
	docker-compose down --remove-orphans
.PHONY: compose-down

linter-golangci: ### check by golangci linter
	golangci-lint run
.PHONY: linter-golangci

test: ### run test
	go test -v -cover -race ./domain/...
.PHONY: test

integration-test: ### run integration-test
	go clean -testcache && go test -v ./integration-test/...
.PHONY: integration-test

mock: ### run mockery
	mockery --all -r --case snake
.PHONY: mock

migrate-create:  ### create new migration
	migrate create -ext sql -dir migrations 'migrate_name'
.PHONY: migrate-create

init: ### init project
	rm -rf go.mod
	go mod init order-app

tidy: ### tidy the project
	rm -rf go.sum
	go mod tidy

swag-v1: ### swag init
	swag init -g domain/controller/http/v1/router.go
.PHONY: swag-v1

migration:
	go run ./cmd/app --mode migration

migration-down:
	go run ./cmd/app --mode migration-down

run: swag-v1
	go run ./cmd/app