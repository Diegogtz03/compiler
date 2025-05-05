package ast

// import (
// 	"compiler/token"
// 	"compiler/utils"
// 	"errors"
// 	"fmt"
// )

// type Variable struct {
// 	Id    string
// 	Type  string
// 	Value string
// }

// var varsQueue = utils.Queue{}

// var syntaxCube = map[string]map[string]map[string]string{
// 	"int": {
// 		"int": {
// 			"+": "int",
// 			"-": "int",
// 			"*": "int",
// 			"/": "int",
// 		},
// 		"float": {
// 			"+": "float",
// 			"-": "float",
// 			"*": "float",
// 			"/": "float",
// 		},
// 	},
// 	"float": {
// 		"int": {
// 			"+": "float",
// 			"-": "float",
// 			"*": "float",
// 			"/": "float",
// 		},
// 		"float": {
// 			"+": "float",
// 			"-": "float",
// 			"*": "float",
// 			"/": "float",
// 		},
// 	},
// }

// func AddVarToQueue(stmt interface{}) (Variable, error) {
// 	id := string(stmt.(*token.Token).Lit)

// 	// add to queue
// 	varsQueue.Enqueue(id)
// 	if _, ok := ProgramFunctions[CurrentModule].Vars[id]; ok {
// 		return Variable{}, errors.New("Variable " + id + " already exists in module " + CurrentModule)
// 	}

// 	var newVar = Variable{
// 		Id:   id,
// 		Type: "int",
// 	}

// 	ProgramFunctions[CurrentModule].Vars[id] = newVar

// 	return newVar, nil
// }

// func setCurrentType(stmt interface{}) (string, error) {
// 	var_type := string(stmt.(*token.Token).Lit)

// 	return var_type, nil
// }

// func AddVarsToTable(stmt interface{}) (Variable, error) {
// 	var_type := string(stmt.(*token.Token).Lit)

// 	for !varsQueue.IsEmpty() {
// 		id := varsQueue.Dequeue()

// 		if _, ok := ProgramFunctions[CurrentModule].Vars[id]; ok {
// 			return Variable{}, fmt.Errorf("Variable " + id + " already exists in module " + CurrentModule)
// 		}

// 		var newVar = Variable{
// 			Id:   id,
// 			Type: var_type,
// 		}

// 		ProgramFunctions[CurrentModule].Vars[id] = newVar
// 	}

// 	return Variable{}, nil
// }

// // current_module
// // current_type
// // current_id
// // current_operator
// // current_sign
// // current_cte

// // 0 -> Crear Dir Funciones
// // 1 -> Dir Funciones add id and type // set modulo actual
// // 2 -> Crear Tabla Vars, ligada a módulo actual
// // 3 -> Set tipo actual
// // 4 -> If var in table --> ERROR else --> Add tabla or queue vars (todo lo necesario) [TYPE NULL or queue with ids]
// // 5 -> Set tipo actual void
// // 6 -> Set list of ids with type or all with null types with current types (add queue to var tables)
// // 7 -> Verify if ID exists in modulo actual, set id actual
// // 8 -> Assign expression result to id actual
// // 9 -> if theres an operator, use it, Assign result of term to current_term, current operator to null
// // 10 --> set current operator
// // 11 --> set current sign
// // 12 --> get given id from var table from module --> assign to current_cte
// // 13 --> set current cte
// // 14 -->

// On program node --> assign to func table --> current module == "global"
// On func node --> assign to func table --> current module == func name -- if exists --> error
// On vars node --> assign to vars table --> current module == func name --> if exists --> error

// on expression node --> check matrix of types to see if the operation is valid
