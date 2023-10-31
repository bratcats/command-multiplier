EXEC_NAME=cmd-multp
PRJ_DIR=$(shell pwd)
BUILD_DIR=$(PRJ_DIR)/bin
all: tidy vet build

tidy:
	go mod tidy

vet:
	go vet ./...

build: clean
	mkdir $(BUILD_DIR)
	go build -C cmd -o $(BUILD_DIR)/$(EXEC_NAME)

clean:
	rm -rf $(BUILD_DIR)