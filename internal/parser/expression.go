package parser

import "github.com/Legolass322/executor/internal/ast"

func (p *parser) parseExpression(bpower int) ast.Expression {
	token := p.consume()

	prefix, existsPrefix := manager.Prefix.parselet(token.Kind)

	if !existsPrefix {
		panic("not exists prefix: " + token.Kind.String())
	}

	left := prefix.parse(p, token)

	for int(bpower) < getBPower(p) {
		token = p.consume()

		infix, _ := manager.Infix.parselet(token.Kind)
		left = infix.parse(p, left, token)
	}
	
	return left
}

func getBPower(p *parser) int {
	token := p.peek()
	parselet, exists := manager.Infix.parselet(token.Kind)

	if !exists {
		return -1
	}

	return int(parselet.getBPower())
}
