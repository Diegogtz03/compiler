package main

import (
	"testing"

	"compiler/ast"
	"compiler/lexer"
	"compiler/parser"
	"compiler/types"
)

// ------------------------ Variable / Function Definition Tests ------------------------

func TestASTDetectsAllFunctionsAndVariablesCorrectly(t *testing.T) {
	src :=
		`program demoSix;

		var x : int;

		void testFunction(a : int, b : float) [
			var c : int;
			{
				c = 1 + 2;
			}
		];

		main {
			testFunction(5, 2.0);
		}

		end`

	l := lexer.NewLexer([]byte(src))
	p := parser.NewParser()

	tree, perr := p.Parse(l)

	if perr != nil {
		t.Fatalf("parse failed: %v", perr)
	}

	expectedFunctionMap := map[string]ast.Function{
		"demoSix": {
			Name: "demoSix",
			Vars: map[string]ast.Variable{
				"x": {
					Id:          "x",
					Type:        types.Int,
					MemoryIndex: 0,
				},
			},
		},
		"testFunction": {
			Name: "testFunction",
			Vars: map[string]ast.Variable{
				"a": {
					Id:          "a",
					Type:        types.Int,
					MemoryIndex: 0,
				},
				"b": {
					Id:          "b",
					Type:        types.Float,
					MemoryIndex: 1,
				},
				"c": {
					Id:          "c",
					Type:        types.Int,
					MemoryIndex: 2,
				},
			},
		},
	}

	PrintFunctionMapWithVars(ast.ProgramFunctions)
	PrintQuadrupleList(ast.QuadrupleList)

	if !AssertFunctionAndVarsMap(ast.ProgramFunctions, expectedFunctionMap) {
		t.Fatalf("function map mismatch")
	}

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

	PrintFunctionMapWithVars(ast.ProgramFunctions)

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
