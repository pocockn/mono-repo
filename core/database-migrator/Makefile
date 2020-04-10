BINARY_NAME=database-migrator
BINARY_UNIX=$(BINARY_NAME)_unix
BUILD_TIME ?= "$(date -u +"%d/%m/%YT%H:%M:%S%z")"
GOBUILD=$(GOCMD) build
GOCMD=go
GOCLEAN=$(GOCMD) clean
VERSION = $(shell cat ./VERSION)

build-time:
	echo "=> Pushing pococknick91/$(BUILD_TIME):${VERSION} to docker"

build-binary:
	@$(GOBUILD) -ldflags "-X main.Version=${VERSION} -X main.BuildTime=${BUILD_TIME}"

build-binary-mac:
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=1 $(GOBUILD) -ldflags "-X main.Version=${VERSION} -X main.BuildTime=${BUILD_TIME}"

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v

build-image:
	docker build -t pocockn/$(BINARY_NAME) --build-arg VERSION=${VERSION} .

DEFAULT: run

push-image: build-linux
	docker build -t pocockn/$(BINARY_NAME) --build-arg VERSION=${VERSION} .
	docker tag pocockn/$(BINARY_NAME) pococknick91/$(BINARY_NAME):${VERSION}
	echo "=> Pushing pococknick91/$(BINARY_NAME):${VERSION} to docker"
	docker push pococknick91/$(BINARY_NAME):${VERSION}

run:
	@$(GOBUILD) -ldflags "-X main.Version=dev -X main.BuildTime=17/01/2017T14:12:35+0000"
	@ENV=development ./$(BINARY_NAME)

vet:
	@go tool vet .

release:
	./release.sh ${VERSION}
