dev:
	@go mod tidy
	@go fmt ./...
	@swag init
	@go run ./main.go -config ./config/dev.yaml

pro:
	@go mod tidy
	@go fmt ./...
	@swag init
	@go run ./main.go -config ./config/pro.yaml

fmt:
	@go fmt ./...