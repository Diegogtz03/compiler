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
		`program fibonacci;
        var n, a, b, count, temp : int; z, e : float;

        void fib(n : int) [
            var a, b, i, temp : int;
            {
                a = 0;
                b = 1;
                i = 0;

                while (i < n) do {
                    temp = a + b;
                    a = b;
                    b = temp;
                    i = i + 1;
                };

                print("Fibonacci of", n, "is", a);
            }
        ];

        void fact(n : int) [
            var result, counter : int;

            {
                result = 1;
                counter = 2;
                while (counter < n + 1) do {
                    result = result * counter;
                    counter = counter + 1;
                };
                print("Factorial of", n, "is", result);
            }
        ];

        main {
            a = 10;
            z = 40.34;
            e = 22.345;
            fib(12);
            fact(5);

            print(z - e, (z - e) * 2);
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
