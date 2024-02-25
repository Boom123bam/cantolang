package evaluator

import (
	"cantolang/lexer"
	"cantolang/object"
	"cantolang/parser"
	"testing"
)

func TestInteger(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"1", 1},
		{"5", 5},
	}
	for _, test := range tests {
		output := testEval(t, test.input)
		intObj, ok := output.(object.Integer)
		if !ok {
			t.Errorf("Expected object.Integer got %T", output)
		}
		if intObj.Value != test.expected {
			t.Errorf("expected %d got %+v (type %T)", test.expected, intObj.Value, intObj.Value)
		}
	}
}

func testEval(t *testing.T, input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	return Eval(program)
}