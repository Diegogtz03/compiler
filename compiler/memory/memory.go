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
			index := GLOBAL_INT_START + CurrentGlobalInt
			CurrentGlobalInt++
			return index
		case types.Local:
			index := LOCAL_INT_START + CurrentLocalInt
			CurrentLocalInt++
			return index
		case types.Temp:
			index := TEMP_INT_START + CurrentTempInt
			CurrentTempInt++
			return index
		}
	case types.Float:
		switch memoryType {
		case types.Global:
			index := GLOBAL_FLOAT_START + CurrentGlobalFloat
			CurrentGlobalFloat++
			return index
		case types.Local:
			index := LOCAL_FLOAT_START + CurrentLocalFloat
			CurrentLocalFloat++
			return index
		case types.Temp:
			index := TEMP_FLOAT_START + CurrentTempFloat
			CurrentTempFloat++
			return index
		}
	case types.Bool:
		switch memoryType {
		case types.Temp:
			index := TEMP_BOOL_START + CurrentTempBool
			CurrentTempBool++
			return index
		}
	}
	return 0
}

func IndexToType(index int) types.Type {
	if index >= GLOBAL_INT_START && index < GLOBAL_FLOAT_START {
		return types.Int
	}
	if index >= GLOBAL_FLOAT_START && index < LOCAL_INT_START {
		return types.Float
	}
	if index >= LOCAL_INT_START && index < LOCAL_FLOAT_START {
		return types.Int
	}
	if index >= LOCAL_FLOAT_START && index < TEMP_INT_START {
		return types.Float
	}
	if index >= TEMP_INT_START && index < TEMP_FLOAT_START {
		return types.Int
	}
	if index >= TEMP_FLOAT_START && index < TEMP_BOOL_START {
		return types.Float
	}
	if index >= TEMP_BOOL_START && index < CONSTANT_INT_START {
		return types.Bool
	}
	if index >= CONSTANT_INT_START && index < CONSTANT_FLOAT_START {
		return types.Int
	}
	if index >= CONSTANT_FLOAT_START && index < CONSTANT_STRING_START {
		return types.Float
	}
	if index >= CONSTANT_STRING_START {
		return types.String
	}
	return types.Error
}

func IndexToTypeAndMemoryType(index int) (types.Type, types.MemoryType) {
	if index >= GLOBAL_INT_START && index < GLOBAL_FLOAT_START {
		return types.Int, types.Global
	}
	if index >= GLOBAL_FLOAT_START && index < LOCAL_INT_START {
		return types.Float, types.Global
	}
	if index >= LOCAL_INT_START && index < LOCAL_FLOAT_START {
		return types.Int, types.Local
	}
	if index >= LOCAL_FLOAT_START && index < TEMP_INT_START {
		return types.Float, types.Local
	}
	if index >= TEMP_INT_START && index < TEMP_FLOAT_START {
		return types.Int, types.Temp
	}
	if index >= TEMP_FLOAT_START && index < TEMP_BOOL_START {
		return types.Float, types.Temp
	}
	if index >= TEMP_BOOL_START && index < CONSTANT_INT_START {
		return types.Bool, types.Temp
	}
	if index >= CONSTANT_INT_START && index < CONSTANT_FLOAT_START {
		return types.Int, types.Constant
	}
	if index >= CONSTANT_FLOAT_START && index < CONSTANT_STRING_START {
		return types.Float, types.Constant
	}
	if index >= CONSTANT_STRING_START {
		return types.String, types.Constant
	}
	return types.Error, types.Global
}

func MakeOperandNegative() (int, error) {
	ConstantIsNegative = true
	return 0, nil
}

func ResetLocalMemory() {
	CurrentLocalInt = 0
	CurrentLocalFloat = 0
	CurrentTempInt = 0
	CurrentTempFloat = 0
	CurrentTempBool = 0
}

func AssignIntConstant(value interface{}) int {
	intValue, _ := value.(*token.Token).Int32Value()

	if ConstantIsNegative {
		intValue = -intValue
		ConstantIsNegative = false
	}

	index := CONSTANT_INT_START + CurrentConstantInt
	CurrentConstantInt++
	ConstantInts = append(ConstantInts, int(intValue))
	return index
}

func AssignFloatConstant(value interface{}) int {
	floatValue, _ := value.(*token.Token).Float64Value()

	if ConstantIsNegative {
		floatValue = -floatValue
		ConstantIsNegative = false
	}

	index := CONSTANT_FLOAT_START + CurrentConstantFloat
	CurrentConstantFloat++
	ConstantFloats = append(ConstantFloats, floatValue)
	return index
}

func AssignStringConstant(value *token.Token) int {
	strValue := value.CharLiteralValue()

	index := CONSTANT_STRING_START + CurrentConstantString
	CurrentConstantString++
	ConstantStrings = append(ConstantStrings, strValue)
	return index
}
