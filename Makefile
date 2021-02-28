PREFIX?=$(shell pwd)
BINDIR := ${PREFIX}/bin

.PHONY: all
all: clean test build

.PHONY: build
build:
	@echo "==> $@ <=="
	@echo "Building binary..."
	@GO111MODULE=on go build -trimpath -o $(BINDIR)/waterjug ./cmd/main.go
	@echo "Binary created on: ./bin/waterjug"

test:
	@echo "==> $@ <=="
	@echo "Running tests..."
	@GO111MODULE=on go test -v ./...

run:
	@echo "==> $@ <=="
	@echo "Running application..."
	@GO111MODULE=on go run ./cmd/main.go

.PHONY: clean
clean:
	@echo "==> $@ <=="
	@echo "Cleaning..."
	@$(RM) -r $(BUILDDIR)
	@$(RM) -r $(BINDIR)
