.DEFAULT_GOAL := build

build:
	CGO_ENABLED=0 go build -o bin/dependabot-tools -tags "${GO_TAGS}" -ldflags "-s -w" .

tidy:
	go mod tidy

test:
	go test -v ./pkg/...