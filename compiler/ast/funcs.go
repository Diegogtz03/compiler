package ast

import (
	"compiler/memory"
	"compiler/token"
	"fmt"
)

// Struct for functions that will be used to store functions for the program
type Function struct {
	Name string
	Vars map[string]Variable
}

// Variable that stores all of the functions for the program, each function has its own variable dictionary
var ProgramFunctions map[string]Function = make(map[string]Function)

// Variable that will be used to store the name of the program for Global Scope lookup (avoid having another separate table for global variables/functions)
var GlobalProgramName string = ""

// Variable that will be used to store the current module the compiler is at
var CurrentModule string = ""

func CreateFuntion(stmt interface{}, isProgram bool) (*Function, error) {
	id := string(stmt.(*token.Token).Lit)

	// Check if the function already exists in the defined functions table
	if _, ok := ProgramFunctions[id]; ok {
		return nil, fmt.Errorf("Function %s already exists", id)
	}

	var newFunc = Function{
		Name: id,
		Vars: make(map[string]Variable),
	}

	if isProgram {
		GlobalProgramName = id
	}

	ProgramFunctions[id] = newFunc
	CurrentModule = id

	return &newFunc, nil
}

func ResetToGlobalScope() (string, error) {
	CurrentModule = GlobalProgramName
	memory.ResetLocalMemory()
	return CurrentModule, nil
}
