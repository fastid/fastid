BINARY_NAME=fastid

.PHONY: help
help:
	@echo ""
	@echo "make build     - Building a project"
	@echo "make run       - Project run"
	@echo "make clean     - Cleans up temporary files"
	@echo "make test      - Runs tests"
	@echo "make cover     - Creates a test coverage reporter"
	@echo ""

.PHONY: build
build:
	go build -o ${BINARY_NAME} cmd/${BINARY_NAME}.go

.PHONY: run
run: build
	./${BINARY_NAME} && make clean

.PHONY: clean
clean:
	go clean
	rm -f ${BINARY_NAME}
	rm -f coverage.out
	rm -f coverage.html
	rm -f gen

.PHONY: test
test:
	go test ./...

.PHONY: cover
cover:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html
