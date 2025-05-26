package ast

import (
	"compiler/memory"
	"compiler/token"
	"compiler/types"
	"fmt"
)

// Struct for functions that will be used to store functions for the program
type Function struct {
	Name           string
	Vars           map[string]Variable
	Params         Params
	Memory         FunctionMemory
	QuadrupleIndex int
	FunctionIndex  int
}

type Params struct {
	Signature []int
	Order     []types.Type
}

type FunctionMemory struct {
	IntLocals   int
	FloatLocals int
	IntTemps    int
	FloatTemps  int
	BoolTemps   int
}

// Variable that stores all of the functions for the program, each function has its own variable dictionary
var ProgramFunctions map[string]Function = make(map[string]Function)

var ProgramFunctionOrder []string = make([]string, 0)

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
		Name:           id,
		Vars:           make(map[string]Variable),
		QuadrupleIndex: len(QuadrupleList),
		FunctionIndex:  len(ProgramFunctionOrder),
	}

	if isProgram {
		GlobalProgramName = id
	}

	ProgramFunctions[id] = newFunc
	CurrentModule = id
	ProgramFunctionOrder = append(ProgramFunctionOrder, id)

	memory.ResetLocalMemory()

	return &newFunc, nil
}

func ResetToGlobalScope() (string, error) {
	// Count the number of locals and temps in the current module (seeing the counter at the end of the function)
	newMemory := FunctionMemory{
		IntLocals:   memory.CurrentLocalInt,
		FloatLocals: memory.CurrentLocalFloat,
		IntTemps:    memory.CurrentTempInt,
		FloatTemps:  memory.CurrentTempFloat,
		BoolTemps:   memory.CurrentTempBool,
	}

	oldFunction := ProgramFunctions[CurrentModule]
	oldFunction.Memory = newMemory
	ProgramFunctions[CurrentModule] = oldFunction

	CurrentModule = GlobalProgramName
	memory.ResetLocalMemory()

	return CurrentModule, nil
}

// FUNCTION CALLING
var CurrentFunctionCallName string = ""
var CurrentParamIndex int = 0

func FunctionCallCreate(name interface{}) (int, error) {
	functionName := string(name.(*token.Token).Lit)

	CurrentFunctionCallName = functionName

	if _, ok := ProgramFunctions[functionName]; !ok {
		return 0, fmt.Errorf("Function %s not found", functionName)
	}

	function := ProgramFunctions[functionName]

	quadruple := types.Quadruple{
		Op:     types.Era,
		Arg1:   -1,
		Arg2:   -1,
		Result: function.FunctionIndex,
	}

	QuadrupleList = append(QuadrupleList, quadruple)

	return 0, nil
}

func FunctionCallFill() (int, error) {
	// End Expression
	EndExpression()

	// Check types of needed and given param
	valueToInsert := OperandStack.Pop()

	memoryType := memory.IndexToType(valueToInsert)

	if memoryType != ProgramFunctions[CurrentFunctionCallName].Params.Order[CurrentParamIndex] {
		return 0, fmt.Errorf("type mismatch for parameter %d in function %s", CurrentParamIndex, CurrentFunctionCallName)
	}

	quadruple := types.Quadruple{
		Op:     types.Parameter,
		Arg1:   valueToInsert,
		Arg2:   -1,
		Result: ProgramFunctions[CurrentFunctionCallName].Params.Signature[CurrentParamIndex],
	}

	QuadrupleList = append(QuadrupleList, quadruple)

	CurrentParamIndex++

	return 0, nil
}

func VerifyParamFill() (int, error) {
	// Size should match
	if CurrentParamIndex < len(ProgramFunctions[CurrentFunctionCallName].Params.Order) || CurrentParamIndex > len(ProgramFunctions[CurrentFunctionCallName].Params.Order) {
		return 0, fmt.Errorf("too many parameters for function %s", CurrentFunctionCallName)
	}

	// Generate GoSub quadruple
	quadruple := types.Quadruple{
		Op:     types.GoSub,
		Arg1:   -1,
		Arg2:   -1,
		Result: ProgramFunctions[CurrentFunctionCallName].QuadrupleIndex,
	}

	QuadrupleList = append(QuadrupleList, quadruple)
	return 0, nil
}
