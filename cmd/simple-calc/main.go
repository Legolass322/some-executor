package main

import (
	"flag"
	"fmt"
	"os"
	"path"

	"github.com/Legolass322/executor/internal/ast"
	"github.com/Legolass322/executor/internal/lexer"
	"github.com/Legolass322/executor/internal/parser"
)

func main() {
	var rootPath string
	flag.StringVar(&rootPath, "r", ".", "Provide project path as an absolute path")
	flag.Parse()

	bytes, err := os.ReadFile(path.Join(rootPath, "examples", "calc.e"))

	if err != nil {
		panic(err.Error())
	}

	source := string(bytes)
	tokens := lexer.Tokenize(source)

	ast := parser.ParseExp(tokens)[0]

	fmt.Println(calc(ast))

}

func calc(expr ast.Expression) float64 {
	switch expr := expr.(type) {
	case nil:
		panic("Nil expr")
	case *ast.NumberExpr:
		return float64(expr.Value)
	case *ast.PrefixUnaryExpr:

		if !expr.Operator.Kind.IsIn(lexer.PLUS, lexer.MINUS) {
			panic("Unexpected token: " + expr.Operator.Kind.String())
		}

		if expr.Operator.Kind == lexer.PLUS {
			return calc(expr.Only)
		}

		return -calc(expr.Only)
	case *ast.InfixBinaryExpr:
		if !expr.Operator.Kind.IsIn(lexer.PLUS, lexer.MINUS, lexer.SLASH, lexer.STAR) {
			panic("Unexpected token: " + expr.Operator.Kind.String())
		}

		switch expr.Operator.Kind {
		case lexer.PLUS:
			return calc(expr.Left) + calc(expr.Right)
		case lexer.MINUS:
			return calc(expr.Left) - calc(expr.Right)
		case lexer.STAR:
			return calc(expr.Left) * calc(expr.Right)
		case lexer.SLASH:
			dividor := calc(expr.Right)
			if dividor < 0.00005 && dividor > -0.00005 {
				panic("Cannot divide by zero")
			}
			return calc(expr.Left) / dividor
		}
	}

	panic("Unexpected")
}
