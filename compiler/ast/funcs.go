package ast

// type Function struct {
// 	Name string
// 	Vars map[string]Variable
// }

// var ProgramFunctions map[string]Function = make(map[string]Function)

// var CurrentModule string = ""

// // 1
// func CreateFunc(stmt interface{}) (Function, error) {
// 	id := string(stmt.(*token.Token).Lit)

// 	if _, ok := ProgramFunctions[id]; ok {
// 		return Function{}, fmt.Errorf("Function " + id + " already exists")
// 	}

// 	fmt.Println("Creating function " + id)

// 	var newFunc = Function{
// 		Name: id,
// 		Vars: make(map[string]Variable),
// 	}

// 	ProgramFunctions[id] = newFunc
// 	CurrentModule = id

// 	return newFunc, nil
// }
