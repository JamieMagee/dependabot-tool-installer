.DEFAULT_GOAL := build

build:
	CGO_ENABLED=0 go build -o bin/dependabot-tools -tags "${GO_TAGS}" -ldflags "-s -w" .

docker:
	docker build -t ghcr.io/jamiemagee/dependabot-tools .

tidy:
	go mod tidy

lint:
	golangci-lint run

validate: tidy lint

unit-test:
	go test -v ./...

integration-test: docker
	@for dir in test/*; do \
		if [ -d "$$dir" ]; then \
			echo "Running integration tests for $$dir"; \
			docker build -t dependabot-tools-test -f $$dir/Dockerfile .; \
		fi; \
	done

test: unit-test integration-test
