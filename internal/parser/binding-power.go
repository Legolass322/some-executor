package parser

import "github.com/Legolass322/executor/internal/lexer"

type bindingPower int

const (
	ASSIGNMENT bindingPower = iota
	CONDITIONAL
	SUM
	PRODUCT
	PREFIX
	INFIX
	CALL
)

// INFIX BPOWER TABLE

var infixBPowerTable = map[lexer.TokenKind]bindingPower{
	lexer.EQUAL: CONDITIONAL,
	lexer.AND: CONDITIONAL,
	lexer.OR: CONDITIONAL,
	lexer.GREATER: CONDITIONAL,
	lexer.GREATER_EQUAL: CONDITIONAL,
	lexer.LESS: CONDITIONAL,
	lexer.LESS_EQUAL: CONDITIONAL,
	lexer.PLUS: SUM,
	lexer.MINUS: SUM,
	lexer.STAR: PRODUCT,
	lexer.SLASH: PRODUCT,
}
