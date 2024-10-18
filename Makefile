TARGET_BIN = futil
GO_CMD_MAIN = cmd/app/main.go

run:
	go run $(GO_CMD_MAIN)

run-build:
	go build -o $(TARGET_BIN) $(GO_CMD_MAIN)

run-test:
	go test ./... -count=1 -cover -race -v
