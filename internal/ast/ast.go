package ast

type Statement interface {
	stmt()
}

type Expression interface {
	Expr() ExprType
}
