package parser

import "github.com/Legolass322/executor/internal/lexer"

type ParseletManager struct {
	Prefix *prefixParseletsHandler
	Infix  *infixParseletsHandler
}

var manager *ParseletManager = nil

func newParseletManager() *ParseletManager {
	if manager == nil {
		manager = &ParseletManager{
			Prefix: &prefixParseletsHandler{
				parselets: make(map[lexer.TokenKind]prefixParselet),
			}, 
			Infix: &infixParseletsHandler{
				parselets: make(map[lexer.TokenKind]infixParselet),
			},
		}
		manager.Prefix.initParselets()
		manager.Infix.initParselets()
	}

	return manager
}

func (pp *prefixParseletsHandler) initParselets() {
	pp.register(lexer.IDENTIFIER, nameParselet{})
	pp.register(lexer.NUMBER, numberParselet{})
	pp.prefix(lexer.PLUS)
	pp.prefix(lexer.MINUS)
	pp.prefix(lexer.NOT)

	pp.register(lexer.PARENTHESIS_START, parenParselet{})
}

func (ip *infixParseletsHandler) initParselets() {
	ip.infix(lexer.PLUS)
	ip.infix(lexer.MINUS)
	ip.infix(lexer.SLASH)
	ip.infix(lexer.STAR)
	ip.infix(lexer.EQUAL)
	ip.infix(lexer.AND)
	ip.infix(lexer.OR)
	ip.infix(lexer.GREATER)
	ip.infix(lexer.GREATER_EQUAL)
	ip.infix(lexer.LESS)
	ip.infix(lexer.LESS_EQUAL)
}
