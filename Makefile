.PHONY: help
.DEFAULT_GOAL := help

deps: ## create vendor directory with go dep.
	cd src && dep ensure

build: deps ## build main.go.
	go build src/cmd/main.go

run: ## run main.go.
	go run src/cmd/main.go

test: ## unit test.
	go test -v ./...

docker-build: ## create docker image for building code.
	docker build -t auto-test .

docker-run: ## run main.go with docker image. run after *make docker-build*
	docker run -e LOG_LEVEL=debug -v $(PWD):/go/src/github.com/g-hyoga/auto-test auto-test go run ./src/cmd/main.go

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
