EXEC_NAME=cmd-multp
BUILD_DIR=bin
all: tidy vet build

tidy:
	go mod tidy

vet:
	go vet ./...

build: clean
	mkdir $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(EXEC_NAME)

clean:
	rm -rf $(BUILD_DIR)