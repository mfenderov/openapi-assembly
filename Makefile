NAME=openapi-assembler
REPO=github.com/mfenderov/${NAME}

dependencies:
	@go mod vendor
	@go mod tidy
clean:
	@rm -rf build/
test:
	@go test -v -race ./...
build:
	@go build -o build/ -v ./...
install:
	@go install ${REPO}
