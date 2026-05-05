VERSION := $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
COMMIT  := $(shell git rev-parse --short HEAD 2>/dev/null || echo "none")
DATE    := $(shell date -u +%Y-%m-%dT%H:%M:%SZ)
LDFLAGS := -ldflags "-X github.com/ismt7/genesis/cmd.version=$(VERSION) \
	-X github.com/ismt7/genesis/cmd.commit=$(COMMIT) \
	-X github.com/ismt7/genesis/cmd.date=$(DATE)"

.PHONY: build run version clean

build:
	go build $(LDFLAGS) -o genesis .

run: build
	./genesis

version: build
	./genesis version

clean:
	rm -f genesis