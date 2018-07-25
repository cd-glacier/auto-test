
deps:
	cd src && dep ensure

build:
	go build src/cmd/main.go

run:
	go run src/cmd/main.go

