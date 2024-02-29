package parser

import (
	"cantolang/ast"
	"cantolang/lexer"
	"cantolang/token"
	"testing"
)

func checkParserErrors(p *Parser, t *testing.T) {
	for i, e := range p.Errors {
		t.Errorf("Errors[%d]: %s", i, e)
	}
}

func TestIntegerStatements(t *testing.T) {
	input := `
	2。
	1。
	`
	expected := []int{2, 1}

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
		intLit, ok := exprStatement.Expression.(*ast.IntegerLiteral)
		if intLit.Value != expected[i] {
			t.Errorf("[%d] expected value '%d' got %d", i, expected[i], intLit.Value)
		}
	}
}

func TestBoolStatements(t *testing.T) {
	input := `
	啱。
	錯。
	`
	expected := []bool{true, false}

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
		bool, ok := exprStatement.Expression.(*ast.Boolean)
		if bool.Value != expected[i] {
			t.Errorf("[%d] expected value '%t' got %t", i, expected[i], bool.Value)
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
		right  int
	}{
		{"-", 2},
		{"-", 1},
		{"唔係", 5},
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
		if right.Value != expected[i].right {
			t.Errorf("[%d] expected value '%d' got %d", i, expected[i].right, right.Value)
		}
	}
}

func TestInfixStatements(t *testing.T) {
	input := `
	1-1。
	1+2+3。
	1+3*2/5。
	10+x。
	hello+world。
	（1 + 2） + 3。
	1 + （2 + 3）。
	1 * （2 + 3）。
	`
	expected := []string{
		"(1 - 1)",
		"((1 + 2) + 3)",
		"(1 + ((3 * 2) / 5))",
		"(10 + x)",
		"(hello + world)",
		"((1 + 2) + 3)",
		"(1 + (2 + 3))",
		"(1 * (2 + 3))",
	}

	l := lexer.New(input)
	p := New(l)
	// p.ParseProgram()
	program := p.ParseProgram()
	checkParserErrors(p, t)
	if len(program.Statements) != len(expected) {
		t.Errorf("len(program) expected %d got %d", len(expected), len(program.Statements))
	}

	for i, st := range program.Statements {
		exprStatement, ok := st.(*ast.ExpressionStatement)
		if !ok {
			t.Errorf("[%d] expected type ast.ExpressionStatement got %T", i, st)
		}
		if exprStatement.Expression.String() != expected[i] {
			t.Logf("'%v'", exprStatement.Expression.(*ast.InfixExpression).Right.(*ast.Identifier))
			t.Errorf("[%d] expected string %s got %s.\n%+v", i, expected[i], exprStatement.Expression.String(), exprStatement.Expression)
		}
	}
}

func TestSingleIdentStatment(t *testing.T) {
	input := `x`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(p, t)
	if len(program.Statements) != 1 {
		t.Errorf("len(program) expected 1 got %d", len(program.Statements))
	}

	es, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Errorf("expected expressionStatement got %T", program.Statements[0])
	}
	if es.Expression.String() != "x" {
		t.Errorf("expression expected x got %s", es.Expression.String())
	}

}

func TestIfStatement(t *testing.T) {
	input := `如果 （a） 嘅話，就「
    2。
」`

	l := lexer.New(input)
	p := New(l)
	// p.ParseProgram()
	program := p.ParseProgram()
	checkParserErrors(p, t)
	if len(program.Statements) != 1 {
		t.Errorf("len(program) expected 1 got %d", len(program.Statements))
	}

	es, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Errorf("expected expressionStatement got %T", program.Statements[0])
	}

	ie, ok := es.Expression.(*ast.IfExpression)
	if !ok {
		t.Errorf("expected ifExpression got %T", es.Expression)
	}

	if ie.Condition.String() != "a" {
		t.Errorf("expected condition 'a' got %s", ie.Condition.String())
	}

	if len(ie.Consequence.Statements) != 1 {
		t.Errorf("expected 1 consq statement got %d", len(ie.Consequence.Statements))
	}

	cons, ok := ie.Consequence.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Errorf("cons not expressionStatement got %T", ie.Consequence.Statements[0])
	}

	if cons.String() != "2" {
		t.Errorf("cons.string() expected '2' got '%s'", cons.String())
	}

}

func TestIfElseStatement(t *testing.T) {
	input := `如果 （a） 嘅話，就「
	    2。
	」唔係就「
	    3。
	」
`

	l := lexer.New(input)
	p := New(l)
	// p.ParseProgram()
	program := p.ParseProgram()
	checkParserErrors(p, t)
	if len(program.Statements) != 1 {
		t.Errorf("len(program) expected 1 got %d", len(program.Statements))
	}

	es, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Errorf("expected expressionStatement got %T", program.Statements[0])
	}

	ie, ok := es.Expression.(*ast.IfExpression)
	if !ok {
		t.Errorf("expected ifExpression got %T", es.Expression)
	}

	if ie.Condition.String() != "a" {
		t.Errorf("expected condition 'a' got %s", ie.Condition.String())
	}

	if len(ie.Consequence.Statements) != 1 {
		t.Errorf("expected 1 consq statement got %d", len(ie.Consequence.Statements))
	}

	cons, ok := ie.Consequence.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Errorf("cons not expressionStatement got %T", ie.Consequence.Statements[0])
	}

	if cons.String() != "2" {
		t.Errorf("cons.string() expected '2' got '%s'", cons.String())
	}

	if len(ie.Alternative.Statements) != 1 {
		t.Errorf("expected 1 alt statement got %d", len(ie.Alternative.Statements))
	}

	alt, ok := ie.Alternative.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Errorf("alt not expressionStatement got %T", ie.Alternative.Statements[0])
	}

	if alt.String() != "3" {
		t.Errorf("alt.string() expected '3' got '%s'", alt.String())
	}

}

func TestReturnStatement(t *testing.T) {
	input := "俾我 1 + 2。"
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(p, t)
	if len(program.Statements) != 1 {
		t.Errorf("len(program) expected 1 got %d", len(program.Statements))
	}

	rs, ok := program.Statements[0].(*ast.ReturnStatement)
	if !ok {
		t.Errorf("expected ReturnStatement got %T", program.Statements[0])
	}
	if rs.Token.TokenType != token.RETURN {
		t.Errorf("expected return tok got %s", rs.Token.TokenType)
	}
	if rs.Expression.String() != "(1 + 2)" {
		t.Errorf("expected expression (1 + 2) got %s", rs.Expression.String())
	}
}

func TestFunctionDef(t *testing.T) {
	input := `聽到 add（x，y） 嘅話，就「
	    俾我 x + y。
	」`

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(p, t)
	if len(program.Statements) != 1 {
		t.Errorf("len(program) expected 1 got %d", len(program.Statements))
	}

	fd, ok := program.Statements[0].(*ast.FunctionDefStatment)
	if !ok {
		t.Errorf("expected FunctionDefStatement got %T", program.Statements[0])
	}

	if fd.String() != "聽到 add(x,y) {俾我(x + y)}" {
		t.Errorf("expected 聽到 add(x,y) {俾我(x + y)} got %s", fd.String())
	}

}

func TestFunctionCall(t *testing.T) {
	input := `add（x，y）`

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(p, t)
	if len(program.Statements) != 1 {
		t.Errorf("len(program) expected 1 got %d", len(program.Statements))
	}

	es, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Errorf("expected FunctionDefStatement got %T", program.Statements[0])
	}

	fce, ok := es.Expression.(*ast.FunctionCallExpression)
	if !ok {
		t.Errorf("expected FunctionCallExpression got %T", es.Expression)
	}
	if fce.Identifier.String() != "add" {
		t.Errorf("expected add got %s", fce.Identifier.String())
	}

}

func TestAssignStatement(t *testing.T) {
	input := "塞 (2 + 3) 入 i"
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(p, t)
	if len(program.Statements) != 1 {
		t.Errorf("len(program) expected 1 got %d", len(program.Statements))
	}

	rs, ok := program.Statements[0].(*ast.AssignStatement)
	if !ok {
		t.Errorf("expected AssignStatement got %T", program.Statements[0])
	}
	if rs.Token.TokenType != token.ASSIGN {
		t.Errorf("expected assign tok got %s", rs.Token.TokenType)
	}
	if rs.Expression.String() != "(2 + 3)" {
		t.Errorf("expected expression (2 + 3) got %s", rs.Expression.String())
	}
}
