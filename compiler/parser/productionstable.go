package parser

import "compiler/ast"

type (
	ProdTab      [numProductions]ProdTabEntry
	ProdTabEntry struct {
		String     string
		Id         string
		NTType     int
		Index      int
		NumSymbols int
		ReduceFunc func([]Attrib, interface{}) (Attrib, error)
	}
	Attrib interface {
	}
)

var productionsTable = ProdTab{
	ProdTabEntry{
		String:     `S' : Start	<<  >>`,
		Id:         "S'",
		NTType:     0,
		Index:      0,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `Start : Programa	<<  >>`,
		Id:         "Start",
		NTType:     1,
		Index:      1,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `Program_Point : program	<< ast.TestAST(), nil >>`,
		Id:         "Program_Point",
		NTType:     2,
		Index:      2,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.TestAST(), nil
		},
	},
	ProdTabEntry{
		String:     `Programa : Program_Point id semicolon Vars Programa_PR main Body end	<<  >>`,
		Id:         "Programa",
		NTType:     3,
		Index:      3,
		NumSymbols: 8,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `Programa : Program_Point id semicolon Programa_PR main Body end	<<  >>`,
		Id:         "Programa",
		NTType:     3,
		Index:      4,
		NumSymbols: 7,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `Programa_PR : Funcs Programa_PR	<<  >>`,
		Id:         "Programa_PR",
		NTType:     4,
		Index:      5,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `Programa_PR : empty	<<  >>`,
		Id:         "Programa_PR",
		NTType:     4,
		Index:      6,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return nil, nil
		},
	},
	ProdTabEntry{
		String:     `Vars : var id Vars_PR two_dots Type semicolon Vars_PR_PR	<<  >>`,
		Id:         "Vars",
		NTType:     5,
		Index:      7,
		NumSymbols: 7,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `Vars_PR : comma id Vars_PR	<<  >>`,
		Id:         "Vars_PR",
		NTType:     6,
		Index:      8,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `Vars_PR : empty	<<  >>`,
		Id:         "Vars_PR",
		NTType:     6,
		Index:      9,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return nil, nil
		},
	},
	ProdTabEntry{
		String:     `Vars_PR_PR : id Vars_PR two_dots Type semicolon Vars_PR_PR	<<  >>`,
		Id:         "Vars_PR_PR",
		NTType:     7,
		Index:      10,
		NumSymbols: 6,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `Vars_PR_PR : empty	<<  >>`,
		Id:         "Vars_PR_PR",
		NTType:     7,
		Index:      11,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return nil, nil
		},
	},
	ProdTabEntry{
		String:     `Type : int_rw	<<  >>`,
		Id:         "Type",
		NTType:     8,
		Index:      12,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `Type : float_rw	<<  >>`,
		Id:         "Type",
		NTType:     8,
		Index:      13,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `Body : curly_open Body_PR curly_close	<<  >>`,
		Id:         "Body",
		NTType:     9,
		Index:      14,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `Body_PR : Statement Body_PR	<<  >>`,
		Id:         "Body_PR",
		NTType:     10,
		Index:      15,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `Body_PR : empty	<<  >>`,
		Id:         "Body_PR",
		NTType:     10,
		Index:      16,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return nil, nil
		},
	},
	ProdTabEntry{
		String:     `Statement : Assign	<<  >>`,
		Id:         "Statement",
		NTType:     11,
		Index:      17,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `Statement : Condition	<<  >>`,
		Id:         "Statement",
		NTType:     11,
		Index:      18,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `Statement : Cycle	<<  >>`,
		Id:         "Statement",
		NTType:     11,
		Index:      19,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `Statement : F_call	<<  >>`,
		Id:         "Statement",
		NTType:     11,
		Index:      20,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `Statement : Print	<<  >>`,
		Id:         "Statement",
		NTType:     11,
		Index:      21,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `Print : print parenthesis_open Expr Print_PR parenthesis_close semicolon	<<  >>`,
		Id:         "Print",
		NTType:     12,
		Index:      22,
		NumSymbols: 6,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `Print : print parenthesis_open string Print_PR parenthesis_close semicolon	<<  >>`,
		Id:         "Print",
		NTType:     12,
		Index:      23,
		NumSymbols: 6,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `Print_PR : comma Expr Print_PR	<<  >>`,
		Id:         "Print_PR",
		NTType:     13,
		Index:      24,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `Print_PR : comma string Print_PR	<<  >>`,
		Id:         "Print_PR",
		NTType:     13,
		Index:      25,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `Print_PR : empty	<<  >>`,
		Id:         "Print_PR",
		NTType:     13,
		Index:      26,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return nil, nil
		},
	},
	ProdTabEntry{
		String:     `Assign : id equal Exp semicolon	<<  >>`,
		Id:         "Assign",
		NTType:     14,
		Index:      27,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `Cycle : while parenthesis_open Expr parenthesis_close do Body semicolon	<<  >>`,
		Id:         "Cycle",
		NTType:     15,
		Index:      28,
		NumSymbols: 7,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `Condition : if parenthesis_open Expr parenthesis_close Body semicolon	<<  >>`,
		Id:         "Condition",
		NTType:     16,
		Index:      29,
		NumSymbols: 6,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `Condition : if parenthesis_open Expr parenthesis_close Body else Body semicolon	<<  >>`,
		Id:         "Condition",
		NTType:     16,
		Index:      30,
		NumSymbols: 8,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `Cte : myint	<<  >>`,
		Id:         "Cte",
		NTType:     17,
		Index:      31,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `Cte : myfloat	<<  >>`,
		Id:         "Cte",
		NTType:     17,
		Index:      32,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `Expr : Exp	<<  >>`,
		Id:         "Expr",
		NTType:     18,
		Index:      33,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `Expr : Exp greater_than Exp	<<  >>`,
		Id:         "Expr",
		NTType:     18,
		Index:      34,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `Expr : Exp less_than Exp	<<  >>`,
		Id:         "Expr",
		NTType:     18,
		Index:      35,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `Expr : Exp not_equal Exp	<<  >>`,
		Id:         "Expr",
		NTType:     18,
		Index:      36,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `Exp : Term	<<  >>`,
		Id:         "Exp",
		NTType:     19,
		Index:      37,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `Exp : Term plus Exp	<<  >>`,
		Id:         "Exp",
		NTType:     19,
		Index:      38,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `Exp : Term minus Exp	<<  >>`,
		Id:         "Exp",
		NTType:     19,
		Index:      39,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `Term : Fact	<<  >>`,
		Id:         "Term",
		NTType:     20,
		Index:      40,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `Term : Fact mult Term	<<  >>`,
		Id:         "Term",
		NTType:     20,
		Index:      41,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `Term : Fact div Term	<<  >>`,
		Id:         "Term",
		NTType:     20,
		Index:      42,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `Fact : parenthesis_open Expr parenthesis_close	<<  >>`,
		Id:         "Fact",
		NTType:     21,
		Index:      43,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `Fact : plus id	<<  >>`,
		Id:         "Fact",
		NTType:     21,
		Index:      44,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `Fact : minus id	<<  >>`,
		Id:         "Fact",
		NTType:     21,
		Index:      45,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `Fact : Cte	<<  >>`,
		Id:         "Fact",
		NTType:     21,
		Index:      46,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `Fact : id	<<  >>`,
		Id:         "Fact",
		NTType:     21,
		Index:      47,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `Fact : plus Cte	<<  >>`,
		Id:         "Fact",
		NTType:     21,
		Index:      48,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `Fact : minus Cte	<<  >>`,
		Id:         "Fact",
		NTType:     21,
		Index:      49,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `Funcs : void id parenthesis_open parenthesis_close bracket_open Body bracket_close semicolon	<<  >>`,
		Id:         "Funcs",
		NTType:     22,
		Index:      50,
		NumSymbols: 8,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `Funcs : void id parenthesis_open parenthesis_close bracket_open Vars Body bracket_close semicolon	<<  >>`,
		Id:         "Funcs",
		NTType:     22,
		Index:      51,
		NumSymbols: 9,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `Funcs : void id parenthesis_open id two_dots Type Funcs_PR parenthesis_close bracket_open Body bracket_close semicolon	<<  >>`,
		Id:         "Funcs",
		NTType:     22,
		Index:      52,
		NumSymbols: 12,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `Funcs : void id parenthesis_open id two_dots Type Funcs_PR parenthesis_close bracket_open Vars Body bracket_close semicolon	<<  >>`,
		Id:         "Funcs",
		NTType:     22,
		Index:      53,
		NumSymbols: 13,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `Funcs_PR : comma id two_dots Type Funcs_PR	<<  >>`,
		Id:         "Funcs_PR",
		NTType:     23,
		Index:      54,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `Funcs_PR : empty	<<  >>`,
		Id:         "Funcs_PR",
		NTType:     23,
		Index:      55,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return nil, nil
		},
	},
	ProdTabEntry{
		String:     `F_call : id parenthesis_open parenthesis_close semicolon	<<  >>`,
		Id:         "F_call",
		NTType:     24,
		Index:      56,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `F_call : id parenthesis_open Expr F_call_PR parenthesis_close semicolon	<<  >>`,
		Id:         "F_call",
		NTType:     24,
		Index:      57,
		NumSymbols: 6,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `F_call_PR : comma Expr F_call_PR	<<  >>`,
		Id:         "F_call_PR",
		NTType:     25,
		Index:      58,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `F_call_PR : empty	<<  >>`,
		Id:         "F_call_PR",
		NTType:     25,
		Index:      59,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return nil, nil
		},
	},
}
