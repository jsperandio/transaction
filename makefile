swagger:
	swag init -g cmd/transaction/main.go

local: swagger
	go run ./cmd/transaction/main.go

lint:
	golangci-lint run ./... --config ./build/golangci-lint/config.yml

test:
	go test -race ./...

build-docker:
	docker build -t transaction -f ./build/docker/dockerfile .

up:
	docker run -p 8081:8081  transaction

run: build-docker up