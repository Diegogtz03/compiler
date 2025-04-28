package main

import (
	"fmt"

	"github.com/alecthomas/participle/v2"
)

func main() {
	parser, err := participle.Build[BabyDuck]()

	if err != nil {
		panic(err)
	}

	input := `
		program babyDuck;

		main

		end
	`

	expr, err := parser.ParseString("", input)

	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", expr)
}
