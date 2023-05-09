APP_NAME=fizzbuzz
BUILD_DIR=./build
DOCKER_TAG ?= latest
DOCKER_PORT ?= 8080

.PHONY: build run test clean docker-build docker-run lint

docker-build:
	docker build -t $(APP_NAME):$(DOCKER_TAG) .

docker-run:
	docker run -p $(DOCKER_PORT):8080 $(APP_NAME):$(DOCKER_TAG)

build:
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(APP_NAME) -v ./cmd/lbc-fizzbuzz

run: build
	$(BUILD_DIR)/$(APP_NAME)

test:
	go test -v ./...

lint:
	docker run --rm -v $(PWD):/app -w /app golangci/golangci-lint:v1.52-alpine golangci-lint run

clean:
	rm -rf $(BUILD_DIR)