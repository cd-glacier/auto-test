package main

import (
	"errors"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"strconv"

	"github.com/k0kubun/pp"
)

func evalPlus(expr *ast.BinaryExpr) (int64, error) {
	xList, ok := expr.X.(*ast.BasicLit)
	if !ok {
		return 0, errors.New("left operand is not BasicLit")
	}

	yList, ok := expr.Y.(*ast.BasicLit)
	if !ok {
		return 0, errors.New("right operand is not BasicLit")
	}

	if expr.Op != token.ADD {
		return 0, errors.New("operator is not +")
	}

	x, err := strconv.ParseInt(xList.Value, 10, 64)
	if err != nil {
		return 0, err
	}

	y, err := strconv.ParseInt(yList.Value, 10, 64)
	if err != nil {
		return 0, err
	}

	return x + y, nil
}

func main() {
	expr, err := parser.ParseExpr("1+2")
	if err != nil {
		log.Println(err)
	}

	if be, ok := expr.(*ast.BinaryExpr); ok {
		if v, err := evalPlus(be); err == nil {
			pp.Println(v)
		} else {
			pp.Printf("error: %s", err)
		}
	}

}
