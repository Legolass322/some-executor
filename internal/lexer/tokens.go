package lexer

import "fmt"

type TokenKind int

const (
	EOF TokenKind = iota
	EOL

	NUMBER
	STRING
	IDENTIFIER

	CURLY_START
	CURLY_END
	BRACKET_START
	BRACKET_END
	PARENTHESIS_START
	PARENTHESIS_END

	ASSIGNMENT
	PLUS_EQ
	MINUS_EQ
	STAR_EQ
	SLASH_EQ

	EQUAL
	NOT_EQUAL
	GREATER
	GREATER_EQUAL
	LESS
	LESS_EQUAL

	NOT
	AND
	OR

	PLUS
	MINUS
	STAR
	SLASH

	DOT
	COMMA
	COLON
	QUESTION

	// Reserved keywords
	TRUE
	FALSE
	IMPORT
	EXPORT
	FROM
	AS
	FUNC
	IF
	WHILE
	FOR
	LET
	CONST
	TYPEOF
	RETURN
)

var reservedLookUp map[string]TokenKind = map[string]TokenKind{
	TRUE.String():   TRUE,
	FALSE.String():  FALSE,
	IMPORT.String(): IMPORT,
	EXPORT.String(): EXPORT,
	FROM.String():   FROM,
	AS.String():     AS,
	FUNC.String():   FUNC,
	IF.String():     IF,
	WHILE.String():  WHILE,
	FOR.String():    FOR,
	LET.String():    LET,
	CONST.String():  CONST,
	TYPEOF.String(): TYPEOF,
	RETURN.String(): RETURN,
}

type Token struct {
	Kind  TokenKind
	Value string
}

func NewToken(kind TokenKind, value string) Token {
	return Token{kind, value}
}

func (token Token) Debug() {
	if token.hasOneOfKind(NUMBER, STRING, IDENTIFIER) {
		fmt.Printf("%s (%s)\n", token.Kind, token.Value)
		return
	}

	fmt.Printf("%s ()\n", token.Kind)
}

func (kind TokenKind) IsIn(kinds ...TokenKind) bool {
	for _, expected := range kinds {
		if kind == expected {
			return true
		}
	}

	return false
}

func (token Token) hasOneOfKind(kinds ...TokenKind) bool {
	return token.Kind.IsIn(kinds...)
}

func (kind TokenKind) String() string {
	switch kind {
	case EOF:
		return "eof"
	case EOL:
		return "eol"
	case NUMBER:
		return "number"
	case STRING:
		return "string"
	case TRUE:
		return "true"
	case FALSE:
		return "false"
	case IDENTIFIER:
		return "identifier"
	case CURLY_START:
		return "{"
	case CURLY_END:
		return "}"
	case BRACKET_START:
		return "["
	case BRACKET_END:
		return "]"
	case PARENTHESIS_START:
		return "("
	case PARENTHESIS_END:
		return ")"
	case ASSIGNMENT:
		return "="
	case PLUS_EQ:
		return "+="
	case MINUS_EQ:
		return "-="
	case STAR_EQ:
		return "*="
	case SLASH_EQ:
		return "/="
	case EQUAL:
		return "=="
	case NOT_EQUAL:
		return "!="
	case GREATER:
		return ">"
	case GREATER_EQUAL:
		return ">="
	case LESS:
		return "<"
	case LESS_EQUAL:
		return "<="
	case NOT:
		return "!"
	case AND:
		return "&&"
	case OR:
		return "||"
	case PLUS:
		return "+"
	case MINUS:
		return "-"
	case STAR:
		return "*"
	case SLASH:
		return "/"
	case DOT:
		return "."
	case COMMA:
		return ","
	case COLON:
		return ":"
	case QUESTION:
		return "?"
	case IMPORT:
		return "import"
	case EXPORT:
		return "export"
	case FROM:
		return "from"
	case AS:
		return "as"
	case FUNC:
		return "func"
	case IF:
		return "if"
	case WHILE:
		return "while"
	case FOR:
		return "for"
	case LET:
		return "let"
	case CONST:
		return "const"
	case TYPEOF:
		return "typeof"
	case RETURN:
		return "return"
	default:
		return "unrecognized"
	}
}
