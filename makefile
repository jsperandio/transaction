swagger:
	swag init -g cmd/transaction/main.go

lint:
	golangci-lint run ./... --config ./build/golangci-lint/config.yml

run: swagger
	go run ./cmd/transaction/main.go