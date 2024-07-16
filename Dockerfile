FROM docker.io/library/golang:1.22 AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN make

FROM docker.io/library/ubuntu:24.04

RUN apt-get update \
  && apt-get install -y \
    ca-certificates \
    openssl \
  && rm -rf /var/lib/apt/lists/*

COPY --from=builder /build/bin/dependabot-tools /usr/local/bin/dependabot-tools

CMD ["/usr/local/bin/dependabot-tools"]