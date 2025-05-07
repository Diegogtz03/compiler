package ast

import (
	"compiler/token"
	"compiler/utils"
	"errors"
	"fmt"
)

// Struct for variables that will be used to store variables for functions
type Variable struct {
	Id    string
	Type  string
	Value string
}

// Queue for variables, used to store upcoming ids which type is not yet defined
var varsQueue = utils.Queue{}

// Variable that keeps track of the current type of the variable to be used
var CurrentType string = ""

type Type int

const (
	Int Type = iota
	Float
	Error
)

// Syntax cube to validate expression types and correct type assignments
var syntaxCube = map[Type]map[Type]map[string]Type{
	Int: {
		Int: {
			"+":  Int,
			"-":  Int,
			"*":  Int,
			"/":  Int,
			">":  Int,
			"<":  Int,
			"!=": Int,
		},
		Float: {
			"+":  Float,
			"-":  Float,
			"*":  Float,
			"/":  Float,
			">":  Int,
			"<":  Int,
			"!=": Int,
		},
	},
	Float: {
		Int: {
			"+":  Float,
			"-":  Float,
			"*":  Float,
			"/":  Float,
			">":  Int,
			"<":  Int,
			"!=": Int,
		},
		Float: {
			"+":  Float,
			"-":  Float,
			"*":  Float,
			"/":  Float,
			">":  Int,
			"<":  Int,
			"!=": Int,
		},
	},
}

func AddVarToQueue(name interface{}) (*Variable, error) {
	id := string(name.(*token.Token).Lit)

	varsQueue.Enqueue(id)

	// Check if the variable already exists in the current module only
	if _, ok := ProgramFunctions[CurrentModule].Vars[id]; ok {
		return nil, errors.New("Variable " + id + " already exists in module " + CurrentModule)
	}

	return nil, nil
}

func SetCurrentType(varType string) (string, error) {
	CurrentType = varType

	return varType, nil
}

func AddVarsToTable(varType string) (*Variable, error) {
	for !varsQueue.IsEmpty() {
		id := varsQueue.Dequeue()

		// Check if the variable already exists in the current module only and also in the global scope
		if _, ok := ProgramFunctions[CurrentModule].Vars[id]; ok {
			return nil, fmt.Errorf("Variable " + id + " already exists in module " + CurrentModule)
		} else if CurrentModule != GlobalProgramName {
			if _, ok := ProgramFunctions[GlobalProgramName].Vars[id]; ok {
				return nil, fmt.Errorf("Global variable " + id + " already exists in program " + GlobalProgramName)
			}
		}

		var newVar = Variable{
			Id:   id,
			Type: varType,
		}

		ProgramFunctions[CurrentModule].Vars[id] = newVar
	}

	return nil, nil
}

func GetVarValue(name interface{}) (interface{}, error) {
	id := string(name.(*token.Token).Lit)

	if v, ok := ProgramFunctions[CurrentModule].Vars[id]; ok {
		if v.Value != "assigned" {
			return nil, fmt.Errorf("Variable '" + id + "' not assigned")
		}
		return v.Value, nil
	} else if CurrentModule != GlobalProgramName {
		if v, ok := ProgramFunctions[GlobalProgramName].Vars[id]; ok {
			if v.Value != "assigned" {
				return nil, fmt.Errorf("Variable '" + id + "' not assigned")
			}
			return v.Value, nil
		}
	}

	return nil, fmt.Errorf("Variable '" + id + "' not defined")
}

// TODO: Implement correct assignment of values to variables
func AssignVarValue(name interface{}) (interface{}, error) {
	id := string(name.(*token.Token).Lit)

	if v, ok := ProgramFunctions[CurrentModule].Vars[id]; ok {
		v.Value = "assigned"
		ProgramFunctions[CurrentModule].Vars[id] = v

		return nil, nil
	} else if CurrentModule != GlobalProgramName {
		if v, ok := ProgramFunctions[GlobalProgramName].Vars[id]; ok {
			v.Value = "assigned"
			ProgramFunctions[GlobalProgramName].Vars[id] = v

			return nil, nil
		}
	}

	return nil, fmt.Errorf("Variable '" + id + "' not defined")
}
