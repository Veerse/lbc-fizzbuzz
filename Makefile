APP_NAME=fizzbuzz
BUILD_DIR=./build
DOCKER_TAG ?= latest
DOCKER_PORT ?= 8080

.PHONY: build test clean docker-build docker-run

docker-build:
	docker build -t $(APP_NAME):$(DOCKER_TAG) .

docker-run:
	docker run -p $(DOCKER_PORT):8080 $(APP_NAME):$(DOCKER_TAG)

build:
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(APP_NAME) -v ./cmd/lbc-fizzbuzz

test:
	go test -v ./...

clean:
	rm -rf $(BUILD_DIR)