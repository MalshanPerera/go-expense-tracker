build:
	@go build -o bin/main cmd/main.go

dev:
	@go build -o ./tmp/$(APP_NAME) ./cmd/$(APP_NAME)/main.go && air