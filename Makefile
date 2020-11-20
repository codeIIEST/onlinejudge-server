-include ./env
GOCMD=go 

GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GORUN=$(GOCMD) run
GOFILES=./server/*.go
MIGRATIONFILES=./migrations/*.go
BINARY_PATH=./bin/
BINARY_NAME=server
BINARY_UNIX=$(BINARY_NAME)_unix

all: test build

migrate:
		$(GORUN) $(MIGRATIONFILES)
build: 
		GOBIN=$(GOBIN) $(GOBUILD) -o $(BINARY_PATH)$(BINARY_NAME) -v $(GOFILES)
test: 
		$(GOTEST) -v ./...
clean: 
		$(GOCLEAN) $(GOFILES)
		rm -f $(BINARY_PATH)$(BINARY_NAME)
		rm -f $(BINARY_PATH)$(BINARY_UNIX)
run:
		$(GOBUILD) -o $(BINARY_PATH)$(BINARY_NAME) $(GOFILES)
		./$(BINARY_PATH)$(BINARY_NAME)

# Cross compilation
build-linux:
		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_PATH)$(BINARY_UNIX) -v
docker-build:
		docker run --rm -it -v "$(GOPATH)":/go -w /go/src/bitbucket.org/rsohlich/makepost golang:latest go build -o "$(BINARY_UNIX)" -v