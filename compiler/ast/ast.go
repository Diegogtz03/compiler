package ast

import "fmt"

type TestNode struct {
	Message string
}

func TestAST() *TestNode {
	fmt.Println("TestAST")
	return &TestNode{Message: "Hello, World!"}
}
