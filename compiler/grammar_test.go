package main

import (
	"testing"

	"compiler/lexer"
	"compiler/parser"
)

func TestParseSample(t *testing.T) {
	srcs := []string{
		`program demo;

		var  x, y, z : int;

		main { 
			print(1 + 2); 
		}

		end`,
		`program demo;

		var  x, y, z : int;

		main { 
			print(1 + 2); 
		}`,
	}

	for _, src := range srcs {
		l := lexer.NewLexer([]byte(src))
		p := parser.NewParser()

		tree, perr := p.Parse(l)

		if perr != nil {
			t.Fatalf("parse failed: %v", perr)
		}

		t.Logf("Parse OK, top-level attribute = %#v", tree)
	}
}
