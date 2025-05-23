package main

import (
	"testing"

	"compiler/ast"
	"compiler/lexer"
	"compiler/parser"
)

// ------------------------ Parser Tests ------------------------
// This test checks if the parser correctly parses a correct sample
func TestTheParserCorrectlyParsesACorrectSample(t *testing.T) {
	src :=
		`program demoOne;

		var  x, y, z : int;

		main {
			while (x < 10) do {
				print(x);

				if (x != 5) {
					x = 10;
				};
			};
		}
		end`

	l := lexer.NewLexer([]byte(src))
	p := parser.NewParser()

	tree, perr := p.Parse(l)

	if perr != nil {
		t.Fatalf("parse failed: %v", perr)
	}

	PrintFunctionMapWithVars(ast.ProgramFunctions)
	PrintQuadrupleList(ast.QuadrupleList)

	t.Logf("parse OK %#v", tree)
}

// This test checks if the parser correctly parses another correct sample
func TestTheParserCorrectlyParsesAnotherCorrectSample(t *testing.T) {
	src :=
		`program demoTwo;

		var  x, y, z : int; p, e, o : float;

		void aFunction(a : int, b : float) [
			var c : int;

			{
				x = 1;
				y = 2;
				x = x + y;
				print(x);
			}
		];

		void anotherFunction(a : int, b : float) [
			var c : int;

			{
				a = 1;
				b = 2;
				c = a + b;
				print(c);
			}
		];

		main {
			aFunction(1, 2.0);

			while (x < 10) do {
				print(x);
				x = x + p;
			};
		}

		end`

	l := lexer.NewLexer([]byte(src))
	p := parser.NewParser()

	tree, perr := p.Parse(l)

	if perr != nil {
		t.Fatalf("parse failed: %v", perr)
	}

	PrintFunctionMapWithVars(ast.ProgramFunctions)
	PrintQuadrupleList(ast.QuadrupleList)

	t.Logf("parse OK %#v", tree)
}

// This test checks if the parser correctly detects a missing end
func TestTheParserDetectsMissingEnd(t *testing.T) {
	src :=
		`program demoThree;

		var  x, y, z : int;

		main {
			print(1 + 2);
		}`

	l := lexer.NewLexer([]byte(src))
	p := parser.NewParser()

	_, perr := p.Parse(l)

	if perr == nil {
		t.Fatalf("expected parse to fail but it succeeded")
	}

	t.Logf("parse correctly failed with: %v", perr)
}

// This test checks if the parser correctly detects a missing semicolon, should FAIL with that error
func TestTheParserDetectsMissingSemicolon(t *testing.T) {
	src :=
		`program demoFour

		var  x, y, z : int;

		main {
			print(1 + 2);
		}`

	l := lexer.NewLexer([]byte(src))
	p := parser.NewParser()

	_, perr := p.Parse(l)

	if perr == nil {
		t.Fatalf("expected parse to fail but it succeeded")
	}

	t.Logf("parse correctly failed with: %v", perr)
}

// This test checks if the parser correctly detects a wrong token, should FAIL with that error
func TestTheParserDetectsWrongTokens(t *testing.T) {
	src :=
		`program demoFive;

		var  x, y, z : string;

		main {
			print(1 + 2);
		}`

	l := lexer.NewLexer([]byte(src))
	p := parser.NewParser()

	_, perr := p.Parse(l)

	if perr == nil {
		t.Fatalf("expected parse to fail but it succeeded")
	}

	t.Logf("parse correctly failed with: %v", perr)
}

// ------------------------ Variable / Function Definition Tests ------------------------
// This test checks if the parser correctly detects a global variable redeclaration, should FAIL with that error

func TestASTDetectsAllFunctionsAndVariablesCorrectly(t *testing.T) {
	src :=
		`program demoSix;

		var x, y, z : int;

		void testFunction(a : int, b : float) [
			var c : int;

			{
				c = 1 + 2;
				print(c);
			}
		];

		void anotherFunction(abc : float, bca : int) [
			var c : float;

			{
				c = 1 + 2;
				print(c);
			}
		];

		main {
			x = 1;
			anotherFunction(1, 2.0);
		}

		end`

	l := lexer.NewLexer([]byte(src))
	p := parser.NewParser()

	tree, perr := p.Parse(l)

	if perr != nil {
		t.Fatalf("parse failed: %v", perr)
	}

	PrintFunctionMapWithVars(ast.ProgramFunctions)
	PrintQuadrupleList(ast.QuadrupleList)
	t.Logf("parse OK %#v", tree)
}

func TestASTDetectsGlobalVariableRedeclaration(t *testing.T) {
	src :=
		`program demoSeven;

		var x, y, z : int;

		void anotherFunction(a : int, b : float) [
			var x : int;

			{
				d = a + b;
				print(d);
			}
		];

		main {
			print(x);
		}

		end`

	l := lexer.NewLexer([]byte(src))
	p := parser.NewParser()

	_, perr := p.Parse(l)

	if perr == nil {
		t.Fatalf("expected parse to fail but it succeeded")
	}

	t.Logf("parse correctly failed with: %v", perr)
}

func TestASTDetectsFunctionRedeclaration(t *testing.T) {
	src :=
		`program demoEight;

		var x, y, z : int;

		void anotherFunction(a : int, b : float) [
			{
				x = 2;
			}
		];

		void anotherFunction(abc : float, bca : int) [
			var c : float;

			{
				c = 3;
			}
		];

		main {
			print(x);
		}

		end`

	l := lexer.NewLexer([]byte(src))
	p := parser.NewParser()

	_, perr := p.Parse(l)

	PrintFunctionMapWithVars(ast.ProgramFunctions)

	if perr == nil {
		t.Fatalf("expected parse to fail but it succeeded")
	}

	t.Logf("parse correctly failed with: %v", perr)
}

func TestASTDetectsUndefinedVariable(t *testing.T) {
	src :=
		`program demoNine;

			var x, y, z : int;

		main {
			print(a);
		}

		end`

	l := lexer.NewLexer([]byte(src))
	p := parser.NewParser()

	_, perr := p.Parse(l)

	PrintFunctionMapWithVars(ast.ProgramFunctions)

	if perr == nil {
		t.Fatalf("expected parse to fail but it succeeded")
	}

	t.Logf("parse correctly failed with: %v", perr)
}

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
