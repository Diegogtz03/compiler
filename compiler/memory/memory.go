package memory

import (
	"compiler/types"
)

type MemoryType int

const (
	Global MemoryType = iota
	Local
	Temp
	Constant
)

var GLOBAL_INT_START = 1000
var GLOBAL_FLOAT_START = 1200

var LOCAL_INT_START = 2000
var LOCAL_FLOAT_START = 2200

var TEMP_INT_START = 3000
var TEMP_FLOAT_START = 3200
var TEMP_BOOL_START = 3400

var CONSTANT_INT_START = 4200
var CONSTANT_FLOAT_START = 4400
var CONSTANT_STRING_START = 4600

// Current memory index
var CurrentGlobalInt = 0
var CurrentGlobalFloat = 0
var CurrentLocalInt = 0
var CurrentLocalFloat = 0
var CurrentTempInt = 0
var CurrentTempFloat = 0
var CurrentTempBool = 0
var CurrentConstantInt = 0
var CurrentConstantFloat = 0
var CurrentConstantString = 0

// Global memory
var GlobalInts = []int{}
var GlobalFloats = []float64{}

var LocalInts = []int{}
var LocalFloats = []float64{}

var TempInts = []int{}
var TempFloats = []float64{}
var TempBools = []bool{}

var ConstantInts = []int{}
var ConstantFloats = []float64{}
var ConstantStrings = []string{}

func AllocateMemory(varType types.Type, memoryType types.MemoryType) int {
	switch varType {
	case types.Int:
		switch memoryType {
		case types.Global:
			CurrentGlobalInt++
			return GLOBAL_INT_START + CurrentGlobalInt
		case types.Local:
			CurrentLocalInt++
			return LOCAL_INT_START + CurrentLocalInt
		case types.Temp:
			CurrentTempInt++
			return TEMP_INT_START + CurrentTempInt
		case types.Constant:
			CurrentConstantInt++
			return CONSTANT_INT_START + CurrentConstantInt
		}
	case types.Float:
		switch memoryType {
		case types.Global:
			CurrentGlobalFloat++
			return GLOBAL_FLOAT_START + CurrentGlobalFloat
		case types.Local:
			CurrentLocalFloat++
			return LOCAL_FLOAT_START + CurrentLocalFloat
		case types.Temp:
			CurrentTempFloat++
			return TEMP_FLOAT_START + CurrentTempFloat
		case types.Constant:
			CurrentConstantFloat++
			return CONSTANT_FLOAT_START + CurrentConstantFloat
		}
	case types.Bool:
		switch memoryType {
		case types.Temp:
			CurrentTempBool++
			return TEMP_BOOL_START + CurrentTempBool
		}
	case types.String:
		switch memoryType {
		case types.Constant:
			CurrentConstantString++
			return CONSTANT_STRING_START + CurrentConstantString
		}
	}
	return 0
}

func IndexToType(index int) types.Type {
	if index >= GLOBAL_INT_START && index < GLOBAL_FLOAT_START {
		return types.Int
	}
	if index >= GLOBAL_FLOAT_START && index < GLOBAL_FLOAT_START {
		return types.Float
	}
	if index >= LOCAL_INT_START && index < LOCAL_FLOAT_START {
		return types.Int
	}
	if index >= LOCAL_FLOAT_START && index < LOCAL_FLOAT_START {
		return types.Float
	}
	if index >= TEMP_INT_START && index < TEMP_FLOAT_START {
		return types.Int
	}
	if index >= TEMP_FLOAT_START && index < TEMP_FLOAT_START {
		return types.Float
	}
	if index >= TEMP_BOOL_START && index < TEMP_BOOL_START {
		return types.Bool
	}
	if index >= CONSTANT_INT_START && index < CONSTANT_FLOAT_START {
		return types.Int
	}
	if index >= CONSTANT_FLOAT_START && index < CONSTANT_STRING_START {
		return types.Float
	}
	if index >= CONSTANT_STRING_START && index < CONSTANT_STRING_START {
		return types.String
	}
	return types.Error
}

func ResetLocalMemory() {
	CurrentLocalInt = 0
	CurrentLocalFloat = 0
}
