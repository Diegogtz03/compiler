# Go Compiler

This project contains a compiler for the "BabyDuck" programming language. It defines its grammar as a single BNF file, transformed into a parser with the library "GOCC".

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
└── tsconfig.json
```

## Tests

1.

## Run Tests

```
cd compiler
go test -v
```
