package lexer

import (
	"cantolang/token"
	"testing"
)

func TestLexer(t *testing.T) {
	input := `
    當 （i 細過 8） 時，就「
        講（i）。
        塞 i 加 1 入 i。
    」

    a 大D。
`

	expectedTokens := []struct {
		Type    string
		Literal string
	}{
		{token.WHILE, "當"},
		{token.OPEN_PAREN, "（"},
		{token.IDENTIFIER, "i"},
		{token.LESS_THAN, "細過"},
		{token.NUMBER, "8"},
		{token.CLOSE_PAREN, "）"},
		{token.SI, "時"},
		{token.COMMA, "，"},
		{token.THEN, "就"},
		{token.OPEN_BRACE, "「"},
		{token.IDENTIFIER, "講"},
		{token.OPEN_PAREN, "（"},
		{token.IDENTIFIER, "i"},
		{token.CLOSE_PAREN, "）"},
		{token.EOL, "。"},
		{token.ASSIGN, "塞"},
		{token.IDENTIFIER, "i"},
		{token.ADD, "加"},
		{token.NUMBER, "1"},
		{token.TO, "入"},
		{token.IDENTIFIER, "i"},
		{token.EOL, "。"},
		{token.CLOSE_BRACE, "」"},
		{token.IDENTIFIER, "a"},
		{token.INCREMENT, "大D"},
		{token.EOL, "。"},
		{token.EOF, ""},
	}
	l := New(input)
	for i, exp := range expectedTokens {
		got := l.ReadToken()
		if got.TokenLiteral != exp.Literal {
			t.Errorf("tests[%d] Expected literal '%s' got '%s'", i, exp.Literal, got.TokenLiteral)
		}
		if got.TokenType != exp.Type {
			t.Errorf("tests[%d] Expected type '%s' got '%s'", i, exp.Type, got.TokenType)
		}

	}
}

func TestSingleTok(t *testing.T) {
	input := `i`
	l := New(input)
	tok := l.ReadToken()
	if tok.TokenType != token.IDENTIFIER {
		t.Errorf("expected ident got %s", tok.TokenType)
	}
	if tok.TokenLiteral != "i" {
		t.Errorf("expected i got %s", tok.TokenLiteral)
	}
}

func TestIfElse(t *testing.T) {
	input := `如果 （a） 嘅話，就「
	    2。
	」唔係就「
	    3。
	」
`

	expectedTokens := []struct {
		Type    string
		Literal string
	}{
		{token.IF, "如果"},
		{token.OPEN_PAREN, "（"},
		{token.IDENTIFIER, "a"},
		{token.CLOSE_PAREN, "）"},
		{token.GEWA, "嘅話"},
		{token.COMMA, "，"},
		{token.THEN, "就"},
		{token.OPEN_BRACE, "「"},
		{token.NUMBER, "2"},
		{token.EOL, "。"},
		{token.CLOSE_BRACE, "」"},
		{token.ELSE, "唔係就"},
		{token.OPEN_BRACE, "「"},
		{token.NUMBER, "3"},
		{token.EOL, "。"},
		{token.CLOSE_BRACE, "」"},
		{token.EOF, ""},
	}
	l := New(input)
	for i, exp := range expectedTokens {
		got := l.ReadToken()
		if got.TokenLiteral != exp.Literal {
			t.Errorf("tests[%d] Expected literal '%s' got '%s'", i, exp.Literal, got.TokenLiteral)
		}
		if got.TokenType != exp.Type {
			t.Errorf("tests[%d] Expected type '%s' got '%s'", i, exp.Type, got.TokenType)
		}

	}
}

func TestFunction(t *testing.T) {
	input := `聽到 add（a，b） 嘅話，就「
	    俾我 a 加 b。
	」`

	expectedTokens := []struct {
		Type    string
		Literal string
	}{
		{token.FUNCTION, "聽到"},
		{token.IDENTIFIER, "add"},
		{token.OPEN_PAREN, "（"},
		{token.IDENTIFIER, "a"},
		{token.COMMA, "，"},
		{token.IDENTIFIER, "b"},
		{token.CLOSE_PAREN, "）"},
		{token.GEWA, "嘅話"},
		{token.COMMA, "，"},
		{token.THEN, "就"},
		{token.OPEN_BRACE, "「"},
		{token.RETURN, "俾我"},
		{token.IDENTIFIER, "a"},
		{token.ADD, "加"},
		{token.IDENTIFIER, "b"},
		{token.EOL, "。"},
		{token.CLOSE_BRACE, "」"},
		{token.EOF, ""},
	}
	l := New(input)
	for i, exp := range expectedTokens {
		got := l.ReadToken()
		if got.TokenLiteral != exp.Literal {
			t.Errorf("tests[%d] Expected literal '%s' got '%s'", i, exp.Literal, got.TokenLiteral)
		}
		if got.TokenType != exp.Type {
			t.Errorf("tests[%d] Expected type '%s' got '%s'", i, exp.Type, got.TokenType)
		}

	}
}

func TestRegularSymbols(t *testing.T) {
	input := `(){},;+-*/`

	expectedTokens := []struct {
		Type    string
		Literal string
	}{
		{token.OPEN_PAREN, "("},
		{token.CLOSE_PAREN, ")"},
		{token.OPEN_BRACE, "{"},
		{token.CLOSE_BRACE, "}"},
		{token.COMMA, ","},
		{token.EOL, ";"},
		{token.ADD, "+"},
		{token.MINUS, "-"},
		{token.MULTIPLY, "*"},
		{token.DIVIDE, "/"},
		{token.EOF, ""},
	}
	l := New(input)
	for i, exp := range expectedTokens {
		got := l.ReadToken()
		if got.TokenLiteral != exp.Literal {
			t.Errorf("tests[%d] Expected literal '%s' got '%s'", i, exp.Literal, got.TokenLiteral)
		}
		if got.TokenType != exp.Type {
			t.Errorf("tests[%d] Expected type '%s' got '%s'", i, exp.Type, got.TokenType)
		}

	}
}

func TestBool(t *testing.T) {
	input := `
	啱 錯。
	塞 啱 係 錯 入 i。
`

	expectedTokens := []struct {
		Type    string
		Literal string
	}{
		{token.TRUE, "啱"},
		{token.FALSE, "錯"},
		{token.EOL, "。"},
		{token.ASSIGN, "塞"},
		{token.TRUE, "啱"},
		{token.EQUAL_TO, "係"},
		{token.FALSE, "錯"},
		{token.TO, "入"},
		{token.IDENTIFIER, "i"},
		{token.EOL, "。"},
	}
	l := New(input)
	for i, exp := range expectedTokens {
		got := l.ReadToken()
		if got.TokenLiteral != exp.Literal {
			t.Errorf("tests[%d] Expected literal '%s' got '%s'", i, exp.Literal, got.TokenLiteral)
		}
		if got.TokenType != exp.Type {
			t.Errorf("tests[%d] Expected type '%s' got '%s'", i, exp.Type, got.TokenType)
		}

	}
}
