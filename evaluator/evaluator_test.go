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
		{"-1", -1},
		{"-5", -5},
		{"1+3", 4},
		{"5-3", 2},
		{"5 - 3 + 3", 5},
		{"5 - (3 + 3)", -1},
		{"如果 (啱) 嘅話，就 {1} 唔係就 {2}", 1},
		{"如果 (錯) 嘅話，就 {1} 唔係就 {2}", 2},
		{"如果 (3) 嘅話，就 {1} 唔係就 {2}", 1},
		{"如果 (6 細過 3) 嘅話，就 {1} 唔係就 {2}", 2},
		{"如果 (2 細過 3) 嘅話，就 {1} 唔係就 {2}", 1},
	}
	for _, test := range tests {
		output := testEval(t, test.input)
		intObj, ok := output.(*object.Integer)
		if !ok {
			t.Errorf("Expected object.Integer got %T", output)
		}
		if intObj.Value != test.expected {
			t.Errorf("expected %d got %+v (type %T)", test.expected, intObj.Value, intObj.Value)
		}
	}
}

func TestBool(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"啱", true},
		{"錯", false},
		{"唔係 啱", false},
		{"唔係 錯", true},
		{"3 係 3", true},
		{"3 係 6", false},
		{"3 大過 6", false},
		{"6 大過 3", true},
		{"3 + 3 係 6", true},
	}
	for _, test := range tests {
		output := testEval(t, test.input)
		boolObj, ok := output.(*object.Boolean)
		if !ok {
			t.Errorf("Expected object.Boolean got %T", output)
		}
		if boolObj.Value != test.expected {
			t.Errorf("expected %t got %+v (type %T)", test.expected, boolObj.Value, boolObj.Value)
		}
	}
}

func TestNull(t *testing.T) {
	tests := []string{
		"1 + 啱",
		"2 係 錯",
		"啱 + 錯",
		"6 大過 false",
	}
	for _, input := range tests {
		output := testEval(t, input)
		_, ok := output.(*object.Null)
		if !ok {
			t.Errorf("Expected object.Null got %T", output)
		}
	}
}

func testEval(t *testing.T, input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	return Eval(program)
}
