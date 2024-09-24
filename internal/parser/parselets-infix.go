package parser

import (
	"github.com/Legolass322/executor/internal/ast"
	"github.com/Legolass322/executor/internal/lexer"
)

// PARSELETS DEFINITION

type infixParselet interface {
	parse(p *parser, left ast.Expression, token lexer.Token) ast.Expression
	getBPower() int
}

type binaryOperationParselet struct{
	bpower bindingPower
}

func (ip *binaryOperationParselet) parse(p *parser, left ast.Expression, token lexer.Token) ast.Expression {
	right := p.parseExpression(ip.getBPower())
	return &ast.InfixBinaryExpr{Left: left, Operator: token, Right: right}
}

func (ip *binaryOperationParselet) getBPower() int {
	return int(ip.bpower)
}

// PARSELETS TABLE

type infixParseletsHandler struct {
	parselets map[lexer.TokenKind]infixParselet
}

func (ip *infixParseletsHandler) register(tKind lexer.TokenKind, parselet infixParselet) {
	ip.parselets[tKind] = parselet
}

func (ip *infixParseletsHandler) infix(tKind lexer.TokenKind) {
	ip.register(tKind, &binaryOperationParselet{bpower: infixBPowerTable[tKind]})
}

func (ip *infixParseletsHandler) parselet(tKind lexer.TokenKind) (infixParselet, bool) {
	parselet, exists := ip.parselets[tKind]
	return parselet, exists
}
