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
	Op     string
	Arg1   int
	Arg2   int
	Result int
}
