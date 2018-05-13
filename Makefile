APP_NAME=triage
BUILD_DIR=cmd/$(APP_NAME)
BINARY_PATH=$(BUILD_DIR)/$(APP_NAME)
REPO_PATH=github.com/nylar/$(APP_NAME)

all: build

build:
	@ go build -o $(BINARY_PATH) $(REPO_PATH)/$(BUILD_DIR)

run:
	@ ./$(BINARY_PATH)

test:
	@ go test ./... -v -race

clean-binary:
	@ rm $(BINARY_PATH)

clean: clean-binary

.PHONY: clean test
