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
		t.Errorf("len(program) expected 2 got %d", len(program.Statements))
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

func TestPrefixStatements(t *testing.T) {
	input := `
	-2。
	-1。
	唔係5。
	`
	expected := []struct {
		prefix string
		right  string
	}{
		{"-", "2"},
		{"-", "1"},
		{"唔係", "5"},
	}

	l := lexer.New(input)
	p := New(l)
	// p.ParseProgram()
	program := p.ParseProgram()
	checkParserErrors(p, t)

	if len(program.Statements) != 3 {
		t.Errorf("len(program) expected 3 got %d", len(program.Statements))
	}

	for i, st := range program.Statements {
		exprStatement, ok := st.(*ast.ExpressionStatement)
		if !ok {
			t.Errorf("[%d] expected type ast.ExpressionStatement got %T", i, st)
		}
		prefixExp, ok := exprStatement.Expression.(*ast.PrefixExpression)
		if !ok {
			t.Errorf("[%d] expected type ast.PrefixExpression got %T", i, exprStatement.Expression)
		}
		right, ok := prefixExp.Right.(*ast.IntegerLiteral)
		if !ok {
			t.Errorf("[%d] expected type ast.IntegerLiteral got %T", i, prefixExp.Right)
		}
		if prefixExp.PrefixToken.TokenLiteral != expected[i].prefix {
			t.Errorf("[%d] expected prefix '%s' got %s", i, expected[i].prefix, prefixExp.PrefixToken.TokenLiteral)
		}
		if right.Token.TokenLiteral != expected[i].right {
			t.Errorf("[%d] expected token literal '%s' got %s", i, expected[i], exprStatement.Token.TokenLiteral)
		}
	}
}

func TestInfixStatements(t *testing.T) {
	input := `
	1-1。
	1+2+3。
	1+3*2/5。
	`
	expected := []string{
		"(1 - 1)",
		"((1 + 2) + 3)",
		"(1 + ((3 * 2) / 5))",
	}

	l := lexer.New(input)
	p := New(l)
	// p.ParseProgram()
	program := p.ParseProgram()
	checkParserErrors(p, t)
	if len(program.Statements) != 3 {
		t.Errorf("len(program) expected 3 got %d", len(program.Statements))
	}

	for i, st := range program.Statements {
		exprStatement, ok := st.(*ast.ExpressionStatement)
		if !ok {
			t.Errorf("[%d] expected type ast.ExpressionStatement got %T", i, st)
		}
		if exprStatement.Expression.String() != expected[i] {
			t.Errorf("[%d] expected string %s got %s", i, expected[i], exprStatement.Expression.String())
		}
	}

}
