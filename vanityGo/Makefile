.PHONY: build run clean install

build:
	go build -o vanity-worker .

run:
	go run .

install:
	go mod download

clean:
	rm -f vanity-worker vanity-worker-linux

build-linux:
	GOOS=linux GOARCH=amd64 go build -o vanity-worker-linux .

docker-build:
	docker build -t vanity-worker-go .

help:
	@echo "Available targets:"
	@echo "  build        - Build the binary"
	@echo "  run          - Run the application"
	@echo "  install      - Download dependencies"
	@echo "  clean        - Remove built binaries"
	@echo "  build-linux  - Build for Linux (cross-compile)"
