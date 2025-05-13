package ast

import "compiler/types"

var QuadrupleList []types.Quadruple

func AddQuadruple(quadruple types.Quadruple) {
	QuadrupleList = append(QuadrupleList, quadruple)
}

func GetQuadrupleList() []types.Quadruple {
	return QuadrupleList
}
