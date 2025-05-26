package main

import (
	"compiler/VM"
	"compiler/ast"
	"compiler/lexer"
	"compiler/parser"
	"fmt"
)

func main() {
	src :=
		`program demoSix;

		var x : int;

		main {
			x = 3;
			print("Hello, World!", "And Diego", x);
		}

		end`

	l := lexer.NewLexer([]byte(src))
	p := parser.NewParser()

	_, perr := p.Parse(l)

	if perr != nil {
		fmt.Printf("parse failed: %s\n", perr)
	}

	VM.RunBabyDuck(ast.QuadrupleList)
}
