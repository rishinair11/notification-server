APP=notification-server-go
VERSION=latest

all: clean deps build

clean:
	@echo "--> cleaning..."
	@rm -rf build
	@rm -rf vendor
	@go clean ./...

prereq:
	@echo "--> prerequisites..."
	@mkdir -p build/bin

deps: prereq
	@echo "--> installing prereqs..."
	@dep ensure

build: prereq
	@echo '--> building...'
	@go fmt ./...
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o build/bin/${APP} src/server.go

package: build
	@echo '--> building docker image...'
	@docker build -t ${APP}:${VERSION} .