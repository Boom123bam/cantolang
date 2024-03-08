package main

import (
	"cantolang/evaluator"
	"cantolang/lexer"
	"cantolang/object"
	"cantolang/parser"
	"cantolang/repl"
	"fmt"
	"os"
)

func main() {
	switch len(os.Args) {
	case 1:
		repl.Start(os.Stdin, os.Stdout)
	case 2:
		filename := os.Args[1]
		data, err := os.ReadFile(filename)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}
		// Convert the byte slice to a string
		input := string(data)
		l := lexer.New(input)
		p := parser.New(l)
		program := p.ParseProgram()
		evaluator.Eval(program, object.NewEnvironment(nil))
	default:
		fmt.Println("usage: go run main.go (filename)")
	}
}
