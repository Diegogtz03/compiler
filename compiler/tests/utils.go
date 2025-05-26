package main

import (
	"compiler/ast"
	"compiler/types"
	"fmt"
	"strconv"
)

// ------------------------ Debugging Functions ------------------------

func PrintFunctionMapWithVars(functions map[string]ast.Function) {
	for _, function := range functions {
		fmt.Println("------- Function: " + function.Name + " -------")
		fmt.Println("Memory: " + strconv.Itoa(int(function.Memory.IntLocals)) + " | " + strconv.Itoa(int(function.Memory.FloatLocals)) + " | " + strconv.Itoa(int(function.Memory.IntTemps)) + " | " + strconv.Itoa(int(function.Memory.FloatTemps)) + " | " + strconv.Itoa(int(function.Memory.BoolTemps)))

		for i, param := range function.Params.Order {
			fmt.Println("Param: " + strconv.Itoa(int(param)) + " | Signature: " + strconv.Itoa(function.Params.Signature[i]))
		}

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
		case types.GotoF:
			opString = "GotoF"
		case types.Goto:
			opString = "Goto"
		case types.GotoT:
			opString = "GotoT"
		case types.GoSub:
			opString = "GoSub"
		case types.Era:
			opString = "Era"
		case types.Parameter:
			opString = "Parameter"
		case types.EndFunc:
			opString = "EndFunc"
		case types.Terminate:
			opString = "Terminate"
		default:
			opString = "Unknown"
		}
		fmt.Println(opString, quadruple.Arg1, quadruple.Arg2, quadruple.Result)
	}
}

// ------------------------ Assert Functions ------------------------

func AssertFunctionAndVarsMap(functionMap map[string]ast.Function, expectedFunctionMap map[string]ast.Function) bool {
	if len(functionMap) != len(expectedFunctionMap) {
		fmt.Printf("Function map length mismatch: %d != %d\n", len(functionMap), len(expectedFunctionMap))
		return false
	}

	for _, function := range functionMap {
		if _, ok := expectedFunctionMap[function.Name]; !ok {
			fmt.Printf("Function %s not found in expected function map\n", function.Name)
			return false
		}
	}

	for _, function := range functionMap {
		for _, variable := range function.Vars {
			if _, ok := expectedFunctionMap[function.Name].Vars[variable.Id]; !ok {
				fmt.Printf("Variable %s not found in expected function map\n", variable.Id)
				return false
			}
		}
	}

	return true
}

func AssertQuadrupleList(quadrupleList []types.Quadruple, expectedQuadrupleList []types.Quadruple) bool {
	if len(quadrupleList) != len(expectedQuadrupleList) {
		fmt.Printf("Quadruple list length mismatch: %d != %d\n", len(quadrupleList), len(expectedQuadrupleList))
		return false
	}

	for i, quadruple := range quadrupleList {
		if quadruple != expectedQuadrupleList[i] {
			fmt.Printf("Quadruple mismatch at index %d: %v != %v\n", i, quadruple, expectedQuadrupleList[i])

			return false
		}
	}

	return true
}
