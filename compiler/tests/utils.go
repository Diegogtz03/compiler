package main

import (
	"compiler/ast"
	"compiler/types"
	"fmt"
	"strconv"
)

func PrintFunctionMapWithVars(functions map[string]ast.Function) {
	for _, function := range functions {
		fmt.Println("------- Function: " + function.Name + " -------")
		for _, variable := range function.Vars {
			fmt.Println("Variable: " + variable.Id + " | Type: " + strconv.Itoa(int(variable.Type)) + " | Memory Index: " + strconv.Itoa(variable.MemoryIndex))
		}
	}
}

func PrintQuadrupleList(quadrupleList []types.Quadruple) {
	for _, quadruple := range quadrupleList {
		opString := ""
		switch quadruple.Op {
		case types.Add:
			opString = "Add"
		case types.Sub:
			opString = "Sub"
		case types.Mul:
			opString = "Mul"
		case types.Div:
			opString = "Div"
		case types.Assign:
			opString = "Assign"
		case types.NotEqual:
			opString = "NotEqual"
		case types.LessThan:
			opString = "LessThan"
		case types.GreaterThan:
			opString = "GreaterThan"
		case types.Print:
			opString = "Print"
		case types.StackDivider:
			opString = "StackDivider"
		case types.ErrorOperator:
			opString = "ErrorOperator"
		default:
			opString = "Unknown"
		}
		fmt.Println(opString, quadruple.Arg1, quadruple.Arg2, quadruple.Result)
	}
}
