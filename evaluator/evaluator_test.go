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
		{"塞 3 入 i; i;", 3},
		{"塞 5 入 i; 塞 3 入 j; i * j;", 15},
		{`
			塞 [1 + 2, "h" + "i"] 入 i;
			i[0];
		`, 3},
	}
	for _, test := range tests {
		output := testEval(t, test.input)
		intObj, ok := output.(*object.Integer)
		if !ok {
			t.Errorf("Expected object.Integer got %T (%+v)", output, output)
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
		{"唔係(6 大過 3)", false},
		{"唔係 唔係(6 大過 3)", true},
		{`"hi" 係 "amogus"`, false},
		{`"fart" 係 "fart"`, true},
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

func TestError(t *testing.T) {
	tests := []struct {
		input   string
		message string
	}{
		{"1 + 啱", "type mismatch"},
		{"2 係 錯", "type mismatch"},
		{"-啱", "invalid prefix"},
		{"唔係 3", "invalid prefix"},
		{"啱 + 錯", "invalid operation"},
		{"啱 大過 錯", "invalid comparison"},
		{"啱 大過 錯 + 錯", "invalid operation"},
		{"如果 (啱 大過 錯) 嘅話，就 {2} 唔係就 {3}", "invalid comparison"},
		{`有幾長（2）`, "invalid argument type"},
	}
	for _, test := range tests {
		output := testEval(t, test.input)
		err, ok := output.(*object.Error)
		if !ok {
			t.Errorf("Expected object.Error got %T", output)
		}
		if err.Message != test.message {
			t.Errorf("Expected msg: %s got %s", test.message, err.Message)
		}
	}
}

func testEval(t *testing.T, input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	return Eval(program, object.NewEnvironment(nil))
}

func TestFunction(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{`
		聽到 identity（x） 嘅話，就「
		     x。
		」;
		identity(15);
		`, 15},

		{`
		聽到 sum（x, y） 嘅話，就「
		     x + y。
		」;
		sum(15, 5);
		`, 20},

		{`
		聽到 identity（x） 嘅話，就「
		     俾我 x。
		」;
		identity(15);
		`, 15},

		{`
		聽到 sum（x, y） 嘅話，就「
		     俾我 x + y。
		」;
		sum(15, 5);
		`, 20},

		{`
		塞 3 入 x;
		聽到 sum（x, y） 嘅話，就「
		     俾我 x + y。
		」;
		sum(15, 5);
		x;
		`, 3},

		{`
		塞 3 入 x;
		聽到 sum（x, y） 嘅話，就「
		     俾我 x + y。
		」;
		sum(1, sum(2, sum(3, 4)));
		`, 10},
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

func TestString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{`"hello world"`, "hello world"},
		{`“塞 塞 塞”`, "塞 塞 塞"},
		{`"hello" + " " + "world"`, "hello world"},
		{`["hello", "world"][0]`, `"hello"`},
		{`
			塞 ["hello", "world"] 入 i;
			i[1];
		`, `"world"`},
	}
	for _, test := range tests {
		output := testEval(t, test.input)
		strObj, ok := output.(*object.String)
		if !ok {
			t.Errorf("Expected object.String got %T (%+v)", output, output)
		}
		if strObj.Value != test.expected {
			t.Errorf("expected %s got %+v (type %T)", test.expected, strObj.Value, strObj.Value)
		}
	}
}

func TestBuiltInFunction(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{`
		有幾長（"hello"）
		`, 5},
	}
	for _, test := range tests {
		output := testEval(t, test.input)
		intObj, ok := output.(*object.Integer)
		if !ok {
			t.Errorf("Expected object.Integer got %T (%+v)", output, output)
		}
		if intObj.Value != test.expected {
			t.Errorf("expected %d got %+v (type %T)", test.expected, intObj.Value, intObj.Value)
		}
	}
}

func TestArray(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{`["hello", "world"]`, `["hello", "world"]`},
		{"[]", "[]"},
		{`[1 + 2, "h" + "i"]`, `[3, "hi"]`},
		{`塞 [1 + 2, "h" + "i"] 入 i; i;`, `[3, "hi"]`},
	}

	for _, test := range tests {
		output := testEval(t, test.input)
		arr, ok := output.(*object.Array)
		if !ok {
			t.Errorf("Expected object.Array got %T (%+v)", output, output)
		}
		if arr.Inspect() != test.expected {
			t.Errorf("expected %s got %s", test.expected, arr.Inspect())
		}
	}
}
