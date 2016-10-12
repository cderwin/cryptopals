VERSION := 0.1
BUILDDIR := build
ENTRYPOINT := github.com/cderwin/cryptopals

OS := linux

.PHONY: all build clean

all: build

build:
	mkdir -p $(BUILDDIR) && \
	go build -o $(BUILDDIR)/cryptopals $(ENTRYPOINT)

clean:
	@rm -rf $(BUILDDIR)
