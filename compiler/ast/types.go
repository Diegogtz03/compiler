package ast

type Node interface {
	Token() string
}

type Program struct {
	Name  string
	Vars  []*Vars
	Funcs []*Func
	// ProgramBody *Body
}

type Vars struct {
	Names []string
	Type  string
}

type Func struct {
	Name   string
	Params []*Vars
	Vars   []*Vars
	// Body   *Body
}

type Body struct {
	Statements []Statement
}

type Statement interface {
	Node
	statement()
}

type StringExpression struct {
	Value string
}

type AssignStatement struct {
	Var  string
	Expr Exp
}

type PrintExpression interface {
	Node
	expression()
}

type PrintStatement struct {
	Expr []*PrintExpression
}

type CycleStatement struct {
	Condition Expression
	Body      Body
}

type ConditionStatement struct {
	Condition Expression
	Body      Body
	ElseBody  Body
}

type FCallStatement struct {
	Id   string
	Args []Expression
}

type Expression interface {
	Node
	expression()
}

type BooleanExpression struct {
	Left  Exp
	Op    string
	Right Exp
}

type IntConstant struct {
	Value int
}

type FloatConstant struct {
	Value float64
}
