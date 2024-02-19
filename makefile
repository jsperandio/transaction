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

db: 
	docker run --rm -P --name transaction-pg -v ./build/docker/schema.sql:/docker-entrypoint-initdb.d/schema.sql -p 172.17.0.1:5432:5432 -e POSTGRES_PASSWORD=pismo -e POSTGRES_USER=pismo   -d postgres:16.2-alpine

up: db
	docker run -e CONF=./global.yaml --name transaction -p 8081:8081 -d transaction

dc-build:
	docker-compose -f ./build/docker/docker-compose.yaml build

dc-up:
	docker-compose -f ./build/docker/docker-compose.yaml up -d

dc-down:
	docker-compose -f ./build/docker/docker-compose.yaml down

dc-logs:
	docker-compose -f ./build/docker/docker-compose.yaml logs