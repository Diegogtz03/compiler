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

	// 3. Create Temp memory for main
	CallStack = append(CallStack, Memory{
		IntTemps:    make([]int, memory.CurrentTempInt),
		FloatTemps:  make([]float64, memory.CurrentTempFloat),
		BoolTemps:   make([]bool, memory.CurrentTempBool),
		IntLocals:   make([]int, memory.CurrentLocalInt),
		FloatLocals: make([]float64, memory.CurrentLocalFloat),
	})

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
			IP = IP_stack[len(IP_stack)-1] + 1
			IP_stack = IP_stack[:len(IP_stack)-1]
			shouldAddIP = false
		case types.Add:
			HandleOperation(types.Add, QuadrupleList[IP].Arg1, QuadrupleList[IP].Arg2, QuadrupleList[IP].Result)
		case types.Sub:
			HandleOperation(types.Sub, QuadrupleList[IP].Arg1, QuadrupleList[IP].Arg2, QuadrupleList[IP].Result)
		case types.Mul:
			HandleOperation(types.Mul, QuadrupleList[IP].Arg1, QuadrupleList[IP].Arg2, QuadrupleList[IP].Result)
		case types.Div:
			HandleOperation(types.Div, QuadrupleList[IP].Arg1, QuadrupleList[IP].Arg2, QuadrupleList[IP].Result)
		case types.NotEqual:
			HandleComparison(types.NotEqual, QuadrupleList[IP].Arg1, QuadrupleList[IP].Arg2, QuadrupleList[IP].Result)
		case types.LessThan:
			HandleComparison(types.LessThan, QuadrupleList[IP].Arg1, QuadrupleList[IP].Arg2, QuadrupleList[IP].Result)
		case types.GreaterThan:
			HandleComparison(types.GreaterThan, QuadrupleList[IP].Arg1, QuadrupleList[IP].Arg2, QuadrupleList[IP].Result)
		case types.Parameter:
			HandleParameter(QuadrupleList[IP].Arg1, QuadrupleList[IP].Result)
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
			fmt.Print(CallStack[len(CallStack)-1].IntTemps[valueMemoryIndex-memory.TEMP_INT_START], " ")
		} else if valueType == types.Float {
			fmt.Print(CallStack[len(CallStack)-1].FloatTemps[valueMemoryIndex-memory.TEMP_FLOAT_START], " ")
		} else if valueType == types.Bool {
			fmt.Print(CallStack[len(CallStack)-1].BoolTemps[valueMemoryIndex-memory.TEMP_BOOL_START], " ")
		}
	} else if valueMemoryType == types.Local {
		if valueType == types.Int {
			fmt.Print(CallStack[len(CallStack)-1].IntLocals[valueMemoryIndex-memory.LOCAL_INT_START], " ")
		} else if valueType == types.Float {
			fmt.Print(CallStack[len(CallStack)-1].FloatLocals[valueMemoryIndex-memory.LOCAL_FLOAT_START], " ")
		}
	} else if valueMemoryType == types.Global {
		if valueType == types.Int {
			fmt.Print(GlobalMem.GlobalInts[valueMemoryIndex-memory.GLOBAL_INT_START], " ")
		} else if valueType == types.Float {
			fmt.Print(GlobalMem.GlobalFloats[valueMemoryIndex-memory.GLOBAL_FLOAT_START], " ")
		}
	} else if valueMemoryType == types.Constant {
		if valueType == types.Int {
			fmt.Print(memory.ConstantInts[valueMemoryIndex-memory.CONSTANT_INT_START], " ")
		} else if valueType == types.Float {
			fmt.Print(memory.ConstantFloats[valueMemoryIndex-memory.CONSTANT_FLOAT_START], " ")
		} else if valueType == types.String {
			fmt.Print(memory.ConstantStrings[valueMemoryIndex-memory.CONSTANT_STRING_START], " ")
		}
	}
}

// All go to temps as results
func HandleOperation(operation types.Operator, value1Index int, value2Index int, resultIndex int) {
	value1Type := memory.IndexToType(value1Index)
	value2Type := memory.IndexToType(value2Index)
	resultType := memory.IndexToType(resultIndex)

	value1 := getValue(value1Index)
	value2 := getValue(value2Index)

	switch resultType {
	case types.Int:
		if operation == types.Add {
			if value1Type == types.Int && value2Type == types.Int {
				CallStack[len(CallStack)-1].IntTemps[resultIndex-memory.TEMP_INT_START] = value1.(int) + value2.(int)
			} else if value1Type == types.Float && value2Type == types.Int {
				CallStack[len(CallStack)-1].IntTemps[resultIndex-memory.TEMP_INT_START] = int(value1.(float64) + float64(value2.(int)))
			} else if value1Type == types.Int && value2Type == types.Float {
				CallStack[len(CallStack)-1].IntTemps[resultIndex-memory.TEMP_INT_START] = int(float64(value1.(int)) + value2.(float64))
			} else if value1Type == types.Float && value2Type == types.Float {
				CallStack[len(CallStack)-1].IntTemps[resultIndex-memory.TEMP_INT_START] = int(value1.(float64) + value2.(float64))
			}
		} else if operation == types.Sub {
			if value1Type == types.Int && value2Type == types.Int {
				CallStack[len(CallStack)-1].IntTemps[resultIndex-memory.TEMP_INT_START] = value1.(int) - value2.(int)
			} else if value1Type == types.Float && value2Type == types.Int {
				CallStack[len(CallStack)-1].IntTemps[resultIndex-memory.TEMP_INT_START] = int(value1.(float64) - float64(value2.(int)))
			} else if value1Type == types.Int && value2Type == types.Float {
				CallStack[len(CallStack)-1].IntTemps[resultIndex-memory.TEMP_INT_START] = int(float64(value1.(int)) - value2.(float64))
			} else if value1Type == types.Float && value2Type == types.Float {
				CallStack[len(CallStack)-1].IntTemps[resultIndex-memory.TEMP_INT_START] = int(value1.(float64) - value2.(float64))
			}
		} else if operation == types.Mul {
			if value1Type == types.Int && value2Type == types.Int {
				CallStack[len(CallStack)-1].IntTemps[resultIndex-memory.TEMP_INT_START] = value1.(int) * value2.(int)
			} else if value1Type == types.Float && value2Type == types.Int {
				CallStack[len(CallStack)-1].IntTemps[resultIndex-memory.TEMP_INT_START] = int(value1.(float64) * float64(value2.(int)))
			} else if value1Type == types.Int && value2Type == types.Float {
				CallStack[len(CallStack)-1].IntTemps[resultIndex-memory.TEMP_INT_START] = int(float64(value1.(int)) * value2.(float64))
			} else if value1Type == types.Float && value2Type == types.Float {
				CallStack[len(CallStack)-1].IntTemps[resultIndex-memory.TEMP_INT_START] = int(value1.(float64) * value2.(float64))
			}
		} else if operation == types.Div {
			if value1Type == types.Int && value2Type == types.Int {
				CallStack[len(CallStack)-1].IntTemps[resultIndex-memory.TEMP_INT_START] = value1.(int) / value2.(int)
			} else if value1Type == types.Float && value2Type == types.Int {
				CallStack[len(CallStack)-1].IntTemps[resultIndex-memory.TEMP_INT_START] = int(value1.(float64) / float64(value2.(int)))
			} else if value1Type == types.Int && value2Type == types.Float {
				CallStack[len(CallStack)-1].IntTemps[resultIndex-memory.TEMP_INT_START] = int(float64(value1.(int)) / value2.(float64))
			} else if value1Type == types.Float && value2Type == types.Float {
				CallStack[len(CallStack)-1].IntTemps[resultIndex-memory.TEMP_INT_START] = int(value1.(float64) / value2.(float64))
			}
		}
	case types.Float:
		if operation == types.Add {
			if value1Type == types.Int && value2Type == types.Int {
				CallStack[len(CallStack)-1].FloatTemps[resultIndex-memory.TEMP_FLOAT_START] = float64(value1.(int)) + float64(value2.(int))
			} else if value1Type == types.Float && value2Type == types.Int {
				CallStack[len(CallStack)-1].FloatTemps[resultIndex-memory.TEMP_FLOAT_START] = value1.(float64) + float64(value2.(int))
			} else if value1Type == types.Int && value2Type == types.Float {
				CallStack[len(CallStack)-1].FloatTemps[resultIndex-memory.TEMP_FLOAT_START] = float64(value1.(int)) + value2.(float64)
			} else if value1Type == types.Float && value2Type == types.Float {
				CallStack[len(CallStack)-1].FloatTemps[resultIndex-memory.TEMP_FLOAT_START] = value1.(float64) + value2.(float64)
			}
		} else if operation == types.Sub {
			if value1Type == types.Int && value2Type == types.Int {
				CallStack[len(CallStack)-1].FloatTemps[resultIndex-memory.TEMP_FLOAT_START] = float64(value1.(int)) - float64(value2.(int))
			} else if value1Type == types.Float && value2Type == types.Int {
				CallStack[len(CallStack)-1].FloatTemps[resultIndex-memory.TEMP_FLOAT_START] = value1.(float64) - float64(value2.(int))
			} else if value1Type == types.Int && value2Type == types.Float {
				CallStack[len(CallStack)-1].FloatTemps[resultIndex-memory.TEMP_FLOAT_START] = float64(value1.(int)) - value2.(float64)
			} else if value1Type == types.Float && value2Type == types.Float {
				CallStack[len(CallStack)-1].FloatTemps[resultIndex-memory.TEMP_FLOAT_START] = value1.(float64) - value2.(float64)
			}
		} else if operation == types.Mul {
			if value1Type == types.Int && value2Type == types.Int {
				CallStack[len(CallStack)-1].FloatTemps[resultIndex-memory.TEMP_FLOAT_START] = float64(value1.(int)) * float64(value2.(int))
			} else if value1Type == types.Float && value2Type == types.Int {
				CallStack[len(CallStack)-1].FloatTemps[resultIndex-memory.TEMP_FLOAT_START] = value1.(float64) * float64(value2.(int))
			} else if value1Type == types.Int && value2Type == types.Float {
				CallStack[len(CallStack)-1].FloatTemps[resultIndex-memory.TEMP_FLOAT_START] = float64(value1.(int)) * value2.(float64)
			} else if value1Type == types.Float && value2Type == types.Float {
				CallStack[len(CallStack)-1].FloatTemps[resultIndex-memory.TEMP_FLOAT_START] = value1.(float64) * value2.(float64)
			}
		} else if operation == types.Div {
			if value1Type == types.Int && value2Type == types.Int {
				CallStack[len(CallStack)-1].FloatTemps[resultIndex-memory.TEMP_FLOAT_START] = float64(value1.(int)) / float64(value2.(int))
			} else if value1Type == types.Float && value2Type == types.Int {
				CallStack[len(CallStack)-1].FloatTemps[resultIndex-memory.TEMP_FLOAT_START] = value1.(float64) / float64(value2.(int))
			} else if value1Type == types.Int && value2Type == types.Float {
				CallStack[len(CallStack)-1].FloatTemps[resultIndex-memory.TEMP_FLOAT_START] = float64(value1.(int)) / value2.(float64)
			} else if value1Type == types.Float && value2Type == types.Float {
				CallStack[len(CallStack)-1].FloatTemps[resultIndex-memory.TEMP_FLOAT_START] = value1.(float64) / value2.(float64)
			}
		}
	}
}

func getValue(index int) any {
	valueType, valueMemoryType := memory.IndexToTypeAndMemoryType(index)

	if valueMemoryType == types.Temp {
		if valueType == types.Int {
			return CallStack[len(CallStack)-1].IntTemps[index-memory.TEMP_INT_START]
		} else if valueType == types.Float {
			return CallStack[len(CallStack)-1].FloatTemps[index-memory.TEMP_FLOAT_START]
		} else if valueType == types.Bool {
			return CallStack[len(CallStack)-1].BoolTemps[index-memory.TEMP_BOOL_START]
		}
	} else if valueMemoryType == types.Local {
		if valueType == types.Int {
			return CallStack[len(CallStack)-1].IntLocals[index-memory.LOCAL_INT_START]
		} else if valueType == types.Float {
			return CallStack[len(CallStack)-1].FloatLocals[index-memory.LOCAL_FLOAT_START]
		}
	} else if valueMemoryType == types.Global {
		if valueType == types.Int {
			return GlobalMem.GlobalInts[index-memory.GLOBAL_INT_START]
		} else if valueType == types.Float {
			return GlobalMem.GlobalFloats[index-memory.GLOBAL_FLOAT_START]
		}
	} else if valueMemoryType == types.Constant {
		if valueType == types.Int {
			return memory.ConstantInts[index-memory.CONSTANT_INT_START]
		} else if valueType == types.Float {
			return memory.ConstantFloats[index-memory.CONSTANT_FLOAT_START]
		} else if valueType == types.String {
			return memory.ConstantStrings[index-memory.CONSTANT_STRING_START]
		}
	}

	panic("Variable not assigned")
}

func HandleComparison(operation types.Operator, value1Index int, value2Index int, resultIndex int) {
	// result is always temp bool
	value1Type := memory.IndexToType(value1Index)
	value2Type := memory.IndexToType(value2Index)

	value1 := getValue(value1Index)
	value2 := getValue(value2Index)

	if operation == types.NotEqual {
		if value1Type == types.Int && value2Type == types.Int {
			CallStack[len(CallStack)-1].BoolTemps[resultIndex-memory.TEMP_BOOL_START] = value1.(int) != value2.(int)
		} else if value1Type == types.Float && value2Type == types.Int {
			CallStack[len(CallStack)-1].BoolTemps[resultIndex-memory.TEMP_BOOL_START] = value1.(float64) != float64(value2.(int))
		} else if value1Type == types.Int && value2Type == types.Float {
			CallStack[len(CallStack)-1].BoolTemps[resultIndex-memory.TEMP_BOOL_START] = float64(value1.(int)) != value2.(float64)
		} else if value1Type == types.Float && value2Type == types.Float {
			CallStack[len(CallStack)-1].BoolTemps[resultIndex-memory.TEMP_BOOL_START] = value1.(float64) != value2.(float64)
		} else if value1Type == types.Bool && value2Type == types.Bool {
			CallStack[len(CallStack)-1].BoolTemps[resultIndex-memory.TEMP_BOOL_START] = value1.(bool) != value2.(bool)
		}
	} else if operation == types.LessThan {
		if value1Type == types.Int && value2Type == types.Int {
			CallStack[len(CallStack)-1].BoolTemps[resultIndex-memory.TEMP_BOOL_START] = value1.(int) < value2.(int)
		} else if value1Type == types.Float && value2Type == types.Int {
			CallStack[len(CallStack)-1].BoolTemps[resultIndex-memory.TEMP_BOOL_START] = value1.(float64) < float64(value2.(int))
		} else if value1Type == types.Int && value2Type == types.Float {
			CallStack[len(CallStack)-1].BoolTemps[resultIndex-memory.TEMP_BOOL_START] = float64(value1.(int)) < value2.(float64)
		} else if value1Type == types.Float && value2Type == types.Float {
			CallStack[len(CallStack)-1].BoolTemps[resultIndex-memory.TEMP_BOOL_START] = value1.(float64) < value2.(float64)
		}
	} else if operation == types.GreaterThan {
		if value1Type == types.Int && value2Type == types.Int {
			CallStack[len(CallStack)-1].BoolTemps[resultIndex-memory.TEMP_BOOL_START] = value1.(int) <= value2.(int)
		} else if value1Type == types.Float && value2Type == types.Int {
			CallStack[len(CallStack)-1].BoolTemps[resultIndex-memory.TEMP_BOOL_START] = value1.(float64) <= float64(value2.(int))
		} else if value1Type == types.Int && value2Type == types.Float {
			CallStack[len(CallStack)-1].BoolTemps[resultIndex-memory.TEMP_BOOL_START] = float64(value1.(int)) <= value2.(float64)
		} else if value1Type == types.Float && value2Type == types.Float {
			CallStack[len(CallStack)-1].BoolTemps[resultIndex-memory.TEMP_BOOL_START] = value1.(float64) <= value2.(float64)
		}
	}
}

func HandleParameter(valueMemoryIndex int, resultMemoryIndex int) {
	// Result is always local memory type
	valueType := memory.IndexToType(valueMemoryIndex)
	resultType := memory.IndexToType(resultMemoryIndex)

	if valueType != resultType {
		panic("Parameter type does not match")
	}

	value := getValue(valueMemoryIndex)

	if resultType == types.Int {
		CallStack[len(CallStack)-1].IntLocals[resultMemoryIndex-memory.LOCAL_INT_START] = value.(int)
	} else if resultType == types.Float {
		CallStack[len(CallStack)-1].FloatLocals[resultMemoryIndex-memory.LOCAL_FLOAT_START] = value.(float64)
	} else if resultType == types.Bool {
		CallStack[len(CallStack)-1].BoolTemps[resultMemoryIndex-memory.TEMP_BOOL_START] = value.(bool)
	}
}
