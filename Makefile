.PHONY: all
all: clean dist unit_test integration_test build package

.PHONY: dist
dist:
	mkdir dist

.PHONY: clean
clean:
	rm -rf ./dist

.PHONY: build
build:
	go build -o ./dist/server ./cmd/server

.PHONY: unit_test
unit_test:
	go test -v -cover ./...

.PHONY: integration_test
integration_test:
	go test -v -count=1 --tags=integration ./app

IMAGE=vollivier>/vehicle-server
TAG?=dev

.PHONY: package
package:
	docker build -t $(IMAGE):$(TAG) .