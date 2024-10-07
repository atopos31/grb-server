dev: init
	@go run ./cmd/start.go -config ./config/dev.yaml

pro: init
	@go build -o ./bin/grb ./cmd/start.go 

init: fmt
	@go mod tidy
	@swag init -g ./cmd/start.go

fmt:
	@go fmt ./...

testAI:
	@go test -v service/test/ai_test.go