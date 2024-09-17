dev: init
	@go run ./cmd/start.go -config ./config/dev.yaml

pro: init
	@go build ./cmd/start.go
	@./cmd/start -config ./config/pro.yaml

init: fmt
	@go mod tidy
	@swag init -g ./cmd/start.go

fmt:
	@go fmt ./...

testAI:
	@go test -v service/test/ai_test.go