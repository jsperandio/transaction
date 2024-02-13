swagger:
	swag init -g cmd/transaction/main.go

run: swagger
	go run ./cmd/transaction/main.go