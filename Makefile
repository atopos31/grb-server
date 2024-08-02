dev: init
	@go run ./main.go -config ./config/dev.yaml

pro: init
	@go build ./main.go
	@./main -config ./config/pro.yaml

init:
	@go mod tidy
	@swag init

fmt:
	@go fmt ./...