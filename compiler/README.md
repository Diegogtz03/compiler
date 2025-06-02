# Go Compiler

This project contains a compiler for the "BabyDuck" programming language. It defines its grammar as a single BNF file, transformed into a parser with the library "GOCC".

## Features

- Custom BNF grammar (`grammar.bnf`)
- Lexical and syntax parsing with GOCC
- Semantic analysis using symbol tables and a semantic cube
- Intermediate code generation (quadruples)
- Memory model for global, local, temporary, and constant variables
- Virtual Machine for executing compiled code
- Support for functions, conditionals, loops, and expressions

## Project Structure

```
/compiler
│ ├── grammar.bnf
├── /ast
│   -> Contains the definitions for AST functions, AKA "neuralgic points"
├── /lexer
│   -> Contains the lexer for the defined grammar (grammar.bnf)
├── /parser
│   -> Contains the parser for the defined grammar (grammar.bnf)
├── /utils
│   -> Contains personal utils used through the project
├── /VM
│   -> Includes the files for the definiton of the VM
├── /tests
│   -> Includes the tests used to maintain the compiler alive during development
└── main.go
```

## Test Coverage

Includes tests for:

- Parser correctness and syntax error detection
- Variable/function declaration and scope errors
- Quadruple generation for expressions and control flow
- Function calling, parameter passing, and memory context
- VM execution of small programs (e.g., Fibonacci, factorial)

Note: Tests only work now by running each separetly since context gets mixed on go's testing library

## Getting Started

1. Install Go and clone this repo.
2. Run `go run main.go` to compile the BabyDuck source on main and execute the VM.
