.DEFAULT_GOAL := build

build:
	CGO_ENABLED=0 go build -o bin/dependabot-tools -tags "${GO_TAGS}" -ldflags "-s -w" .

tidy:
	go mod tidy

lint:
	golangci-lint run

validate: tidy lint

test:
	go test -v ./pkg/...

integration: build
	docker build -t dependabot-tools .

	@for dir in test/*; do \
		if [ -d "$$dir" ]; then \
			echo "Running integration tests for $$dir"; \
			docker build -t dependabot-tools-test -f $$dir/Dockerfile .; \
		fi; \
	done
