VERSION := 0.1
BUILDDIR := build
ENTRYPOINT := github.com/cderwin/cryptopals

.PHONY: all build clean

all: build

build:
	mkdir -p $(BUILDDIR) && \
	GOOS=darwin ARCH=amd64 go build -o $(BUILDDIR)/cryptopals $(ENTRYPOINT)

clean:
	@rm -rf $(BUILDDIR)
