package main

import (
	"testing"

	"compiler/ast"
	"compiler/lexer"
	"compiler/parser"
)

// ------------------------ Expression Tests ------------------------

func TestCorrectlyGeneratesQuadruplesForExpressions(t *testing.T) {
	src :=
		`program demoEleven;

		var a, b, c, z : int;

		main {
			a = a + b + c + z;
		}
		end`

	l := lexer.NewLexer([]byte(src))
	p := parser.NewParser()

	tree, perr := p.Parse(l)

	PrintFunctionMapWithVars(ast.ProgramFunctions)

	PrintQuadrupleList(ast.QuadrupleList)

	if perr != nil {
		t.Fatalf("parse failed: %v", perr)
	}

	t.Logf("parse OK %#v", tree)
}

func TestCorrectlyGeneratesQuadruplesForExpressionsWithFunctions(t *testing.T) {
	src :=
		`program demoTwelve;

		var a, b, c, z : int;

		void anotherFunction(p : int, q : float) [
			var o : int;

			{
				o = p + q;
				print(o, "Hello, World", "Diego Gtz");
			}
		];

		main {

		}

		end`

	l := lexer.NewLexer([]byte(src))
	p := parser.NewParser()

	tree, perr := p.Parse(l)

	PrintFunctionMapWithVars(ast.ProgramFunctions)

	PrintQuadrupleList(ast.QuadrupleList)

	if perr != nil {
		t.Fatalf("parse failed: %v", perr)
	}

	t.Logf("parse OK %#v", tree)
}

// ------------------------ Statement Tests ------------------------

func TestCorrectlyGeneratesQuadruplesForIfStatement(t *testing.T) {
	src :=
		`program demoThirteen;

		var x, y, z : int;

		main {
			if (x != 5 * 3 + (4*3)) {
				x = 10;

				if (x != 10) {
					x = 10;
				} else {
					x = 20;
				};
			} else {
				x = 20;
			};
		}

		end`

	l := lexer.NewLexer([]byte(src))
	p := parser.NewParser()

	tree, perr := p.Parse(l)

	PrintFunctionMapWithVars(ast.ProgramFunctions)

	PrintQuadrupleList(ast.QuadrupleList)

	if perr != nil {
		t.Fatalf("parse failed: %v", perr)
	}

	t.Logf("parse OK %#v", tree)
}

// Error, doesn't jump to correct label
func TestCorrectlyGeneratesQuadruplesForWhileStatement(t *testing.T) {
	src :=
		`program demoFourteen;

		var x, y, z : int;

		main {
			print(x);

			while (x != (x * 2)) do {
				x = x + 1;
			};
		}

		end`

	l := lexer.NewLexer([]byte(src))
	p := parser.NewParser()

	tree, perr := p.Parse(l)

	PrintFunctionMapWithVars(ast.ProgramFunctions)

	PrintQuadrupleList(ast.QuadrupleList)

	if perr != nil {
		t.Fatalf("parse failed: %v", perr)
	}

	t.Logf("parse OK %#v", tree)
}

func TestCorrectlyGeneratesQuadruplesForPrintStatement(t *testing.T) {
	src :=
		`program demoFifteen;

		var x, y, z : int;

		main {
			print(x, "Hello, World", z);
		}

		end`

	l := lexer.NewLexer([]byte(src))
	p := parser.NewParser()

	tree, perr := p.Parse(l)

	PrintFunctionMapWithVars(ast.ProgramFunctions)

	PrintQuadrupleList(ast.QuadrupleList)

	if perr != nil {
		t.Fatalf("parse failed: %v", perr)
	}

	t.Logf("parse OK %#v", tree)
}
