package main

import (
	"fmt"

	"github.com/alecthomas/participle/v2"
)

type Expr struct {
	Left  int    `parser:"@Int"`
	Op    string `parser:"@('+' | '-' | '*' | '/')"`
	Right int    `parser:"@Int"`
}

func main() {
	parser, err := participle.Build[Expr]()
	if err != nil {
		panic(err)
	}

	input := "3 + 4"
	ast := &Expr{}
	_, err = parser.ParseString("", input, ast)
	if err != nil {
		panic(err)
	}

	result := ast.Left + ast.Right
	fmt.Printf("%d %s %d = %d\n", ast.Left, ast.Op, ast.Right, result)
}
