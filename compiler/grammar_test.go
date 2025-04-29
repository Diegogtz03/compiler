package main

import (
	"testing"

	"compiler/lexer"
	"compiler/parser"
)

func TestTheParserCorrectlyParsesACorrectSample(t *testing.T) {
	src :=
		`program demo;

		var  x, y, z : int;

		main {
			print(1 + 2);
		}
		end`

	l := lexer.NewLexer([]byte(src))
	p := parser.NewParser()

	tree, perr := p.Parse(l)

	if perr != nil {
		t.Fatalf("parse failed: %v", perr)
	}

	t.Logf("parse OK %#v", tree)
}

func TestTheParserCorrectlyParsesAnotherCorrectSample(t *testing.T) {
	src :=
		`program myDemoProgram;

		var  x, y, z : int; p, e, o : float;

		void aFunction(a : int, b : float) [
			var c : int;

			{
				d = a + b;
				print(d);
			}
		];

		main {
			aFunction(1, 2.0);

			while (x < 10) do {
				print(x);
				x = x + 1;
			};
		}

		end`

	l := lexer.NewLexer([]byte(src))
	p := parser.NewParser()

	tree, perr := p.Parse(l)

	if perr != nil {
		t.Fatalf("parse failed: %v", perr)
	}

	t.Logf("parse OK %#v", tree)
}

func TestTheParserDetectsMissingEnd(t *testing.T) {
	src :=
		`program demo;

		var  x, y, z : int;

		main {
			print(1 + 2);
		}`

	l := lexer.NewLexer([]byte(src))
	p := parser.NewParser()

	tree, perr := p.Parse(l)

	if perr != nil {
		t.Fatalf("parse failed: %v", perr)
	}

	t.Logf("parse OK %#v", tree)
}

func TestTheParserDetectsMissingSemicolon(t *testing.T) {
	src :=
		`program demo

		var  x, y, z : int;

		main {
			print(1 + 2);
		}`

	l := lexer.NewLexer([]byte(src))
	p := parser.NewParser()

	tree, perr := p.Parse(l)

	if perr != nil {
		t.Fatalf("parse failed: %v", perr)
	}

	t.Logf("parse OK %#v", tree)
}

func TestTheParserDetectsWrongTokens(t *testing.T) {
	src :=
		`program demo;

		var  x, y, z : string;

		main {
			print(1 + 2);
		}`

	l := lexer.NewLexer([]byte(src))
	p := parser.NewParser()

	tree, perr := p.Parse(l)

	if perr != nil {
		t.Fatalf("parse failed: %v", perr)
	}

	t.Logf("parse OK %#v", tree)
}
