package VM

import (
	"compiler/ast"
	"compiler/memory"
	"compiler/types"
	"fmt"
)

var CallStack []Memory
var GlobalMem GlobalMemory

func RunBabyDuck(QuadrupleList []types.Quadruple) {
	// 1. Create Global memory for int and float
	GlobalMem.GlobalInts = make([]int, memory.CurrentGlobalInt)
	GlobalMem.GlobalFloats = make([]float64, memory.CurrentGlobalFloat)

	// 2. Reset Call Stack
	CallStack = make([]Memory, 0)

	var IP int = 0
	var IP_stack []int = make([]int, 0)
	var shouldAddIP bool = true

	for QuadrupleList[IP].Op != types.Terminate {
		shouldAddIP = true

		switch QuadrupleList[IP].Op {
		case types.Era:
			CallStack = append(CallStack, generateMemoryForFunction(QuadrupleList[IP].Result))
		case types.Goto:
			IP = QuadrupleList[IP].Result
			shouldAddIP = false
		case types.GoSub:
			IP_stack = append(IP_stack, IP)
			IP = QuadrupleList[IP].Result
			shouldAddIP = false
		case types.GotoF:
			if !getBoolValue(QuadrupleList[IP].Arg1) {
				IP = QuadrupleList[IP].Result
				shouldAddIP = false
			}
		case types.Assign:
			AssignValue(QuadrupleList[IP].Arg1, QuadrupleList[IP].Result)
		case types.Print:
			PrintValue(QuadrupleList[IP].Arg1)
		case types.EndFunc:
			CallStack = CallStack[:len(CallStack)-1]
			IP = IP_stack[len(IP_stack)-1]
			IP_stack = IP_stack[:len(IP_stack)-1]
			shouldAddIP = false
		case types.Terminate:
			return
		default:
			fmt.Println("Unknown operator")
		}

		if shouldAddIP {
			IP++
		}
	}
}

func generateMemoryForFunction(functionIndex int) Memory {
	function := ast.ProgramFunctions[ast.ProgramFunctionOrder[functionIndex]]

	memory := Memory{
		IntLocals:   make([]int, function.Memory.IntLocals),
		FloatLocals: make([]float64, function.Memory.FloatLocals),
		IntTemps:    make([]int, function.Memory.IntTemps),
		FloatTemps:  make([]float64, function.Memory.FloatTemps),
		BoolTemps:   make([]bool, function.Memory.BoolTemps),
	}

	return memory
}

func getBoolValue(index int) bool {
	return CallStack[len(CallStack)-1].BoolTemps[index-memory.TEMP_BOOL_START]
}

// Handles the assignment of a value to any memory location
func AssignValue(valueMemoryIndex int, resultMemoryIndex int) {
	valueType, valueMemoryType := memory.IndexToTypeAndMemoryType(valueMemoryIndex)
	_, resultMemoryType := memory.IndexToTypeAndMemoryType(resultMemoryIndex)

	if valueMemoryType == types.Temp {
		if resultMemoryType == types.Local {
			if valueType == types.Int {
				CallStack[len(CallStack)-1].IntLocals[resultMemoryIndex-memory.LOCAL_INT_START] = CallStack[len(CallStack)-1].IntTemps[valueMemoryIndex-memory.TEMP_INT_START]
			} else if valueType == types.Float {
				CallStack[len(CallStack)-1].FloatLocals[resultMemoryIndex-memory.LOCAL_FLOAT_START] = CallStack[len(CallStack)-1].FloatTemps[valueMemoryIndex-memory.TEMP_FLOAT_START]
			}
		} else if resultMemoryType == types.Global {
			if valueType == types.Int {
				GlobalMem.GlobalInts[resultMemoryIndex-memory.GLOBAL_INT_START] = CallStack[len(CallStack)-1].IntTemps[valueMemoryIndex-memory.TEMP_INT_START]
			} else if valueType == types.Float {
				GlobalMem.GlobalFloats[resultMemoryIndex-memory.GLOBAL_FLOAT_START] = CallStack[len(CallStack)-1].FloatTemps[valueMemoryIndex-memory.TEMP_FLOAT_START]
			}
		}
	} else if valueMemoryType == types.Local {
		if resultMemoryType == types.Local {
			if valueType == types.Int {
				CallStack[len(CallStack)-1].IntLocals[resultMemoryIndex-memory.LOCAL_INT_START] = CallStack[len(CallStack)-1].IntLocals[valueMemoryIndex-memory.LOCAL_INT_START]
			} else if valueType == types.Float {
				CallStack[len(CallStack)-1].FloatLocals[resultMemoryIndex-memory.LOCAL_FLOAT_START] = CallStack[len(CallStack)-1].FloatLocals[valueMemoryIndex-memory.LOCAL_FLOAT_START]
			}
		} else if resultMemoryType == types.Global {
			if valueType == types.Int {
				GlobalMem.GlobalInts[resultMemoryIndex-memory.GLOBAL_INT_START] = CallStack[len(CallStack)-1].IntLocals[valueMemoryIndex-memory.LOCAL_INT_START]
			} else if valueType == types.Float {
				GlobalMem.GlobalFloats[resultMemoryIndex-memory.GLOBAL_FLOAT_START] = CallStack[len(CallStack)-1].FloatLocals[valueMemoryIndex-memory.LOCAL_FLOAT_START]
			}
		}
	} else if valueMemoryType == types.Global {
		if resultMemoryType == types.Local {
			if valueType == types.Int {
				CallStack[len(CallStack)-1].IntLocals[resultMemoryIndex-memory.LOCAL_INT_START] = GlobalMem.GlobalInts[valueMemoryIndex-memory.GLOBAL_INT_START]
			} else if valueType == types.Float {
				CallStack[len(CallStack)-1].FloatLocals[resultMemoryIndex-memory.LOCAL_FLOAT_START] = GlobalMem.GlobalFloats[valueMemoryIndex-memory.GLOBAL_FLOAT_START]
			}
		} else if resultMemoryType == types.Global {
			if valueType == types.Int {
				GlobalMem.GlobalInts[resultMemoryIndex-memory.GLOBAL_INT_START] = GlobalMem.GlobalInts[valueMemoryIndex-memory.GLOBAL_INT_START]
			} else if valueType == types.Float {
				GlobalMem.GlobalFloats[resultMemoryIndex-memory.GLOBAL_FLOAT_START] = GlobalMem.GlobalFloats[valueMemoryIndex-memory.GLOBAL_FLOAT_START]
			}
		}
	} else if valueMemoryType == types.Constant {
		if resultMemoryType == types.Local {
			if valueType == types.Int {
				CallStack[len(CallStack)-1].IntLocals[resultMemoryIndex-memory.LOCAL_INT_START] = memory.ConstantInts[valueMemoryIndex-memory.CONSTANT_INT_START]
			} else if valueType == types.Float {
				CallStack[len(CallStack)-1].FloatLocals[resultMemoryIndex-memory.LOCAL_FLOAT_START] = memory.ConstantFloats[valueMemoryIndex-memory.CONSTANT_FLOAT_START]
			}
		} else if resultMemoryType == types.Global {
			if valueType == types.Int {
				GlobalMem.GlobalInts[resultMemoryIndex-memory.GLOBAL_INT_START] = memory.ConstantInts[valueMemoryIndex-memory.CONSTANT_INT_START]
			} else if valueType == types.Float {
				GlobalMem.GlobalFloats[resultMemoryIndex-memory.GLOBAL_FLOAT_START] = memory.ConstantFloats[valueMemoryIndex-memory.CONSTANT_FLOAT_START]
			}
		}
	}
}

func PrintValue(valueMemoryIndex int) {
	valueType, valueMemoryType := memory.IndexToTypeAndMemoryType(valueMemoryIndex)

	if valueMemoryType == types.Temp {
		if valueType == types.Int {
			fmt.Println(CallStack[len(CallStack)-1].IntTemps[valueMemoryIndex-memory.TEMP_INT_START])
		} else if valueType == types.Float {
			fmt.Println(CallStack[len(CallStack)-1].FloatTemps[valueMemoryIndex-memory.TEMP_FLOAT_START])
		} else if valueType == types.Bool {
			fmt.Println(CallStack[len(CallStack)-1].BoolTemps[valueMemoryIndex-memory.TEMP_BOOL_START])
		}
	} else if valueMemoryType == types.Local {
		if valueType == types.Int {
			fmt.Println(CallStack[len(CallStack)-1].IntLocals[valueMemoryIndex-memory.LOCAL_INT_START])
		} else if valueType == types.Float {
			fmt.Println(CallStack[len(CallStack)-1].FloatLocals[valueMemoryIndex-memory.LOCAL_FLOAT_START])
		}
	} else if valueMemoryType == types.Global {
		if valueType == types.Int {
			fmt.Println(GlobalMem.GlobalInts[valueMemoryIndex-memory.GLOBAL_INT_START])
		} else if valueType == types.Float {
			fmt.Println(GlobalMem.GlobalFloats[valueMemoryIndex-memory.GLOBAL_FLOAT_START])
		}
	} else if valueMemoryType == types.Constant {
		if valueType == types.Int {
			fmt.Println(memory.ConstantInts[valueMemoryIndex-memory.CONSTANT_INT_START])
		} else if valueType == types.Float {
			fmt.Println(memory.ConstantFloats[valueMemoryIndex-memory.CONSTANT_FLOAT_START])
		} else if valueType == types.String {
			fmt.Println(memory.ConstantStrings[valueMemoryIndex-memory.CONSTANT_STRING_START])
		}
	}
}

// Add
// Sub
// Mul
// Div

// NotEqual
// LessThan
// GreaterThan

// Parameter
