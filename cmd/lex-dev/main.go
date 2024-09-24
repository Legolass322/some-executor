package main

import (
	"os"

	"github.com/Legolass322/executor/internal/lexer"
)

func main() {
	bytes, err := os.ReadFile("./examples/00.e")

	if err != nil {
		panic(err.Error())
	}

	source := string(bytes)
	tokens := lexer.Tokenize(source)

	for _, token := range tokens {
		token.Debug()
	}
}
