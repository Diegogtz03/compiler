package main

import (
	"compiler/ast"
	"fmt"
)

func PrintFunctionMapWithVars(functions map[string]ast.Function) {
	for _, function := range functions {
		fmt.Println("------- Function: " + function.Name + " -------")
		for _, variable := range function.Vars {
			fmt.Println("Variable: " + variable.Id + " | Type: " + variable.Type)
		}
	}
}
