dev: init
	@go run ./cmd/start.go -config ./config/dev.yaml

pro: init 
	@go build -o ./bin/grb ./cmd/start.go 
	@nohup ./bin/grb -config ./config/pro.yaml > /dev/null 2>&1 &

init: fmt
	@go mod tidy
	@wire ./service
	@swag init -g ./cmd/start.go

fmt:
	@go fmt ./...

testAI:
	@go test -v service/test/ai_test.go