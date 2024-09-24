package parser

import (
	"fmt"

	"github.com/Legolass322/executor/internal/ast"
	"github.com/Legolass322/executor/internal/lexer"
)

type parser struct {
	//todo
	// errors []error

	pos    int
	tokens []lexer.Token
}

func newParser(tokens []lexer.Token) *parser {
	newParseletManager()
	return &parser{pos: 0, tokens: tokens}
}

func Parse(tokens []lexer.Token) ast.BlockStatement {
	Body := make([]ast.Statement, 0)

	p := newParser(tokens)

	for p.hasTokens() {
		Body = append(Body, parse_stmt(p))
	}

	return ast.BlockStatement{Body: Body}
}

func ParseExp(tokens []lexer.Token) []ast.Expression {
	p := newParser(tokens)

	expressions := make([]ast.Expression, 0)

	for p.hasTokens() {
		// todo -1
		expressions = append(expressions, p.parseExpression(-1))
	}

	return expressions
}

func (p *parser) peek() lexer.Token {
	return p.tokens[p.pos]
}

func (p *parser) peekKind() lexer.TokenKind {
	return p.tokens[p.pos].Kind
}

func (p *parser) consume() lexer.Token {
	t := p.peek()
	p.pos++
	return t
}

func (p *parser) consumeOnly(tKind lexer.TokenKind) (lexer.Token, error) {
	t := p.consume()

	if t.Kind == tKind {
		return t, nil
	}

	return t, fmt.Errorf("expected %s, but received %s", tKind.String(), t.Kind.String())
}

func (p *parser) hasTokens() bool {
	return p.pos < len(p.tokens) && p.peekKind() != lexer.EOF
}
