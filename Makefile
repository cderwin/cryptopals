VERSION := 0.1
BUILDDIR := build
ENTRYPOINT := github.com/cderwin/cryptopals

OS := linux

.PHONY: all build clean test

all: build

build:
	mkdir -p $(BUILDDIR) && \
	go build -o $(BUILDDIR)/cryptopals $(ENTRYPOINT)

test:
	go test -v ./...

clean:
	@rm -rf $(BUILDDIR)
