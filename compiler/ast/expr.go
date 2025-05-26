package ast

import (
	"compiler/memory"
	"compiler/types"
	"compiler/utils"
	"fmt"
)

var QuadrupleList []types.Quadruple

var OperatorStack = utils.TypeStack{} // types.Operator
var OperandStack = utils.Stack{}      // int (indexes)

var operatorHierarchy = map[types.Operator]int{
	types.ErrorOperator: 0,
	types.StackDivider:  4,
	types.Assign:        4,
	types.Add:           2,
	types.Sub:           2,
	types.Mul:           1,
	types.Div:           1,
	types.NotEqual:      2,
	types.LessThan:      2,
	types.GreaterThan:   2,
	types.Print:         3,
	types.Goto:          3,
	types.GotoF:         3,
	types.GotoT:         3,

	types.Era:       4,
	types.Parameter: 3,
	types.GoSub:     4,
}

// Jump Stack
var JumpStack = utils.Stack{}

func PushOperator(op types.Operator) (types.Operator, error) {
	if OperatorStack.IsEmpty() {
		OperatorStack.Push(op)
		return op, nil
	} else if op == types.StackDivider || OperatorStack.Top() == types.StackDivider {
		OperatorStack.Push(op)
		return op, nil
	} else {
		if operatorHierarchy[OperatorStack.Top()] <= operatorHierarchy[op] {
			GenerateQuadruple(OperatorStack.Pop())
			OperatorStack.Push(op)
		} else {
			OperatorStack.Push(op)
		}
	}

	return op, nil
}

func PushOperand(index int) (int, error) {
	OperandStack.Push(index)

	return index, nil
}

func CloseFakeStack() (types.Operator, error) {
	// pop all operators from OperatorStack until getting to StackDivider
	for OperatorStack.Top() != types.StackDivider {
		GenerateQuadruple(OperatorStack.Pop())
	}

	// pop StackDivider
	OperatorStack.Pop()

	return types.StackDivider, nil
}

func GenerateQuadruple(op types.Operator) {
	// fmt.Printf("Generating quadruple for %v\n", op)
	// fmt.Printf("OperandStack: %v\n", OperandStack)
	// fmt.Printf("OperatorStack: %v\n", OperatorStack)

	op2, op1, resultType := getOperators(op)

	// fmt.Printf("op2: %v\n", op2)
	// fmt.Printf("op1: %v\n", op1)
	// fmt.Printf("resultType: %v\n", resultType)
	// fmt.Println("--------------------------------")

	var quadruple types.Quadruple

	if op == types.Assign {
		quadruple = types.Quadruple{
			Op:     op,
			Arg1:   op2,
			Arg2:   -1,
			Result: op1,
		}
	} else if op == types.Print {
		quadruple = types.Quadruple{
			Op:     op,
			Arg1:   op2,
			Arg2:   -1,
			Result: -1,
		}
	} else if op == types.Goto {
		quadruple = types.Quadruple{
			Op:     op,
			Arg1:   -1,
			Arg2:   -1,
			Result: -1,
		}

		JumpStack.Push(len(QuadrupleList))
	} else if op == types.GotoF {
		// Check if op2 is a bool
		if memory.IndexToType(op2) != types.Bool {
			panic("Error: argument must be a bool")
		}

		quadruple = types.Quadruple{
			Op:     op,
			Arg1:   op2,
			Arg2:   -1,
			Result: -1,
		}

		JumpStack.Push(len(QuadrupleList))
	} else {
		if resultType == types.Error {
			panic("Error: Invalid operation")
		}

		// memoryType := types.MemoryType(memory.Global)

		// if GlobalProgramName != CurrentModule {
		// 	memoryType = types.MemoryType(memory.Local)
		// } else if resultType == types.Bool {
		// 	memoryType = types.MemoryType(memory.Temp)
		// }

		memoryType := types.MemoryType(memory.Temp)

		var tempIndex int = memory.AllocateMemory(resultType, memoryType)

		quadruple = types.Quadruple{
			Op:     op,
			Arg1:   op1,
			Arg2:   op2,
			Result: tempIndex,
		}
		// push tempIndex to OperandStack
		OperandStack.Push(tempIndex)
	}

	QuadrupleList = append(QuadrupleList, quadruple)
}

// Used for while statement
func PushJumpStack() (int, error) {
	JumpStack.Push(len(QuadrupleList))
	return len(QuadrupleList), nil
}

func getOperators(op types.Operator) (int, int, types.Type) {
	if op == types.Print {
		return OperandStack.Pop(), -1, types.Error
	} else if op == types.Goto {
		return -1, -1, types.Error
	} else if op == types.GotoF {
		return OperandStack.Pop(), -1, types.Error
	}

	op1 := OperandStack.Pop()
	op2 := OperandStack.Pop()

	// get types from op1 and op2
	op1Type := memory.IndexToType(op1)
	op2Type := memory.IndexToType(op2)

	// check semantics with semantic cube
	var resultType types.Type = syntaxCube[op1Type][op2Type][op]

	if resultType == types.Error {
		panic("Error: Invalid operation")
	}

	return op1, op2, resultType
}

func EndExpression() (*types.Operator, error) {
	// Pop all of the remaining operators from the stack
	for !OperatorStack.IsEmpty() {
		GenerateQuadruple(OperatorStack.Pop())
	}

	return nil, nil
}

func PopJumpStack(shouldSkipLine bool) (int, error) {
	quadrupleIndex := JumpStack.Pop()
	fmt.Printf("quadrupleIndex: %v\n", quadrupleIndex)
	fmt.Printf("len(QuadrupleList): %v\n", len(QuadrupleList))

	quadruple := QuadrupleList[quadrupleIndex]

	if quadruple.Op == types.GotoF {
		if shouldSkipLine {
			quadruple.Result = len(QuadrupleList) + 1
		} else {
			quadruple.Result = len(QuadrupleList)
		}
	} else if quadruple.Op == types.Goto {
		quadruple.Result = len(QuadrupleList)
	}

	QuadrupleList[quadrupleIndex] = quadruple

	return 0, nil
}

func CyclePopJumpStack() (int, error) {
	quadrupleIndex := JumpStack.Pop()
	quadruple := QuadrupleList[quadrupleIndex]
	quadruple.Result = len(QuadrupleList) + 1
	QuadrupleList[quadrupleIndex] = quadruple

	fmt.Printf("QuadrupleList: %v\n", QuadrupleList)

	// Generate quadruple for Goto at end of cycle
	quadruple = types.Quadruple{
		Op:     types.Goto,
		Arg1:   -1,
		Arg2:   -1,
		Result: JumpStack.Pop(),
	}

	QuadrupleList = append(QuadrupleList, quadruple)

	return 0, nil
}

func MainGoTo() (int, error) {
	quadruple := types.Quadruple{
		Op:     types.Goto,
		Arg1:   -1,
		Arg2:   -1,
		Result: -1,
	}

	JumpStack.Push(len(QuadrupleList))

	QuadrupleList = append(QuadrupleList, quadruple)

	return 0, nil
}

func FillMainGoTo() (int, error) {
	quadrupleIndex := JumpStack.Pop()
	quadruple := QuadrupleList[quadrupleIndex]
	quadruple.Result = len(QuadrupleList)
	QuadrupleList[quadrupleIndex] = quadruple

	return 0, nil
}

func GenerateTerminate() (int, error) {
	quadruple := types.Quadruple{
		Op:     types.Terminate,
		Arg1:   -1,
		Arg2:   -1,
		Result: -1,
	}

	QuadrupleList = append(QuadrupleList, quadruple)

	return 0, nil
}

func EndFunction() (int, error) {
	quadruple := types.Quadruple{
		Op:     types.EndFunc,
		Arg1:   -1,
		Arg2:   -1,
		Result: -1,
	}

	QuadrupleList = append(QuadrupleList, quadruple)

	return 0, nil
}
