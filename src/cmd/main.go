package main

import (
	"flag"
	"go/parser"
	"go/token"
	"log"

	"github.com/k0kubun/pp"
)

func main() {
	var (
		target = flag.String("target", "", "target file")
	)
	flag.Parse()

	fs := token.NewFileSet()
	f, err := parser.ParseFile(fs, *target, nil, 0)
	if err != nil {
		log.Println(err)
	}
	pp.Println(f)
}
