APP_NAME=api

run:
	@go run main.go
build:
	@go build -o $(APP_NAME) main.go
