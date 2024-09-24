package ast

import "github.com/Legolass322/executor/internal/lexer"

type ExprType int

const (
	EXPR_NUMBER ExprType = iota
	EXPR_STRING
	EXPR_TRUE
	EXPR_FALSE
	EXPR_BINARY
	EXPR_UNARY
	EXPR_NAME
)

// LITERALS

// todo: FloatExpr, IntExpr
type NumberExpr struct {
	Value int
}

func (e NumberExpr) Expr() ExprType {
	return EXPR_NUMBER
}

type StringExpr struct {
	Value string
}

func (e StringExpr) Expr() ExprType {
	return EXPR_STRING
}

type TrueExpr struct {
	Value string
}

func (e TrueExpr) Expr() ExprType {
	return EXPR_TRUE
}

type FalseExpr struct {
	Value string
}

func (e FalseExpr) Expr() ExprType {
	return EXPR_FALSE
}

// OPERATIONS

type InfixBinaryExpr struct {
	Left     Expression
	Operator lexer.Token
	Right    Expression
}

func (e InfixBinaryExpr) Expr() ExprType {
	return EXPR_BINARY
}

type PrefixUnaryExpr struct {
	Operator lexer.Token
	Only     Expression
}

func (e PrefixUnaryExpr) Expr() ExprType {
	return EXPR_UNARY
}

// NAMES

type NameExpr struct {
	Token lexer.Token
}

func (e NameExpr) Expr() ExprType {
	return EXPR_NAME
}
