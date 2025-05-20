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
	types.Goto:          5,
	types.GotoF:         5,
	types.GotoT:         5,
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
	fmt.Printf("Generating quadruple for %v\n", op)
	fmt.Printf("OperandStack: %v\n", OperandStack)
	fmt.Printf("OperatorStack: %v\n", OperatorStack)
	fmt.Println("--------------------------------")

	op2, op1, resultType := getOperators(op)

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
	} else {
		if resultType == types.Error {
			panic("Error: Invalid operation")
		}

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

func getOperators(op types.Operator) (int, int, types.Type) {
	if op == types.Print {
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

func GenerateGoto(op types.Operator) {
	JumpStack.Push(len(QuadrupleList) - 1)

	PushOperator(op)
}

func PopJumpStack() {
	quadrupleIndex := JumpStack.Pop()

	quadruple := QuadrupleList[quadrupleIndex]

	quadruple.Result = len(QuadrupleList)

	QuadrupleList[quadrupleIndex] = quadruple
}

// TODO:
// - Add constants to memory (negatives?)
// - Add if statement
// - Add while statement
