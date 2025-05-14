package ast

import (
	"compiler/memory"
	"compiler/types"
	"compiler/utils"
)

var QuadrupleList []types.Quadruple

var OperatorStack = utils.TypeStack{} // types.Operator
var OperandStack = utils.Stack{}      // int (indexes)

var operatorHierarchy = map[types.Operator]int{
	types.ErrorOperator: 0,
	types.StackDivider:  10,
	types.Assign:        10,
	types.Add:           3,
	types.Sub:           3,
	types.Mul:           2,
	types.Div:           2,
	types.NotEqual:      1,
	types.LessThan:      1,
	types.GreaterThan:   1,
	types.Print:         4,
}

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

func PushOperand(index int) {
	OperandStack.Push(index)
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

// Todo: adapt for assignment, etc...
func GenerateQuadruple(op types.Operator) {
	op2, op1 := OperandStack.Pop(), OperandStack.Pop()

	// get types from op1 and op2
	op1Type := memory.IndexToType(op1)
	op2Type := memory.IndexToType(op2)

	// check semantics with semantic cube
	var resultType types.Type = syntaxCube[op1Type][op2Type][op]

	if resultType == types.Error {
		panic("Error: Invalid operation")
	}

	var quadruple types.Quadruple

	if op == types.Assign {
		quadruple = types.Quadruple{
			Op:     op,
			Arg1:   op2,
			Arg2:   -1,
			Result: op1,
		}
	} else {
		var tempIndex int = memory.AllocateMemory(resultType, types.MemoryType(memory.Global))

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

func EndExpression() (*types.Operator, error) {
	// Pop all of the remaining operators from the stack
	for !OperatorStack.IsEmpty() {
		GenerateQuadruple(OperatorStack.Pop())
	}

	return nil, nil
}
