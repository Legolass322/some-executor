package main

import (
	"flag"
	"fmt"
	"os"
	"path"

	"github.com/Legolass322/executor/internal/lexer"
	"github.com/Legolass322/executor/internal/parser"
	"github.com/sanity-io/litter"
)

func main() {
	var rootPath string
    flag.StringVar(&rootPath, "r", ".", "Provide project path as an absolute path")
    flag.Parse()
	
	fmt.Println(rootPath)
	
	bytes, err := os.ReadFile(path.Join(rootPath, "examples", "03.e"))

	if err != nil {
		panic(err.Error())
	}

	source := string(bytes)
	tokens := lexer.Tokenize(source)

	fmt.Println("Tree:")

	ast := parser.ParseExp(tokens)
	litter.Dump(ast)
}
