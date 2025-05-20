package memory

import (
	"compiler/token"
	"compiler/types"
)

type MemoryType int

const (
	Global MemoryType = iota
	Local
	Temp
	Constant
)

var ConstantIsNegative = false

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
var CurrentConstantInt = 1
var CurrentConstantFloat = 0
var CurrentConstantString = 0

// Constant memory
var ConstantInts = []int{-1}
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
		}
	case types.Bool:
		switch memoryType {
		case types.Temp:
			CurrentTempBool++
			return TEMP_BOOL_START + CurrentTempBool
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

func MakeOperandNegative() (int, error) {
	ConstantIsNegative = true
	return 0, nil
}

func ResetLocalMemory() {
	CurrentLocalInt = 0
	CurrentLocalFloat = 0
}

func AssignIntConstant(value interface{}) int {
	intValue, _ := value.(*token.Token).Int32Value()

	if ConstantIsNegative {
		intValue = -intValue
		ConstantIsNegative = false
	}

	CurrentConstantInt++
	ConstantInts = append(ConstantInts, int(intValue))

	return CONSTANT_INT_START + CurrentConstantInt
}

func AssignFloatConstant(value interface{}) int {
	floatValue, _ := value.(*token.Token).Float64Value()

	if ConstantIsNegative {
		floatValue = -floatValue
		ConstantIsNegative = false
	}

	CurrentConstantFloat++
	ConstantFloats = append(ConstantFloats, floatValue)

	return CONSTANT_FLOAT_START + CurrentConstantFloat
}

func AssignStringConstant(value *token.Token) int {
	strValue := value.CharLiteralValue()

	CurrentConstantString++
	ConstantStrings = append(ConstantStrings, strValue)

	return CONSTANT_STRING_START + CurrentConstantString
}
