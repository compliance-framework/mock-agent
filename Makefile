VERSION ?= dev
LDFLAGS := -s -w -X main.buildVersion=$(VERSION)

.PHONY: build test docker-build

build:
	CGO_ENABLED=0 go build -trimpath -ldflags "$(LDFLAGS)" -o dist/mock-agent .

test:
	go test ./... -coverprofile cover.out

# Builds all three images: mock-agent, mock-agent-ci and mock-agent-custodian.
docker-build:
	docker build --build-arg VERSION=$(VERSION) -f Dockerfile -t mock-agent:$(VERSION) .
	docker build --build-arg VERSION=$(VERSION) -f Dockerfile-ci -t mock-agent-ci:$(VERSION) .
	docker build --build-arg VERSION=$(VERSION) -f Dockerfile-custodian -t mock-agent-custodian:$(VERSION) .
