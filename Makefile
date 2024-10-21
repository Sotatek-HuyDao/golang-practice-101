TARGET_BIN = futil
GO_CMD_MAIN = cmd/app/main.go
VERSION = 0.0.1

run:
	go run $(GO_CMD_MAIN)

run-build:
	go build -ldflags="-X 'github.com/Sotatek-HuyDao/golang-practice-101/internal/version.version=$(VERSION)'" -o $(TARGET_BIN) $(GO_CMD_MAIN)

run-test:
	go test ./... -count=1 -cover -race -v
