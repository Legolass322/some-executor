package lexer

import (
	"fmt"
	"regexp"
)

type regexHandler func(lex *lexer, regex *regexp.Regexp)

type regexPattern struct {
	regex   *regexp.Regexp
	handler regexHandler
}

type lexer struct {
	source   string
	Tokens   []Token
	pos      int
	patterns []regexPattern
}

func (lex *lexer) advanceN(n int) {
	lex.pos += n
}

func (lex *lexer) push(token Token) {
	lex.Tokens = append(lex.Tokens, token)
}

func (lex *lexer) remainder() string {
	return lex.source[lex.pos:]
}

func (lex *lexer) isEof() bool {
	return lex.pos >= len(lex.source)
}

func newLexer(source string) *lexer {
	return &lexer{
		pos:    0,
		source: source,
		Tokens: make([]Token, 0),
		patterns: []regexPattern{
			{regexp.MustCompile(`\n`), defaultHandler(EOL, "\n")},
			{regexp.MustCompile(`\s+`), skipHandler},
			{regexp.MustCompile(`//.*`), skipHandler},

			{regexp.MustCompile(`[0-9]+(\.[0-9]+)?`), numberHandler},
			{regexp.MustCompile(`"((\\")|[^"])*"`), stringHandler},
			{regexp.MustCompile(`[_a-zA-Z][_a-zA-Z0-9]*`), symbolHandler},

			{regexp.MustCompile(`\{`), defaultKindHandler(CURLY_START)},
			{regexp.MustCompile(`\}`), defaultKindHandler(CURLY_END)},
			{regexp.MustCompile(`\[`), defaultKindHandler(BRACKET_START)},
			{regexp.MustCompile(`\}`), defaultKindHandler(BRACKET_END)},
			{regexp.MustCompile(`\(`), defaultKindHandler(PARENTHESIS_START)},
			{regexp.MustCompile(`\)`), defaultKindHandler(PARENTHESIS_END)},
			{regexp.MustCompile(`==`), defaultKindHandler(EQUAL)},
			{regexp.MustCompile(`!=`), defaultKindHandler(NOT_EQUAL)},
			{regexp.MustCompile(`<=`), defaultKindHandler(LESS_EQUAL)},
			{regexp.MustCompile(`>=`), defaultKindHandler(GREATER_EQUAL)},
			{regexp.MustCompile(`-=`), defaultKindHandler(MINUS_EQ)},
			{regexp.MustCompile(`/=`), defaultKindHandler(SLASH_EQ)},
			{regexp.MustCompile(`\+=`), defaultKindHandler(PLUS_EQ)},
			{regexp.MustCompile(`\*=`), defaultKindHandler(STAR_EQ)},
			{regexp.MustCompile(`!`), defaultKindHandler(NOT)},
			{regexp.MustCompile(`&&`), defaultKindHandler(AND)},
			{regexp.MustCompile(`\|\|`), defaultKindHandler(OR)},
			{regexp.MustCompile(`<`), defaultKindHandler(LESS)},
			{regexp.MustCompile(`>`), defaultKindHandler(GREATER)},
			{regexp.MustCompile(`=`), defaultKindHandler(ASSIGNMENT)},
			{regexp.MustCompile(`\+`), defaultKindHandler(PLUS)},
			{regexp.MustCompile(`-`), defaultKindHandler(MINUS)},
			{regexp.MustCompile(`\*`), defaultKindHandler(STAR)},
			{regexp.MustCompile(`/`), defaultKindHandler(SLASH)},
			{regexp.MustCompile(`\.`), defaultKindHandler(DOT)},
			{regexp.MustCompile(`,`), defaultKindHandler(COMMA)},
			{regexp.MustCompile(`:`), defaultKindHandler(COLON)},
			{regexp.MustCompile(`\?`), defaultKindHandler(QUESTION)},
		},
	}
}

func Tokenize(source string) []Token {
	lex := newLexer(source)

	for !lex.isEof() {
		matched := false

		for _, pattern := range lex.patterns {
			loc := pattern.regex.FindStringIndex(lex.remainder())

			if loc != nil && loc[0] == 0 {
				pattern.handler(lex, pattern.regex)
				matched = true
				break
			}
		}

		// todo: write unrecognized and go to the next closest token
		if !matched {
			panic(fmt.Sprintf("Lexer::Error -> unrecognized token near %s\n", lex.remainder()))
		}
	}

	lex.push(NewToken(EOF, "EOF"))

	return lex.Tokens
}

func defaultHandler(kind TokenKind, value string) regexHandler {
	return func(lex *lexer, regex *regexp.Regexp) {
		lex.advanceN(len(value))
		lex.push(NewToken(kind, value))
	}
}

func defaultKindHandler(kind TokenKind) regexHandler {
	return defaultHandler(kind, kind.String())
}

func numberHandler(lex *lexer, regex *regexp.Regexp) {
	match := regex.FindString(lex.remainder())
	lex.push(NewToken(NUMBER, match))
	lex.advanceN(len(match))
}

func stringHandler(lex *lexer, regex *regexp.Regexp) {
	match := regex.FindStringIndex(lex.remainder())
	stringLiteral := lex.remainder()[match[0]+1 : match[1]-1]
	lex.push(NewToken(STRING, stringLiteral))
	lex.advanceN(len(stringLiteral) + 2)
}

func symbolHandler(lex *lexer, regex *regexp.Regexp) {
	match := regex.FindString(lex.remainder())

	if kind, exists := reservedLookUp[match]; exists {
		lex.push(NewToken(kind, match))
	} else {
		lex.push(NewToken(IDENTIFIER, match))
	}

	lex.advanceN(len(match))
}

func skipHandler(lex *lexer, regex *regexp.Regexp) {
	match := regex.FindString(lex.remainder())
	lex.advanceN(len(match))
}
