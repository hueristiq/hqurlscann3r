GOCMD=go
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
GOTEST=$(GOCMD) test
GOBUILD=$(GOCMD) build

build:
	$(GOBUILD) -v -ldflags="-extldflags=-static" -o "sigurlscann3r" cmd/sigurlscann3r/main.go

test: 
	$(GOTEST) -v ./...

tidy:
	$(GOMOD) tidy