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
	expr, err := parser.ParseString("", input)

	if err != nil {
		panic(err)
	}

	result := expr.Left + expr.Right
	fmt.Printf("%d %s %d = %d\n", expr.Left, expr.Op, expr.Right, result)
}
