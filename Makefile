deps:
	cd src && dep ensure

build: deps
	go build src/cmd/main.go

run:
	go run src/cmd/main.go

test:
	go test -v ./...

docker-build:
	docker build -t auto-test .

docker-run:
	docker run -e LOG_LEVEL=debug -v $(PWD):/go/src/github.com/g-hyoga/auto-test auto-test go run ./src/cmd/main.go
