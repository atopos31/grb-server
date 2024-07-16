dev:
	@go mod tidy
	@swag init
	@go fmt ./...
	@go run ./main.go -config ./config/dev.yaml

pro:
	@go mod tidy
	@swag init
	@go fmt ./...
	@go build ./main.go
	@./main.go -config ./config/pro.yaml

fmt:
	@go fmt ./...