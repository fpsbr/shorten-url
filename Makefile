GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get

BINARY_NAME=bin/url-shortener

MAIN_FILE=url-shortener/cmd/api

build:
	$(GOBUILD) -o $(BINARY_NAME) $(MAIN_FILE)

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

deps:
	$(GOGET) -v ./...

run:
	$(GOBUILD) -o $(BINARY_NAME) $(MAIN_FILE)
	./$(BINARY_NAME)

default: build

.PHONY: build clean test deps run
