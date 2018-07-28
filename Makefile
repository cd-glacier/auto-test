deps:
	cd src && dep ensure

build: deps
	go build src/cmd/main.go

run:
	go run src/cmd/main.go

test:
	go test -v ./...

mutation-deps:
	go get -t -v github.com/zimmski/go-mutesting/...

go-mutesting:
	go-mutesting ./src/cmd/main.go

