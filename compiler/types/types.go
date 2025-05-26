package types

type MemoryType int

const (
	Global MemoryType = iota
	Local
	Temp
	Constant
)

type Type int

const (
	Int Type = iota
	Float
	Bool
	String
	Error
)

type Quadruple struct {
	Op     Operator
	Arg1   int
	Arg2   int
	Result int
}

type Operator int

const (
	Add Operator = iota
	Sub
	Mul
	Div
	Assign
	NotEqual
	LessThan
	GreaterThan
	Print
	StackDivider
	ErrorOperator
	Goto
	GotoF
	GotoT
	GoSub
	Era
	Parameter
	EndFunc
	Terminate
)
