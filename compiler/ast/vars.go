package ast

import (
	"compiler/memory"
	"compiler/token"
	"compiler/types"
	"compiler/utils"
	"errors"
	"fmt"
)

// Struct for variables that will be used to store variables for functions
type Variable struct {
	Id          string
	Type        types.Type
	MemoryIndex int
}

// Queue for variables, used to store upcoming ids which type is not yet defined
var varsQueue = utils.Queue{}

// Variable that keeps track of the current type of the variable to be used
var CurrentType types.Type = types.Error

// Tells us if we need to insert this as a parameter in the function signature
var IsParameter bool = false

// Syntax cube to validate expression types and correct type assignments
var syntaxCube = map[types.Type]map[types.Type]map[types.Operator]types.Type{
	types.Int: {
		types.Int: {
			types.Add:         types.Int,
			types.Sub:         types.Int,
			types.Mul:         types.Int,
			types.Div:         types.Float,
			types.GreaterThan: types.Bool,
			types.LessThan:    types.Bool,
			types.NotEqual:    types.Bool,
		},
		types.Float: {
			types.Add:         types.Float,
			types.Sub:         types.Float,
			types.Mul:         types.Float,
			types.Div:         types.Float,
			types.GreaterThan: types.Bool,
			types.LessThan:    types.Bool,
			types.NotEqual:    types.Bool,
		},
		types.Bool: {
			types.Add:         types.Error,
			types.Sub:         types.Error,
			types.Mul:         types.Error,
			types.Div:         types.Error,
			types.GreaterThan: types.Error,
			types.LessThan:    types.Error,
			types.NotEqual:    types.Error,
		},
	},
	types.Float: {
		types.Int: {
			types.Add:         types.Float,
			types.Sub:         types.Float,
			types.Mul:         types.Float,
			types.Div:         types.Float,
			types.GreaterThan: types.Bool,
			types.LessThan:    types.Bool,
			types.NotEqual:    types.Bool,
		},
		types.Float: {
			types.Add:         types.Float,
			types.Sub:         types.Float,
			types.Mul:         types.Float,
			types.Div:         types.Float,
			types.GreaterThan: types.Bool,
			types.LessThan:    types.Bool,
			types.NotEqual:    types.Bool,
		},
		types.Bool: {
			types.Add:         types.Error,
			types.Sub:         types.Error,
			types.Mul:         types.Error,
			types.Div:         types.Error,
			types.GreaterThan: types.Error,
			types.LessThan:    types.Error,
			types.NotEqual:    types.Error,
		},
	},
	types.Bool: {
		types.Int: {
			types.Add:         types.Error,
			types.Sub:         types.Error,
			types.Mul:         types.Error,
			types.Div:         types.Error,
			types.GreaterThan: types.Error,
			types.LessThan:    types.Error,
			types.NotEqual:    types.Error,
		},
		types.Float: {
			types.Add:         types.Error,
			types.Sub:         types.Error,
			types.Mul:         types.Error,
			types.Div:         types.Error,
			types.GreaterThan: types.Error,
			types.LessThan:    types.Error,
			types.NotEqual:    types.Error,
		},
		types.Bool: {
			types.Add:         types.Error,
			types.Sub:         types.Error,
			types.Mul:         types.Error,
			types.Div:         types.Error,
			types.GreaterThan: types.Error,
			types.LessThan:    types.Error,
			types.NotEqual:    types.Error,
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

func SetCurrentType(varType types.Type) (types.Type, error) {
	CurrentType = varType

	return varType, nil
}

func InsertingParameter() (interface{}, error) {
	IsParameter = true

	return nil, nil
}

func AddVarsToTable(varType types.Type) (*Variable, error) {
	var memoryType types.MemoryType

	if CurrentModule == GlobalProgramName {
		memoryType = types.MemoryType(memory.Global)
	} else {
		memoryType = types.MemoryType(memory.Local)
	}

	for !varsQueue.IsEmpty() {
		id := varsQueue.Dequeue()

		// Check if the variable already exists in the current module only and also in the global scope
		if _, ok := ProgramFunctions[CurrentModule].Vars[id]; ok {
			return nil, fmt.Errorf("Variable %s already exists in module %s", id, CurrentModule)
		}
		// else if CurrentModule != GlobalProgramName {
		// 	if _, ok := ProgramFunctions[GlobalProgramName].Vars[id]; ok {
		// 		return nil, fmt.Errorf("Variable %s already exists in module %s", id, CurrentModule)
		// 	}
		// }

		// Reserve space in "memory" for the variable
		var index int = memory.AllocateMemory(varType, memoryType)

		var newVar = Variable{
			Id:          id,
			Type:        varType,
			MemoryIndex: index,
		}

		ProgramFunctions[CurrentModule].Vars[id] = newVar

		if IsParameter {
			currentFunction := ProgramFunctions[CurrentModule]
			currentFunction.Params.Order = append(currentFunction.Params.Order, varType)
			currentFunction.Params.Signature = append(currentFunction.Params.Signature, index)
			ProgramFunctions[CurrentModule] = currentFunction

			IsParameter = false
		}
	}

	return nil, nil
}

func GetVarIndex(stmt interface{}) (int, error) {
	id := string(stmt.(*token.Token).Lit)

	if _, ok := ProgramFunctions[CurrentModule].Vars[id]; ok {
		var index int = ProgramFunctions[CurrentModule].Vars[id].MemoryIndex
		PushOperand(index)

		// Push negative quadruple multiplication
		if memory.ConstantIsNegative {
			PushOperand(memory.CONSTANT_INT_START)
			PushOperator(types.Mul)
		}

		return index, nil
	} else if CurrentModule != GlobalProgramName {
		if _, ok := ProgramFunctions[GlobalProgramName].Vars[id]; ok {
			var index int = ProgramFunctions[GlobalProgramName].Vars[id].MemoryIndex
			PushOperand(index)

			// Push negative quadruple multiplication
			if memory.ConstantIsNegative {
				PushOperand(memory.CONSTANT_INT_START)
				PushOperator(types.Mul)
			}

			return index, nil
		}
	}
	return 0, fmt.Errorf("Variable %s not found in module %s", id, CurrentModule)
}
