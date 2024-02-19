swagger:
	swag init -g cmd/transaction/main.go

local: swagger
	go run ./cmd/transaction/main.go

lint:
	golangci-lint run ./... --config ./build/golangci-lint/config.yml

test:
	go test -race ./...

docker-build:
	docker build -t transaction -f ./build/docker/dockerfile .

run:
	docker run -e CONF=./global.yaml --name transaction -p 8081:8081 -d transaction

db: 
	docker run --rm -P --name transaction-pg -p 5432:5432 -e POSTGRES_PASSWORD=pismo -e POSTGRES_USER=pismo -d postgres:16.2-alpine

up: db
	docker run -e CONF=./global.yaml --name transaction -p 8081:8081 -d transaction