package main

import (
	"testing"

	"compiler/ast"
	"compiler/lexer"
	"compiler/parser"
)

func TestTheParserCorrectlyParsesAModule(t *testing.T) {
	src :=
		`program demoOne;

		void testFunction() [
			{
				print("Hello, World!");
			}
		];

		main {
			testFunction();
		}
		end`

	l := lexer.NewLexer([]byte(src))
	p := parser.NewParser()

	tree, perr := p.Parse(l)

	PrintQuadrupleList(ast.QuadrupleList)

	if perr != nil {
		t.Fatalf("parse failed: %v", perr)
	}

	t.Logf("parse OK %#v", tree)
}

func TestTheParserCorrectlyParsesTwoModules(t *testing.T) {
	src :=
		`program demoTwo;

		void testFunction() [
			{
				print("Hello, World!");
			}
		];

		void testFunctio2() [
			{
				testFunction();
			}
		];

		main {
			testFunctio2();
		}
		end`

	l := lexer.NewLexer([]byte(src))
	p := parser.NewParser()

	tree, perr := p.Parse(l)

	PrintQuadrupleList(ast.QuadrupleList)

	if perr != nil {
		t.Fatalf("parse failed: %v", perr)
	}

	t.Logf("parse OK %#v", tree)
}

func TestTheParserCorrectlyParsesParameters(t *testing.T) {
	src :=
		`program demoThree;

		void testFunction(a : int, b : float) [
			{
				print(a, b);
			}
		];

		main {
			testFunction(1, 2.0);
		}
		end`

	l := lexer.NewLexer([]byte(src))
	p := parser.NewParser()

	tree, perr := p.Parse(l)

	PrintQuadrupleList(ast.QuadrupleList)

	if perr != nil {
		t.Fatalf("parse failed: %v", perr)
	}

	t.Logf("parse OK %#v", tree)
}

func TestTheParserDetectsIncorrectParameters(t *testing.T) {
	src :=
		`program demoFour;

		void testFunction(a : int, b : float) [
			{
				print(a, b);
			}
		];

		main {
			testFunction(1);
		}
		end`

	l := lexer.NewLexer([]byte(src))
	p := parser.NewParser()

	tree, perr := p.Parse(l)

	PrintQuadrupleList(ast.QuadrupleList)

	if perr == nil {
		t.Fatalf("expected parse to fail but it succeeded")
	}

	t.Logf("parse OK %#v", tree)
}
