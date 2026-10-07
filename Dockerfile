# syntax=docker/dockerfile:1

FROM golang:1.26.1 AS builder

ARG VERSION=dev

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

# The Makefile owns the build flags (static binary, version ldflags).
RUN GOOS=linux make build VERSION=${VERSION}

# Main image (agent: distroless).
FROM gcr.io/distroless/static-debian12:nonroot

ARG VERSION=dev
LABEL org.opencontainers.image.source="https://github.com/compliance-framework/mock-agent" \
	org.opencontainers.image.version="${VERSION}"

# Every image ships the static binary at /app/mock-agent: mock-agent-action's
# Dockerfile copies it from there and fails its build if it is missing.
COPY --from=builder /src/dist/mock-agent /app/mock-agent

ENTRYPOINT ["/app/mock-agent"]
