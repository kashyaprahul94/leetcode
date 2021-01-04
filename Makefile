install:
	go mod download

tidy:
	go mod tidy

run:
	go run ./cmd/leetcode.go $(ARGS)

test:
	go test -v ./...

compile:
	GOOS=linux GOARCH=amd64 go build $(ARGS) ./cmd/leetcode.go

build: tidy compile