package parser

import (
	"strconv"

	"github.com/Legolass322/executor/internal/ast"
	"github.com/Legolass322/executor/internal/lexer"
)

// PARSELETS DEFINITION

type prefixParselet interface {
	parse(p *parser, token lexer.Token) ast.Expression
}

type prefixOperatorParselet struct{
	bpower bindingPower
}

func (parselet *prefixOperatorParselet) getBPower() int {
	return int(parselet.bpower)
}

func (parselet *prefixOperatorParselet) parse(p *parser, token lexer.Token) ast.Expression {
	operand := p.parseExpression(parselet.getBPower())
	return &ast.PrefixUnaryExpr{Operator: token, Only: operand}
}

type nameParselet struct{}

func (parselet nameParselet) parse(p *parser, token lexer.Token) ast.Expression {
	return &ast.NameExpr{Token: token}
}

type numberParselet struct{}

func (parselet numberParselet) parse(p *parser, token lexer.Token) ast.Expression {
	// todo
	value, err := strconv.Atoi(token.Value);

	if err != nil {
		token.Debug()
		panic("Cannot convert number")
	}

	return &ast.NumberExpr{Value: value}
}

type parenParselet struct{}

func (parselet parenParselet) parse(p *parser, token lexer.Token) ast.Expression {
	inner := p.parseExpression(0)
	_, err := p.consumeOnly(lexer.PARENTHESIS_END)

	// todo
	if err != nil {
		panic(err)
	}

	return inner
}

// PARSELETS TABLE

type prefixParseletsHandler struct {
	parselets map[lexer.TokenKind]prefixParselet
}

func (pp *prefixParseletsHandler) register(tKind lexer.TokenKind, parselet prefixParselet) {
	pp.parselets[tKind] = parselet
}

func (pp *prefixParseletsHandler) prefix(tKind lexer.TokenKind) {
	pp.register(tKind, &prefixOperatorParselet{bpower: PREFIX})
}

func (pp *prefixParseletsHandler) parselet(tKind lexer.TokenKind) (prefixParselet, bool) {
	parselet, exists := pp.parselets[tKind]
	return parselet, exists
}
