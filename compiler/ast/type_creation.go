package ast

import (
	"compiler/token"
	"fmt"
	"strconv"
)

func CreateProgram(name, vars, funcs interface{}) (*Program, error) {
	nameStr := string(name.(*token.Token).Lit)
	var v = []*Vars{}
	var f = []*Func{}

	fmt.Println("funcs", funcs)
	fmt.Println("vars", vars)

	if funcs != nil {
		funcsStr := funcs.([]*Func)
		f = funcsStr
	}

	if vars != nil {
		fmt.Println("vars", vars)
		varsStr := vars.([]*Vars)
		v = varsStr
	}

	// bodyStr := body.(*Body)

	return &Program{
		Name:  nameStr,
		Vars:  v,
		Funcs: f,
		// ProgramBody: bodyStr,
	}, nil
}

func CreateFunc(name interface{}) (*Func, error) {
	nameStr := string(name.(*token.Token).Lit)
	// paramsStr := params.([]*Vars)
	// varsStr := vars.([]*Vars)
	// bodyStr := body.(*Body)

	return &Func{
		Name: nameStr,
		// Params: paramsStr,
		// Vars:   varsStr,
		// Body:   bodyStr,
	}, nil
}

func CreateVars(names, var_type interface{}) *Vars {
	namesStr := names.([]string)
	var_typeStr := string(var_type.(*token.Token).Lit)

	return &Vars{
		Names: namesStr,
		Type:  var_typeStr,
	}
}

// CONSTANTS

func CreateIntConstant(value interface{}) (*IntConstant, error) {
	valueStr := string(value.(*token.Token).Lit)

	intValue, err := strconv.Atoi(valueStr)

	if err != nil {
		return nil, err
	}

	return &IntConstant{
		Value: intValue,
	}, nil
}

func CreateFloatConstant(value interface{}) (*FloatConstant, error) {
	valueStr := string(value.(*token.Token).Lit)

	floatValue, err := strconv.ParseFloat(valueStr, 64)

	if err != nil {
		return nil, err
	}

	return &FloatConstant{
		Value: floatValue,
	}, nil
}
