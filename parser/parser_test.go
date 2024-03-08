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
	tests := []struct {
		input    string
		expected string
	}{
		{`聽到 add（x，y） 嘅話，就「
			    俾我 x + y。
			」`, "聽到 add(x,y) {俾我(x + y)}",
		},
		{`聽到 one（） 嘅話，就「
			    俾我 1。
			」`, "聽到 one() {俾我1}",
		},
	}
	for _, test := range tests {
		l := lexer.New(test.input)
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

		if fd.String() != test.expected {
			t.Errorf("expected %s got %s", test.expected, fd.String())
		}

	}

}

func TestFunctionCall(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{`add（x，y）`, "add"},
		{`one（）`, "one"},
	}

	for _, test := range tests {
		l := lexer.New(test.input)
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
		if fce.Identifier.String() != test.expected {
			t.Errorf("expected add got %s", fce.Identifier.String())
		}
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

func TestStringStatements(t *testing.T) {
	input := `
	"joe"。
	"hi world"。
	`
	expected := []string{"joe", "hi world"}

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
		strLit, ok := exprStatement.Expression.(*ast.StringLiteral)
		if strLit.Value != expected[i] {
			t.Errorf("[%d] expected value '%s' got %s", i, expected[i], strLit.Value)
		}
	}
}

func TestArrayStatements(t *testing.T) {
	input := `
	[1,2,3]。
	["hello", "world"]。
	[]
	`
	expected := []string{"[1, 2, 3]", `["hello", "world"]`, "[]"}

	l := lexer.New(input)
	p := New(l)
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
		arr, ok := exprStatement.Expression.(*ast.ArrayLiteral)
		if arr.String() != expected[i] {
			t.Errorf("[%d] expected value '%s' got %s", i, expected[i], arr.String())
		}
	}
}

func TestIndexStatements(t *testing.T) {
	input := `
	[1,2,3][69]。
	[1,2,3][1 + 1]。
	`

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(p, t)

	if len(program.Statements) != 2 {
		t.Errorf("len(program) expected 2 got %d", len(program.Statements))
	}

	exprStatement, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Errorf("expected type ast.ExpressionStatement got %T", program.Statements[0])
	}
	exp, ok := exprStatement.Expression.(*ast.IndexExpression)
	if !ok {
		t.Errorf("expected ast.IndexExpression got %T", exprStatement.Expression)
	}
	num, ok := exp.Index.(*ast.IntegerLiteral)
	if !ok {
		t.Errorf("expected type ast.IntegerLiteral got %T", exp.Index)
	}
	if num.Value != 69 {
		t.Errorf("exp 69 got %d", num.Value)
	}

	exprStatement, ok = program.Statements[1].(*ast.ExpressionStatement)
	if !ok {
		t.Errorf("expected type ast.ExpressionStatement got %T", program.Statements[0])
	}
	exp, ok = exprStatement.Expression.(*ast.IndexExpression)
	if !ok {
		t.Errorf("expected ast.IndexExpression got %T", exprStatement.Expression)
	}
	ie, ok := exp.Index.(*ast.InfixExpression)
	if !ok {
		t.Errorf("expected type ast.IntegerLiteral got %T", exp.Index)
	}
	if ie.Infix.TokenType != token.ADD {
		t.Errorf("exp + got %s", ie.Infix)
	}

}

func TestWhileLoop(t *testing.T) {
	input := `
	當 （i 細過 8） 時，就「
	    塞 i+1 入 i。
    」`

	l := lexer.New(input)
	p := New(l)
	// p.ParseProgram()
	program := p.ParseProgram()
	checkParserErrors(p, t)
	if len(program.Statements) != 1 {
		t.Errorf("len(program) expected 1 got %d", len(program.Statements))
	}

	loop, ok := program.Statements[0].(*ast.WhileLoop)
	if !ok {
		t.Errorf("expected WhileLoop got %T", program.Statements[0])
	}

	if loop.Condition.String() != "(i 細過 8)" {
		t.Errorf("expected condition '(i 細過 8)' got %s", loop.Condition.String())
	}

	if len(loop.Body.Statements) != 1 {
		t.Errorf("expected 1 consq statement got %d", len(loop.Body.Statements))
	}

	body, ok := loop.Body.Statements[0].(*ast.AssignStatement)
	if !ok {
		t.Errorf("body not expressionStatement got %T", loop.Body.Statements[0])
	}

	if body.String() != "塞(i + 1)-> i" {
		t.Errorf("cons.string() expected '塞(i + 1)-> i' got '%s'", body.String())
	}

}

func TestIncrement(t *testing.T) {
	input := `
	i 大D;
	a 細D;
    `

	l := lexer.New(input)
	p := New(l)
	// p.ParseProgram()
	program := p.ParseProgram()
	checkParserErrors(p, t)
	if len(program.Statements) != 2 {
		t.Errorf("len(program) expected 2 got %d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.IncrementDecrementStatement)
	if !ok {
		t.Errorf("expected IncrementDecrementStatement got %T", program.Statements[0])
	}
	if stmt.Identifier != "i" {
		t.Errorf("expected Identifier 'i' got %s", stmt.Identifier)
	}
	if !stmt.IsIncrement {
		t.Errorf("stmt is not increment")
	}

	stmt2, ok := program.Statements[1].(*ast.IncrementDecrementStatement)
	if !ok {
		t.Errorf("expected IncrementDecrementStatement got %T", program.Statements[1])
	}
	if stmt2.Identifier != "a" {
		t.Errorf("expected Identifier 'a' got %s", stmt.Identifier)
	}
	if stmt2.IsIncrement {
		t.Errorf("stmt is not decrement")
	}
}
