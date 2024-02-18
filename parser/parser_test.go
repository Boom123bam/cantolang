package parser

import (
	"cantolang/ast"
	"cantolang/lexer"
	"cantolang/token"
	"testing"
)

func checkParserErrors(p *Parser, t *testing.T) {
	for i, e := range p.errors {
		t.Errorf("Errors[%d]: %s", i, e)
	}
}

func TestInitializationStatements(t *testing.T) {
	input := `
	叫佢 a。
	叫佢 o_k。
	叫佢 hi。
	`
	expected := []string{"a", "o_k", "hi"}

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(p, t)

	if len(program.Statements) != 3 {
		t.Errorf("len(program) expected 3 got %d", len(program.Statements))
	}

	for i, st := range program.Statements {
		initStatement, ok := st.(*ast.InitializeStatement)
		if !ok {
			t.Errorf("[%d] expected type ast.InitializeStatement got %T", i, st)
		}
		if initStatement.Token.TokenType != token.INITIALIZE {
			t.Errorf("[%d] expected tokenType '%s' got %s", i, token.INITIALIZE, initStatement.Token.TokenType)
		}
		if initStatement.Identifier != expected[i] {
			t.Errorf("[%d] expected identifier '%s' got %s", i, expected[i], initStatement.Identifier)
		}
	}
}

func TestIntegerStatements(t *testing.T) {
	input := `
	2。
	1。
	`
	expected := []string{"2", "1"}

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(p, t)

	if len(program.Statements) != 2 {
		t.Errorf("len(program) expected 3 got %d", len(program.Statements))
	}

	for i, st := range program.Statements {
		exprStatement, ok := st.(*ast.ExpressionStatement)
		if !ok {
			t.Errorf("[%d] expected type ast.ExpressionStatement got %T", i, st)
		}
		if exprStatement.Token.TokenType != token.NUMBER {
			t.Errorf("[%d] expected tokenType '%s' got %s", i, token.NUMBER, exprStatement.Token.TokenType)
		}
		if exprStatement.Token.TokenLiteral != expected[i] {
			t.Errorf("[%d] expected token literal '%s' got %s", i, expected[i], exprStatement.Token.TokenLiteral)
		}
	}
}
